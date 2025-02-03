package orm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Find realiza una busqueda en la base de datos por id y guarda los resultados en `m.Data`
// se puede de forma opcional seleccionar las columnas que desea buscar en la base de datos
// find valida que las columnas existan en la base de datos siempre y cuando halla creado la migracion
// si no se especifica, se usan las columnas de la migracion
// si no hay migraciones, se usan todas las columnas del slice columns
func (m *Model) Find(id any, columns ...string) error {

	// Validar tipo de ID
	switch id.(type) {
	case int, int8, int16, int32, int64, string:
		// tipos válidos no se hace nada
	default:
		return fmt.Errorf("id debe ser numérico o string")
	}

	if err := m.SetSelectedColumns(columns); err != nil {
		return err
	}

	// Usar un map para los drivers soportados
	findFuncs := map[string]func(any) error{
		"mongodb":    m.findMongoDB,
		"mysql":      m.findMySQL,
		"postgresql": m.findMySQL,
	}

	if findFunc, ok := findFuncs[dbDriver]; ok {
		return findFunc(id)
	}

	return fmt.Errorf("driver de base de datos '%s' no soportado", dbDriver)
}

// findMySQL hace la busqueda en mysql o postgresql
func (m *Model) findMySQL(id any) error {

	// Construir la consulta SQL
	query := fmt.Sprintf(
		"SELECT %s FROM %s WHERE id = ?",
		strings.Join(m.selectedColumns, ", "),
		m.tableName,
	)

	// Ejecutar la consulta
	row := db.QueryRow(query, id)

	// Crear un slice de punteros a interfaces para almacenar los valores
	values := make([]any, len(m.selectedColumns))
	valuePtrs := make([]any, len(m.selectedColumns))
	for i := range values {
		// hacer que los punteros apunten a los valores correspondientes
		valuePtrs[i] = &values[i]
	}

	// Crear slice de punteros con capacidad predefinida segun claude
	// values := make([]any, 0, len(m.selectedColumns))
	// valuePtrs := make([]any, 0, len(m.selectedColumns))

	// for range m.selectedColumns {
	// 	var v any
	// 	values = append(values, v)
	// 	valuePtrs = append(valuePtrs, &v)
	// }

	if err := row.Scan(valuePtrs...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("registro no encontrado")
		}
		return fmt.Errorf("error al escanear resultado: %w", err)
	}

	result := make(map[string]any, len(m.selectedColumns))
	for i, colName := range m.selectedColumns {
		result[colName] = values[i]
	}

	m.Data = []map[string]any{result}
	return nil
}

// findMongoDB hace la busqueda en mongodb
func (m *Model) findMongoDB(id any) error {
	// Timeout en el contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Convertir el ID a ObjectID si es necesario
	objID, err := m.convertToObjectID(id)
	if err != nil {
		return err
	}

	// Definir la colección basada en el nombre del modelo
	collection := dbc.Database(databaseName).Collection(m.tableName)

	// Construir el filtro de búsqueda
	filter := bson.M{"_id": objID}

	// Construir la proyección para devolver solo las columnas especificadas
	var opts *options.FindOneOptions
	projection := make(bson.M, len(m.selectedColumns))
	for _, col := range m.selectedColumns {
		projection[col] = 1
	}
	opts = options.FindOne().SetProjection(projection)

	var result map[string]any
	// realizar la consulta
	if err := collection.FindOne(ctx, filter, opts).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("documento no encontrado")
		}
		return fmt.Errorf("error en consulta MongoDB: %w", err)
	}

	m.Data = []map[string]any{result}
	return nil
}
