package migration

var ColumnTypesMap = map[string]map[string]string{
	"mongodb": {
		"binary":      "binData",
		"varbinary":   "binData",
		"tinyblob":    "binData",
		"blob":        "binData",
		"mediumblob":  "binData",
		"longblob":    "binData",
		"char":        "string",
		"string":      "string",
		"enum":        "string",
		"tinytext":    "string",
		"text":        "string",
		"mediumtext":  "string",
		"longtext":    "string",
		"json":        "object",
		"jsonb":       "object",
		"int":         "int32",
		"int8":        "int32",
		"int16":       "int32",
		"int32":       "int32",
		"int64":       "int64",
		"uint":        "int64",
		"uint8":       "int32",
		"uint16":      "int32",
		"uint32":      "int64",
		"uint64":      "int64",
		"float32":     "double",
		"float64":     "double",
		"bool":        "bool",
		"time":        "date",
		"date":        "date",
		"datetime":    "date",
		"timestamp":   "date",
		"timestamptz": "date",
		"decimal":     "decimal128",
		"objectId":    "objectId",
		"array":       "array",
		"document":    "document",
	},
	"mysql": {
		"binary":      "binary",
		"varbinary":   "VARBINARY",
		"tinyblob":    "TINYBLOB",
		"blob":        "BLOB",
		"mediumblob":  "MEDIUMBLOB",
		"longblob":    "LONGBLOB",
		"char":        "CHAR",
		"string":      "VARCHAR",
		"enum":        "ENUM",
		"tinytext":    "TINYTEXT",
		"text":        "TEXT",
		"mediumtext":  "MEDIUMTEXT",
		"longtext":    "LONGTEXT",
		"json":        "JSON",
		"jsonb":       "JSON",
		"int":         "INT",
		"int8":        "TINYINT",
		"int16":       "SMALLINT",
		"int32":       "MEDIUMINT",
		"int64":       "BIGINT",
		"uint":        "INT UNSIGNED",
		"uint8":       "TINYINT UNSIGNED",
		"uint16":      "SMALLINT UNSIGNED",
		"uint32":      "MEDIUMINT UNSIGNED",
		"uint64":      "BIGINT UNSIGNED",
		"float32":     "FLOAT",
		"float64":     "DOUBLE",
		"bool":        "BOOLEAN",
		"time":        "TIME",
		"date":        "DATE",
		"datetime":    "DATETIME",
		"timestamp":   "TIMESTAMP",
		"timestamptz": "TIMESTAMP",
		"decimal":     "DECIMAL",
	},
	"postgresql": {
		"binary":      "BYTEA",
		"varbinary":   "BYTEA",
		"tinyblob":    "BYTEA",
		"blob":        "BYTEA",
		"mediumblob":  "BYTEA",
		"longblob":    "BYTEA",
		"char":        "CHAR",
		"string":      "VARCHAR",
		"enum":        "VARCHAR",
		"tinytext":    "TEXT",
		"text":        "TEXT",
		"mediumtext":  "TEXT",
		"longtext":    "TEXT",
		"json":        "JSON",
		"jsonb":       "JSONB",
		"int":         "INTEGER",
		"int8":        "SMALLINT",
		"int16":       "SMALLINT",
		"int32":       "INTEGER",
		"int64":       "BIGINT",
		"uint":        "INTEGER",
		"uint8":       "SMALLINT",
		"uint16":      "SMALLINT",
		"uint32":      "INTEGER",
		"uint64":      "BIGINT",
		"float32":     "REAL",
		"float64":     "DOUBLE PRECISION",
		"bool":        "BOOLEAN",
		"time":        "TIME",
		"date":        "DATE",
		"datetime":    "TIMESTAMP",
		"timestamp":   "TIMESTAMP",
		"timestamptz": "TIMESTAMPTZ",
		"decimal":     "NUMERIC",
	},
}

