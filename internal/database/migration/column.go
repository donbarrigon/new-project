package migration

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/donbarrigon/new-project/pkg/formatter"
)

// TinyInt crea una columna de tipo TINYINT con las opciones proporcionadas.
// alias de Int8
func TinyInt(name string, options ...string) *Column {
	return Int8(name, options...)
}

// Int8 crea una columna de tipo TINYINT con las opciones proporcionadas.
func Int8(name string, options ...string) *Column {
	return defaultColumn(name, "int8", options...)
}

// SmallInt crea una columna de tipo SMALLINT con las opciones proporcionadas.
// alias de Int16
func SmallInt(name string, options ...string) *Column {
	return Int16(name, options...)
}

// Int16 crea una columna de tipo SMALLINT con las opciones proporcionadas.
func Int16(name string, options ...string) *Column {
	return defaultColumn(name, "int16", options...)
}

// Integer crea una columna de tipo INTEGER con las opciones proporcionadas.
// alias de Int32
func Integer(name string, options ...string) *Column {
	return Int32(name, options...)
}

// Int32 crea una columna de tipo INTEGER con las opciones proporcionadas.
func Int32(name string, options ...string) *Column {
	return defaultColumn(name, "int32", options...)
}

// BigInt crea una columna de tipo BIGINT con las opciones proporcionadas.
// alias de Int64
func BigInt(name string, options ...string) *Column {
	return Int64(name, options...)
}

// Int64 crea una columna de tipo BIGINT con las opciones proporcionadas.
func Int64(name string, options ...string) *Column {
	return defaultColumn(name, "int64", options...)
}

// UTinyInt crea una columna de tipo TINYINT UNSIGNED con las opciones proporcionadas.
// alias de UInt8
func UTinyInt(name string, options ...string) *Column {
	return UInt8(name, options...)
}

// UInt8 crea una columna de tipo TINYINT UNSIGNED con las opciones proporcionadas.
func UInt8(name string, options ...string) *Column {
	return defaultColumn(name, "uint8", options...)
}

// USmallInt crea una columna de tipo SMALLINT UNSIGNED con las opciones proporcionadas.
// alias de UInt16
func USmallInt(name string, options ...string) *Column {
	return UInt16(name, options...)
}

// UInt16 crea una columna de tipo SMALLINT UNSIGNED con las opciones proporcionadas.
func UInt16(name string, options ...string) *Column {
	return defaultColumn(name, "uint16", options...)
}

// UInteger crea una columna de tipo SMALLINT UNSIGNED con las opciones proporcionadas.
// alias de UInt32
func UInteger(name string, options ...string) *Column {
	return UInt32(name, options...)
}

// UInt32 crea una columna de tipo INTEGER UNSIGNED con las opciones proporcionadas.
func UInt32(name string, options ...string) *Column {
	return defaultColumn(name, "uint32", options...)
}

// UBigInt crea una columna de tipo BIGINT UNSIGNED con las opciones proporcionadas.
// alias de UInt64
func UBigInt(name string, options ...string) *Column {
	return UInt64(name, options...)
}

// UInt64 crea una columna de tipo BIGINT UNSIGNED con las opciones proporcionadas.
func UInt64(name string, options ...string) *Column {
	return defaultColumn(name, "uint64", options...)
}

// TinyIncrements crea una columna id de clave primaria de tipo TINYINT UNSIGNED con las opciones proporcionadas.
func TinyIncrements(options ...string) *Column {
	return standardIncrementsColumn("uint8", options...)
}

// SmallIncrements crea una columna id de clave primaria de tipo SMALLINT UNSIGNED con las opciones proporcionadas.
func SmallIncrements(options ...string) *Column {
	return standardIncrementsColumn("uint16", options...)
}

// Increments crea una columna id de clave primaria de tipo INT UNSIGNED con las opciones proporcionadas.
func Increments(options ...string) *Column {
	return standardIncrementsColumn("uint32", options...)
}

// BigIncrements crea una columna id de clave primaria de tipo BIGINT UNSIGNED con las opciones proporcionadas.
func BigIncrements(options ...string) *Column {
	return standardIncrementsColumn("uint64", options...)
}

// Float32 crea una columna de tipo FLOAT con las opciones proporcionadas.
func Float32(name string, options ...string) *Column {
	return defaultColumn(name, "float32", options...)
}

// Float64 crea una columna de tipo DOUBLE con las opciones proporcionadas.
func Float64(name string, options ...string) *Column {
	return defaultColumn(name, "float64", options...)
}

// Time crea una columna tipo TIME con las opciones proporcionadas.
func Time(name string, options ...string) *Column {
	return defaultColumn(name, "time", options...)
}

