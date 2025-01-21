package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/erespereza/new-project/pkg/formatter"
	_ "github.com/go-sql-driver/mysql"
)

// Connect establece una conexión con la base de datos
func Connect() {
	// Obtiene las variables de la base de datos desde el entorno (DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_CHARSET, DB_COLLATION)
	// Recomendación: Utilizar variables de entorno en lugar de hardcodear los datos de conexión a la base de datos para proteger la información sensible.
	// Consulte https://12factor.net/config para más información sobre cómo se deben manejar las variables de entorno en aplicaciones Go.

	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "" {
		dbDriver = "mysql"
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbCharset := os.Getenv("DB_CHARSET")
	dbCollation := os.Getenv("DB_COLLATION")

	// Construye la cadena de conexión
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&collation=%s&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbCharset, dbCollation)

	// Abre la conexión con la base de datos
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al abrir la conexión con la base de datos: %v", err)
	}

	// Verifica que la conexión sea válida
	if err := db.Ping(); err != nil {
		log.Fatalf("Error al verificar la conexión con la base de datos: %v", err)
	}

	log.Println("Conexión con la base de datos establecida.")
}

// Find busca un registro en la base de datos basado en el ID
// id puede ser int, float o string
// columns es opcional, si no se proporciona se usan todos los campos del struct
// se valida que almenos una de las columnas se valida para la busqueda
func (e *ExtendsModel) Find(id any, columns ...string) error {

	// Validar tipo de ID
	switch id.(type) {
	case int, int32, int64, float32, float64, string:
		// tipos válidos no se hace nada
	default:
		return fmt.Errorf("id debe ser numérico o string")
	}

	// Determinar las columnas a consultar
	var selectedColumns []string
	if len(columns) > 0 {
		// Usar las columnas proporcionadas
		selectedColumns = columns
	} else {
		// Usar las columnas del modelo
		selectedColumns = e.columns
	}

	// Si no hay columnas seleccionadas, retornar error
	if len(selectedColumns) == 0 {
		return fmt.Errorf("no hay campos válidos para consultar")
	}

	// Construir la consulta SQL
	query := fmt.Sprintf(
		"SELECT %s FROM %s WHERE id = $1",
		strings.Join(selectedColumns, ", "),
		e.tableName,
	)

	// Ejecutar la consulta
	row := db.QueryRow(query, id)

	if err := e.scanRows(e.model, row, selectedColumns); err != nil {
		return err
	}

	return nil
}

func (e *ExtendsModel) scanRows(m *Model, row *sql.Row, columns []string) error {

	// obtener el valor dinámico del modelo
	modelValue := reflect.ValueOf(m)
	// Obtener el tipo del struct
	structType := modelValue.Elem().Type()

	// Crear slice para almacenar los valores de las columnas
	// se hacen dos slices uno con los valores otro con las referencias por que a row.Scan hay que pasarle punteros
	values := make([]interface{}, len(columns))
	scanRefs := make([]interface{}, len(columns))
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
	for i, colName := range columns {
		// Convertir nombre de columna a PascalCase para que coincida con el del struct
		colName = formatter.ToPascalCase(colName)

		// Buscar el campo correspondiente en el struct
		var field reflect.Value
		for j := 0; j < structType.NumField(); j++ {
			if structType.Field(j).Name == colName {
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
			// aca hay un detalle que si no se puede convertir queda nulo ayayayay
		}
	}

	return nil
}