var ConstraintsMap = map[string]map[string]string{
	"mongodb": {
		"primary_key":    "_id",      // MongoDB usa _id como clave primaria
		"unique":         "unique",   // Restricción de unicidad en un índice
		"not_null":       "required", // Equivalente a NOT NULL en esquemas de validación
		"auto_increment": "objectId", // MongoDB usa ObjectId para generar identificadores únicos automáticamente
		"check":          "expr",     // Se pueden usar expresiones en validaciones de esquemas
		"default":        "default",  // Valor por defecto en esquemas de validación
		"index":          "index",    // Índice en MongoDB
		"foreign_key":    "lookup",   // Se simula con agregaciones ($lookup)
		"on_delete":      "cascade",  // Se puede simular con triggers en la aplicación
		"on_update":      "manual",   // No hay ON UPDATE, debe manejarse manualmente
	},
	"mysql": {
		"primary_key":    "PRIMARY KEY",    // Clave primaria
		"unique":         "UNIQUE",         // Restricción única
		"not_null":       "NOT NULL",       // No nulo
		"auto_increment": "AUTO_INCREMENT", // MySQL: Autoincremento
		"check":          "CHECK",          // Restricción de verificación
		"default":        "DEFAULT",        // Valor por defecto
		"index":          "INDEX",          // Índice
		"foreign_key":    "FOREIGN KEY",    // Clave foránea
		"on_delete":      "ON DELETE",      // Restricción para eliminar en clave foránea
		"on_update":      "ON UPDATE",      // Restricción para actualizar en clave foránea
	},
	"postgresql": {
		"primary_key":    "PRIMARY KEY", // Clave primaria
		"unique":         "UNIQUE",      // Restricción única
		"not_null":       "NOT NULL",    // No nulo
		"auto_increment": "SERIAL",      // PostgreSQL: Autoincremento
		"check":          "CHECK",       // Restricción de verificación
		"default":        "DEFAULT",     // Valor por defecto
		"index":          "INDEX",       // Índice
		"foreign_key":    "FOREIGN KEY", // Clave foránea
		"on_delete":      "ON DELETE",   // Restricción para eliminar en clave foránea
		"on_update":      "ON UPDATE",   // Restricción para actualizar en clave foránea
	},
}

type Column struct {
	Name          string            // Nombre de la columna
	Type          string            // Tipo de dato (por ejemplo, VARCHAR, INT, DECIMAL)           // Longitud (si aplica, como VARCHAR(255) o DECIMAL(10,2))
	Precision     *int              // Precisión para tipos como VARCHAR(255) o DECIMAL(10,2)
	Scale         *int              // Escala para tipos como DECIMAL (opcional)
	Required      bool              // Indica si la columna permite valores NULL
	AutoIncrement bool              // Indica si la columna es autoincremental
	PrimaryKey    bool              // Indica si es una clave primaria
	Unique        bool              // Indica si tiene una restricción de UNIQUE
	ForeignKey    ForeignKey        // Relación con clave foránea (opcional)
	Default       *string           // Valor por defecto (si aplica)
	Check         *string           // Expresión para restricciones CHECK (opcional)
	Comment       *string           // Comentario de la columna (opcional)
	Index         bool              // Indica si se crea un índice simple en esta columna
	Generated     *string           // Definición para columnas generadas (GENERATED AS o similar)
	OnUpdate      *string           // Valor en operaciones UPDATE (como CURRENT_TIMESTAMP)
	Constraints   map[string]string // Otros constraints personalizados
}

type Index struct {
	Columns []string // Columnas que componen el índice
	Unique  bool     // Indica si el índice es único
	Name    string   // Nombre opcional del índice
}

// ALTER TABLE ordenes
// ADD CONSTRAINT fk_usuario
// FOREIGN KEY (usuario_id)
// REFERENCES usuarios(id);
type ForeignKey struct {
	Column    string // Columna que actúa como clave foránea
	Reference string // Columna de la tabla de referencia
	Table     string // Tabla de referencia
	OnDelete  string // Acción en cascada (CASCADE, SET NULL, etc.)
	OnUpdate  string // Acción al actualizar
}

type Table struct {
	Name               string            // Nombre de la tabla
	Columns            []Column          // Lista de columnas
	Engine             string            // Motor de la tabla (MyISAM, InnoDB, etc.)
	Charset            string            // Conjunto de caracteres (utf8mb4, etc.)
	Collation          string            // Collation de la tabla
	PrimaryKeys        []string          // ColumnNames como claves primarias
	Constraints        map[string]string // Restricciones
	AutoIncrementStart int               // Valor inicial del auto_increment
	Temporary          bool              // Indica si es una tabla temporal
	Comment            string            // Comentario de la tabla
}

type Schema struct {
	Name      string  // Nombre del schema
	Tables    []Table // Mapa de tablas del schema
	Charset   string  // Charset por defecto para el schema
	Collation string  // Collation por defecto para el schema
}