// Date crea una columna tipo DATE con las opciones proporcionadas.
func Date(name string, options ...string) *Column {
	return defaultColumn(name, "date", options...)
}

// DateTime crea una columna tipo DATETIME con las opciones proporcionadas.
func DateTime(name string, options ...string) *Column {
	return defaultColumn(name, "datetime", options...)
}

// Timestamp crea una columna tipo TIMESTAMP con las opciones proporcionadas.
func Timestamp(name string, options ...string) *Column {
	return defaultColumn(name, "timestamp", options...)
}

// Timestamp crea una columna tipo TIMESTAMPTZ con las opciones proporcionadas.
func TimestampTz(name string, options ...string) *Column {
	return defaultColumn(name, "timestamptz", options...)
}

// TinyText crea una columna tipo TINYTEXT con las opciones proporcionadas.
func TinyText(name string, options ...string) *Column {
	return defaultColumn(name, "tinytext", options...)
}

// Text crea una columna tipo TEXT con las opciones proporcionadas.
func Text(name string, options ...string) *Column {
	return defaultColumn(name, "text", options...)
}

// MediumText crea una columna tipo MEDIUMTEXT con las opciones proporcionadas.
func MediumText(name string, options ...string) *Column {
	return defaultColumn(name, "mediumtext", options...)
}

// LongText crea una columna tipo LONGTEXT con las opciones proporcionadas.
func LongText(name string, options ...string) *Column {
	return defaultColumn(name, "longtext", options...)
}

// TinyBlob crea una columna tipo TINYBLOB con las opciones proporcionadas.
func TinyBlob(name string, options ...string) *Column {
	return defaultColumn(name, "tinyblob", options...)
}

// Blob crea una columna tipo BLOB con las opciones proporcionadas.
func Blob(name string, options ...string) *Column {
	return defaultColumn(name, "blob", options...)
}

// MediumBlob crea una columna tipo MEDIUMBLOB con las opciones proporcionadas.
func MediumBlob(name string, options ...string) *Column {
	return defaultColumn(name, "mediumblob", options...)
}

// LongBlob crea una columna tipo LONGBLOB con las opciones proporcionadas.
func LongBlob(name string, options ...string) *Column {
	return defaultColumn(name, "longblob", options...)
}

// BYTEA crea una columna tipo LONGBLOB con las opciones proporcionadas.
// alias de LongBlob para los que usan postgresql
func Bytea(name string, options ...string) *Column {
	return LongBlob(name, options...)
}

// JSON crea una columna tipo JSON con las opciones proporcionadas.
func Json(name string, options ...string) *Column {
	return defaultColumn(name, "json", options...)
}

// JSONB crea una columna tipo JSONB con las opciones proporcionadas.
func Jsonb(name string, options ...string) *Column {
	return defaultColumn(name, "jsonb", options...)
}

// Char crea una columna tipo CHAR con las opciones proporcionadas.
// Char($name) crea una columna $name CHAR(255)
// Char($name, $length) return $name CHAR($length)
// Char($name, "40") para que la funcion reconosca $length solo debes escribir un numero en un string
// y si $length se pasa como string culpa a las caracteristicas de go
// pero toco hacerlo asi para que se comportara como yo queria
func Char(name string, options ...string) *Column {
	var length int = 255  // length por defecto para CHAR
	if len(options) > 0 { // Verificar si hay opciones proporcionadas
		if parsedLength, err := strconv.Atoi(options[0]); err == nil && parsedLength > 0 {
			length = parsedLength // Si la primera opción se convierte a un número válido y es mayor a 0, se usa como longitud
			options = options[1:] // Remover el valor de longitud del resto de las opciones
		}
	}

	column := &Column{
		Name:        formatter.ToSnakeCase(name),
		Type:        "char",
		Nullable:    true,
		Precision:   &length, // Usar la longitud calculada
		Constraints: make(map[string]string),
	}

	// Procesar las demás opciones como "not_null", "default", etc.
	processOptions(column, options...)
	return column
}

// String crea una columna tipo VARCHAR con las opciones proporcionadas.
// String($name) crea una columna $name VARCHAR(255)
// String($name, $length) return $name VARCHAR($length)
// String($name, "40") para que la funcion reconosca $length solo debes escribir un numero en un string
// y si $length se pasa como string culpa a las caracteristicas de go
// pero toco hacerlo asi para que se comportara como yo queria
func String(name string, options ...string) *Column {
	var length int = 255  // length por defecto para VARCHAR
	if len(options) > 0 { // Verificar si hay opciones proporcionadas
		if parsedLength, err := strconv.Atoi(options[0]); err == nil && parsedLength > 0 {
			length = parsedLength // Si la primera opción se convierte a un número válido y es mayor a 0, se usa como longitud
			options = options[1:] // Remover el valor de longitud del resto de las opciones
		}
	}

	column := &Column{
		Name:        formatter.ToSnakeCase(name),
		Type:        "varchar",
		Nullable:    true,
		Precision:   &length, // Usar la longitud calculada
		Constraints: make(map[string]string),
	}

	// Procesar las demás opciones como "not_null", "default", etc.
	processOptions(column, options...)
	return column
}

