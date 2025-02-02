package orm

import (
	"github.com/donbarrigon/new-project/internal/cache"
	"github.com/donbarrigon/new-project/internal/database/migration"
	"github.com/donbarrigon/new-project/lib/formatter"
)

// Model defines the contract for database models, inspired by Laravel's Eloquent ORM.
type ModelInterface interface {

	// Table le dice al Model la tabla o coleccion con la que va a trabajar y define si existe tiene o no migracion.
	Table(name string)

	// Fillable establece los atributos que son asignables en masa (mass-assignment).
	// Recibe una lista de nombres de atributos que se pueden llenar de manera masiva.
	Fillable(fields ...string)

	// Guarded establece los atributos que no deben ser asignados de manera masiva.
	// Recibe una lista de nombres de atributos que están protegidos de ser asignados masivamente.
	Guarded(fields ...string)

	// BeforeSave se ejecuta antes de crear o actualizar el modelo en la base de datos.
	BeforeSave(hook func() error) error

	// AfterSave se ejecuta despues de crear o actualizar el modelo en la base de datos.
	AfterSave(hook func() error) error

	// BeforeDelete se ejecuta antes de eliminar el modelo de la base de datos.
	BeforeDelete(hook func() error) error

	// AfterDelete se ejecuta despues de eliminar el modelo de la base de datos.
	AfterDelete(hook func() error) error

	// BeforeCreate se ejecuta antes de crear el modelo en la base de datos.
	BeforeCreate(hook func() error) error

	// AfterCreate se ejecuta despues de crear el modelo en la base de datos.
	AfterCreate(hook func() error) error

	// BeforeUpdate se ejecuta antes de actualizar el modelo de la base de datos.
	BeforeUpdate(hook func() error) error

	// AfterUpdate se ejecuta despues de actualizar el modelo de la base de datos.
	AfterUpdate(hook func() error) error

	// Load loads a relationship (HasMany, BelongsTo, ManyToMany, etc.).
	//Load(name ...string) any

	// Find encuentra un modelo por el valor de primary key.
	// Find(id any, columns ...string) error
}

// estructura base para modelos
type Model struct {
	tableName    string           // tableName es el nombre de la tabla
	table        *migration.Table //estrutura de la base de datos para la migracion
	hasMigration bool             // Indica si el modelo tiene una migración asociada
	fillable     []string         // Fillable establece los atributos que son asignables en masa (mass-assignment).
	guarded      []string         // Guarded establece los atributos que no deben ser asignados de manera masiva.

	// variables que se usaran al construir la consulta
	// selectColumns []string
	// where         string
	// orderBy       string
	// join          []Join
	// limit int
}

func (m *Model) Table(name string) {
	m.tableName = formatter.ToTableName(name)
	// toma la estructura de la tabla segun la migracion
	m.table = cache.GetTable(m.tableName)
	if m.table != nil {
		m.hasMigration = true
	}
}

func (m *Model) Fillable(fields ...string) {
	m.fillable = fields
}

func (m *Model) Guarded(fields ...string) {
	m.guarded = fields
}

// funciones abstractas
func (e *Model) BeforeSave(hook func() error) error {
	return hook()
}

func (e *Model) AfterSave(hook func() error) error {
	return hook()
}

func (e *Model) BeforeDelete(hook func() error) error {
	return hook()
}

func (e *Model) AfterDelete(hook func() error) error {
	return hook()
}

func (e *Model) BeforeCreate(hook func() error) error {
	return hook()
}

func (e *Model) AfterCreate(hook func() error) error {
	return hook()
}

func (e *Model) BeforeUpdate(hook func() error) error {
	return hook()
}

func (e *Model) AfterUpdate(hook func() error) error {
	return hook()
}

// // ID es una estructura base para modelos que requieran un id con autoincremento
// type ID struct {
// 	Id uint64 `json:"id" db:"BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY"`
// }

// // TinyID es una estructura base para modelos que requieran un id con autoincremento de tipo TINYINT
// type TinyID struct {
// 	Id uint8 `json:"id" db:"TINYINT UNSIGNED AUTO_INCREMENT PRIMARY KEY"`
// }

