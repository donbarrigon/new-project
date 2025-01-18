package model

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

// Find busca un registro en la base de datos basado en el ID
// model debe ser un puntero a struct
// id puede ser int, float o string
// columns es opcional, si no se proporciona se usan todos los campos del struct
func Find(model any, id any, columns ...string) error {
	// Validar que model sea un puntero a struct
	modelValue := reflect.ValueOf(model)
	if modelValue.Kind() != reflect.Ptr || modelValue.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("model debe ser un puntero a struct")
	}

	// Validar tipo de ID
	switch id.(type) {
	case int, int32, int64, float32, float64, string:
		// tipos válidos no se hace nada
	default:
		return fmt.Errorf("id debe ser numérico o string")
	}

	// Obtener el tipo del struct
	structType := modelValue.Elem().Type()
	tableName := strings.ToLower(structType.Name())

	// Determinar las columnas a consultar
	var selectedColumns []string
	if len(columns) > 0 {
		// Usar las columnas proporcionadas
		selectedColumns = columns
	} else {
		// Usar todas las columnas del struct
		for i := 0; i < structType.NumField(); i++ {
			field := structType.Field(i)
			// Obtener el tag db si existe, sino usar el nombre del campo
			dbTag := field.Tag.Get("db")
			if dbTag == "" {
				dbTag = strings.ToLower(field.Name)
			}
			selectedColumns = append(selectedColumns, dbTag)
		}
	}

	// Construir la consulta SQL
	query := fmt.Sprintf(
		"SELECT %s FROM %s WHERE id = $1",
		strings.Join(selectedColumns, ", "),
		tableName,
	)

	// Ejecutar la consulta
	row := db.QueryRow(query, id)

	// Crear slice para almacenar los valores de las columnas
	values := make([]interface{}, len(selectedColumns))
	scanRefs := make([]interface{}, len(selectedColumns))
	for i := range values {
		scanRefs[i] = &values[i]
	}

	// Escanear resultados
	if err := row.Scan(scanRefs...); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("registro no encontrado")
		}
		return fmt.Errorf("error al escanear resultados: %v", err)
	}

	// Asignar valores al struct
	structValue := modelValue.Elem()
	for i, colName := range selectedColumns {
		// Buscar el campo correspondiente en el struct
		var field reflect.Value
		for j := 0; j < structType.NumField(); j++ {
			if strings.ToLower(structType.Field(j).Name) == colName ||
				structType.Field(j).Tag.Get("db") == colName {
				field = structValue.Field(j)
				break
			}
		}

		if field.IsValid() && field.CanSet() {
			// Convertir y asignar el valor
			val := reflect.ValueOf(values[i])
			if val.Type().ConvertibleTo(field.Type()) {
				field.Set(val.Convert(field.Type()))
			}
		}
	}

	return nil
}