// Decimal crea una columna tipo DECIMAL con las opciones proporcionadas.
func Decimal(name string, precision int, scale int, options ...string) *Column {
	column := &Column{
		Name:        formatter.ToSnakeCase(name),
		Type:        "decimal",
		Nullable:    true,
		Precision:   &precision,
		Scale:       &scale,
		Constraints: make(map[string]string),
	}
	processOptions(column, options...)
	return column
}

// Boolean crea una columna tipo BOOLEAN con las opciones proporcionadas.
// por defecto no pueder ser nulo
// se agrega restricion check para permitir solo TRUE/FALSE
func Boolean(name string, options ...string) *Column {
	name = formatter.ToSnakeCase(name)
	defaultValue := "false"                                     // Valor predeterminado por defecto
	checkConstraint := fmt.Sprintf("%s IN (TRUE, FALSE)", name) // Restricción para solo permitir TRUE/FALSE

	column := &Column{
		Name:        name,
		Type:        "boolean",
		Nullable:    false, // Por defecto, no permite valores nulos
		Default:     &defaultValue,
		Check:       &checkConstraint,
		Constraints: make(map[string]string),
	}

	// Procesar las opciones adicionales
	processOptions(column, options...)

	// Si la columna no permite NULL, se actualiza la restricción CHECK
	if !column.Nullable {
		checkConstraint = fmt.Sprintf("%s IN (TRUE, FALSE) AND %s IS NOT NULL", name, name)
		column.Check = &checkConstraint
	}

	return column
}

// CreateAt retorna una columna de tipo TIMESTAMP que se utiliza para la creación
func CreateAt(options ...string) *Column {
	defaultValue := "CURRENT_TIMESTAMP"
	column := &Column{
		Name:        "created_at",
		Type:        "timestamp",
		Nullable:    false,
		Default:     &defaultValue,
		Constraints: make(map[string]string),
	}
	processOptions(column, options...)
	return column
}

// UpdateAt retorna una columna de tipo TIMESTAMP que se utiliza para la actualización
func UpdateAt(options ...string) *Column {
	onUpdateValue := "CURRENT_TIMESTAMP"
	column := &Column{
		Name:        "updated_at",
		Type:        "timestamp",
		Nullable:    true,
		OnUpdate:    &onUpdateValue,
		Constraints: make(map[string]string),
	}
	processOptions(column, options...)
	return column
}

// DeleteAt retorna una columna de tipo TIMESTAMP que se utiliza para la eliminación (soft delete)
func DeleteAt(options ...string) *Column {
	column := &Column{
		Name:        "deleted_at",
		Type:        "timestamp",
		Nullable:    true,
		Index:       true,
		Constraints: make(map[string]string),
	}
	processOptions(column, options...)
	return column
}

// crea una variable de tipo ENUM en mysql en postgreSQL la simula con un VARCHAR con CHECK
func Enum(name string, values []string, options ...string) *Column {
	column := &Column{
		Name:        formatter.ToSnakeCase(name),
		Type:        "enum",
		Nullable:    true,
		Constraints: make(map[string]string),
	}
	// Agregar el constraint "CHECK" para simular un ENUM
	column.Constraints["enum"] = strings.Join(wrapValues(values), ", ")
	processOptions(column, options...)
	return column
}

// Binary crea una columna de tipo BINARY con las opciones proporcionadas.
func Binary(name string, options ...string) *Column {
	var length int = 255  // length por defecto
	if len(options) > 0 { // Verificar si hay opciones
		if parsedLength, err := strconv.Atoi(options[0]); err == nil && parsedLength > 0 {
			length = parsedLength // Si la primera opción se convierte a un número válido y es mayor a 0, se usa como longitud
			options = options[1:] // quito el primer elemneto
		}
	}

	column := &Column{
		Name:        formatter.ToSnakeCase(name),
		Type:        "binary",
		Precision:   &length,
		Nullable:    true,
		Constraints: make(map[string]string),
	}
	processOptions(column, options...)
	return column
}

