package database

import (
	"fmt"
	"strings"

	"github.com/erespereza/new-project/pkg/formatter"
)

var ColumnTypesMYSQL = map[string]string{
	"char":      "CHAR",    // CHAR(255) max 255
	"string":    "VARCHAR", // VARCHAR(255) podría almacenar hasta 65,535 bytes / 4 = ???
	"text":      "TEXT",
	"int":       "INT",
	"int8":      "TINYINT",
	"int16":     "SMALLINT",
	"int32":     "MEDIUMINT",
	"int64":     "BIGINT",
	"uint":      "INT UNSIGNED",
	"uint8":     "TINYINT UNSIGNED",
	"uint16":    "SMALLINT UNSIGNED",
	"uint32":    "MEDIUMINT UNSIGNED",
	"uint64":    "BIGINT UNSIGNED",
	"float32":   "FLOAT",
	"float64":   "DOUBLE",
	"bool":      "BOOLEAN",
	"time":      "TIME",
	"date":      "DATE",
	"datetime":  "DATETIME",
	"timestamp": "TIMESTAMP",
	"decimal":   "DECIMAL",
}

var ColumnTypesPostgreSQL = map[string]string{
	"char":      "CHAR",    // Debe tener un tamaño (opcional, por ejemplo CHAR(255))
	"string":    "VARCHAR", // Debe tener un tamaño (opcional, por ejemplo VARCHAR(255))
	"text":      "TEXT",
	"int":       "INTEGER",
	"int8":      "SMALLINT",
	"int16":     "SMALLINT",
	"int32":     "INTEGER",
	"int64":     "BIGINT",
	"uint":      "INTEGER",
	"uint8":     "SMALLINT",
	"uint16":    "SMALLINT",
	"uint32":    "INTEGER",
	"uint64":    "BIGINT",
	"float32":   "REAL",
	"float64":   "DOUBLE PRECISION",
	"bool":      "BOOLEAN",
	"time":      "TIME",
	"date":      "DATE",
	"datetime":  "TIMESTAMP",
	"timestamp": "TIMESTAMP",
	"decimal":   "NUMERIC",
}

var ConstraintsMySQL = map[string]string{
	"primary_key":    "PRIMARY KEY",    // Clave primaria
	"unique":         "UNIQUE",         // Restricción única
	"not_null":       "NOT NULL",       // No nulo
	"auto_increment": "AUTO_INCREMENT", // MySQL: Autoincremento
	"serial":         "AUTO_INCREMENT", // MySQL: Autoincremento
	"check":          "CHECK",          // Restricción de verificación
	"default":        "DEFAULT",        // Valor por defecto
	"index":          "INDEX",          // Índice
	"foreign_key":    "FOREIGN KEY",    // Clave foránea
	"on_delete":      "ON DELETE",      // Restricción para eliminar en clave foránea
	"on_update":      "ON UPDATE",      // Restricción para actualizar en clave foránea
}

var ConstraintsPostgreSQL = map[string]string{
	"primary_key":    "PRIMARY KEY", // Clave primaria
	"unique":         "UNIQUE",      // Restricción única
	"not null":       "NOT NULL",    // No nulo
	"auto_increment": "SERIAL",      // PostgreSQL: Autoincremento
	"serial":         "SERIAL",      // PostgreSQL: Autoincremento
	"check":          "CHECK",       // Restricción de verificación
	"default":        "DEFAULT",     // Valor por defecto
	"index":          "INDEX",       // Índice
	"foreign_key":    "FOREIGN KEY", // Clave foránea
	"on_delete":      "ON DELETE",   // Restricción para eliminar en clave foránea
	"on_update":      "ON UPDATE",   // Restricción para actualizar en clave foránea
}

type Column struct {
	Name   string // Nombre de la columna
	Type   string // Tipo de dato (por ejemplo, VARCHAR, INT, DECIMAL)
	Length *int   // Longitud (si aplica, como VARCHAR(255) o DECIMAL(10,2))
	// Precision     *int              // Precisión para tipos como DECIMAL (opcional) se puede usar length para esto y es cun campo menos
	Scale         *int              // Escala para tipos como DECIMAL (opcional)
	Nullable      bool              // Indica si la columna permite valores NULL
	AutoIncrement bool              // Indica si la columna es autoincremental
	PrimaryKey    bool              // Indica si es una clave primaria
	Unique        bool              // Indica si tiene una restricción de UNIQUE
	ForeignKey    *ForeignKey       // Relación con clave foránea (opcional)
	Default       *string           // Valor por defecto (si aplica)
	Check         *string           // Expresión para restricciones CHECK (opcional)
	Comment       *string           // Comentario de la columna (opcional)
	Index         bool              // Indica si se crea un índice simple en esta columna
	Generated     *string           // Definición para columnas generadas (GENERATED AS o similar)
	OnUpdate      *string           // Valor en operaciones UPDATE (como CURRENT_TIMESTAMP)
	Constraints   map[string]string // Otros constraints personalizados
}

type Constraint struct {
	Name string // Nombre del constraint
	Rule string // Regla del constraint (ejemplo: "age > 0")
}

type Index struct {
	Columns []string // Columnas que componen el índice
	Unique  bool     // Indica si el índice es único
	Name    string   // Nombre opcional del índice
}

type ForeignKey struct {
	Column           string // Columna que actúa como clave foránea
	ReferencedTable  string // Tabla de referencia
	ReferencedColumn string // Columna de la tabla de referencia
	OnDelete         string // Acción en cascada (CASCADE, SET NULL, etc.)
	OnUpdate         string // Acción al actualizar
}

type Migration struct {
	TableName          string       // Nombre de la tabla
	Columns            []Column     // Lista de columnas
	Engine             string       // Motor de la tabla (MyISAM, InnoDB, etc.)
	Charset            string       // Conjunto de caracteres (utf8mb4, etc.)
	Collation          string       // Collation de la tabla
	PrimaryKeys        []string     // Claves primarias
	Indexes            []Index      // Índices
	ForeignKeys        []ForeignKey // Claves foráneas
	Constraints        []Constraint // Restricciones
	AutoIncrementStart int          // Valor inicial del auto_increment
	Temporary          bool         // Indica si es una tabla temporal
	Comment            string       // Comentario de la tabla
}

func Integer(name string, options ...string) Column {
	column := Column{
		Name:        name,
		Type:        "int",
		Nullable:    true, // Por defecto, la columna permite NULL
		Constraints: make(map[string]string),
	}

	Options(&column, options...)

	return column
}

func Options(column *Column, options ...string) {
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