// // SmallID es una estructura base para modelos que requieran un id con autoincremento de tipo SMALLINT
// type SmallID struct {
// 	Id uint16 `json:"id" db:"SMALLINT UNSIGNED AUTO_INCREMENT PRIMARY KEY"`
// }

// // IntegerID es una estructura base para modelos que requieran un id con autoincremento de tipo INTEGER
// type IntegerID struct {
// 	Id uint32 `json:"id" db:"INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY"`
// }

// // UUID es una estructura base para modelos que requieran un id unico de tipo UUID
// type UUID struct {
// 	UUID string `json:"uuid" db:"CHAR(36) UNIQUE NOT NULL DEFAULT (UUID())"`
// }

// // CreateAt agrega el campo create_at para registrar la fecha de creación.
// type CreatedAt struct {
// 	CreatedAt time.Time `json:"created_at" db:"DEFAULT CURRENT_TIMESTAMP"`
// }

// // UpdatedAt agrega el campo updated_at para registrar la fecha de actualización.
// type UpdatedAt struct {
// 	UpdatedAt time.Time `json:"updated_at" db:"DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
// }

// // SoftDelete agrega el campo deleted_at para registrar la fecha de eliminación lógica.
// type SoftDelete struct {
// 	DeletedAt time.Time `json:"deleted_at" db:"INDEX"`
// }

// // Active agrega el campo active para establecer si un modelo está activo.
// type Active struct {
// 	Active bool `json:"active" db:"BOOLEAN DEFAULT 1"`
// }

// // ActiveAt agrega el campo active_at para registrar la fecha de activación.
// type ActiveAt struct {
// 	ActiveAt time.Time `json:"active_at" db:"DEFAULT CURRENT_TIMESTAMP"`
// }

// // Priority agrega el campo priority para establecer la prioridad de un modelo.
// type Priority struct {
// 	Priority int `json:"priority" db:"INT DEFAULT 0"`
// }

// // Timestamps agrega los campos create_at y updated_at para registrar la fecha de creación y actualización.
// type Timestamps struct {
// 	CreatedAt
// 	UpdatedAt
// }

// // AllTimestamps agrega los campos create_at, updated_at y deleted_at para registrar la fecha de creación, actualización y eliminación lógica.
// type AllTimestamps struct {
// 	Timestamps
// 	SoftDelete
// }

// Asigna permite al struct embebido saber cual es el struct modelo !HERMOSA CARACTERISTICA
// se analiza el strcut y almacenan valores relevantes que despues seran usados
// Se hace una validación para ignorar campos que sean:
// Structs (fieldKind == reflect.Struct)
// Slices (fieldKind == reflect.Slice)
// Arrays (fieldKind == reflect.Array)
// Punteros a structs (fieldKind == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct)
// func (e *ExtendsModel) SetModel(m Model) {

// 	// Hacer reflect al modelo
// 	modelValue := reflect.ValueOf(m)

// 	// Verificar que el valor sea un puntero
// 	if modelValue.Kind() != reflect.Ptr || modelValue.IsNil() {
// 		panic("SetModel requiere un puntero al modelo")
// 	}

// 	//guardar el puntero al modelo
// 	e.model = &m

// 	// Obtener el tipo del struct
// 	structType := modelValue.Elem().Type()
// 	e.tableName = formatter.ToTableName(structType.Name())

// 	// Usar solo los campos públicos que no sean structs o colecciones
// 	for i := 0; i < structType.NumField(); i++ {
// 		field := structType.Field(i)

// 		// Verificar si el campo es público
// 		if field.PkgPath != "" {
// 			continue // Campo privado, ignorar
// 		}

// 		// Verificar si el campo es un struct o una colección de structs
// 		fieldKind := field.Type.Kind()
// 		if fieldKind == reflect.Struct ||
// 			fieldKind == reflect.Slice ||
// 			fieldKind == reflect.Array ||
// 			(fieldKind == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct) {
// 			continue // Es un struct o colección, ignorar
// 		}

// 		// Añadir el nombre del campo al slice de columnas
// 		e.columns = append(e.columns, formatter.ToSnakeCase(field.Name))
// 	}
// }