// VarBinary crea una columna de tipo VARBINARY con las opciones proporcionadas.
func VarBinary(name string, options ...string) *Column {
	var length int = 255  // length por defecto
	if len(options) > 0 { // Verificar si hay opciones
		if parsedLength, err := strconv.Atoi(options[0]); err == nil && parsedLength > 0 {
			length = parsedLength // Si la primera opción se convierte a un número válido y es mayor a 0, se usa como longitud
			options = options[1:] // quito el primer elemneto
		}
	}

	column := &Column{
		Name:        formatter.ToSnakeCase(name),
		Type:        "varbinary",
		Precision:   &length,
		Nullable:    true,
		Constraints: make(map[string]string),
	}
	processOptions(column, options...)
	return column
}

// defaultColumn crea una columna numérica de tipo (Type) con las opciones proporcionadas.
// no usar para desarrollar esta es una funcion axiliar
// no deberias usar funciones privadas para crear las columnas de las migraciones
func defaultColumn(name string, t string, options ...string) *Column {
	column := &Column{
		Name:        formatter.ToSnakeCase(name),
		Type:        t,
		Nullable:    true,
		Constraints: make(map[string]string),
	}
	processOptions(column, options...)
	return column
}

// standardIncrementsColumn crea una columna id de clave primaria (type) UNSIGNED con las opciones proporcionadas.
// no usar para desarrollar esta es una funcion axiliar
// no deberias usar funciones privadas para crear las columnas de las migraciones
func standardIncrementsColumn(t string, options ...string) *Column {
	column := &Column{
		Name:          "id",
		Type:          t,
		Nullable:      false,
		AutoIncrement: true,
		PrimaryKey:    true,
		Constraints:   make(map[string]string),
	}
	column.Constraints["primary_key"] = fmt.Sprintf("PRIMARY KEY (%s)", "id")
	processOptions(column, options...)
	return column
}

// Options funcion para procesar las opciones de una columna
// no usar para desarrollar esta es una funcion axiliar
// no deberias usar funciones privadas para crear las columnas de las migraciones
func processOptions(column *Column, options ...string) {
	// Procesar cada opción proporcionada
	for _, option := range options {
		switch formatter.ToSnakeCase(option) {
		case "not_null":
			column.Nullable = false
		case "nullable":
			column.Nullable = true
		case "auto_increment":
			column.AutoIncrement = true
		case "serial":
			column.AutoIncrement = true
		case "primary_key":
			column.PrimaryKey = true
		case "unique":
			column.Unique = true
		case "index":
			column.Index = true
		default:
			// Manejar opciones con prefijo y valores
			if len(option) > 8 && strings.ToLower(option[:8]) == "default:" {
				defaultValue := option[8:] // Obtener el valor después de "default:"
				column.Default = &defaultValue
			} else if len(option) > 8 && strings.ToLower(option[:8]) == "comment:" {
				comment := option[8:] // Obtener el comentario después de "comment:"
				column.Comment = &comment
			} else if len(option) > 6 && strings.ToLower(option[:6]) == "check:" {
				check := option[6:] // Obtener la expresión de check
				column.Check = &check
			} else if len(option) > 9 && strings.ToLower(option[:9]) == "onupdate:" {
				onUpdate := option[9:] // Obtener el valor para OnUpdate
				column.OnUpdate = &onUpdate
			} else if len(option) > 3 && strings.ToLower(option[:3]) == "fk:" {
				// Procesar claves foráneas con formato "fk:referenced_table(column)"
				foreignKeyParts := option[3:] // Obtener lo que está después de "fk:"
				parts := strings.SplitN(foreignKeyParts, "(", 2)
				if len(parts) == 2 && strings.HasSuffix(parts[1], ")") {
					referencedTable := parts[0]
					referencedColumn := strings.TrimSuffix(parts[1], ")")
					column.ForeignKey = &ForeignKey{
						ReferencedTable:  referencedTable,
						ReferencedColumn: referencedColumn,
					}
				}
			} else {
				// Manejar constraints personalizados
				constraintParts := strings.SplitN(option, ":", 2)
				if len(constraintParts) == 2 {
					key := constraintParts[0]
					value := constraintParts[1]
					column.Constraints[key] = value
				} else {
					fmt.Printf("Advertencia: la opción '%s' no es reconocida.\n", option)
				}
			}
		}
	}
}

// wrapValues funcion auxiliar para envolver los valores con comillas simples
// no usar para desarrollar esta es una funcion axiliar
// no deberias usar funciones privadas para crear las columnas de las migraciones
func wrapValues(values []string) []string {
	for i, v := range values {
		values[i] = fmt.Sprintf("'%s'", v)
	}
	return values
}

// cosas que me fatan por hacer
// foreignId
// foreignIdFor
// foreignUlid
// foreignUuid
// geography
// geometry
// ipAddress
// macAddress
// morphs
// nullableMorphs
// nullableTimestamps
// nullableUlidMorphs
// nullableUuidMorphs
// rememberToken
// set
// ulidMorphs
// uuidMorphs
// ulid
// uuid
// vector
// year
