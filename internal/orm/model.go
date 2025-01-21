package orm

import (
	"reflect"
	"time"

	"github.com/erespereza/new-project/pkg/formatter"
)

// Model defines the contract for database models, inspired by Laravel's Eloquent ORM.
type Model interface {
	// TableName returns the name of the table associated with the model.
	TableName() string

	// Columns retorna las columnas en del modelo
	Columns() []string

	// SetModel le dice al ExtendsModel el modelo con el cual debe trabajar
	// por que es caracteristica de GO que los embebidos no puedan acceder a los atributos del heredero
	SetModel(m Model)

	// BeforeSave se ejecuta antes de crear o actualizar el modelo en la base de datos.
	BeforeSave() error

	// AfterSave se ejecuta despues de crear o actualizar el modelo en la base de datos.
	AfterSave() error

	// BeforeDelete se ejecuta antes de eliminar el modelo de la base de datos.
	BeforeDelete() error

	// AfterDelete se ejecuta despues de eliminar el modelo de la base de datos.
	AfterDelete() error

	// BeforeCreate se ejecuta antes de crear el modelo en la base de datos.
	BeforeCreate() error

	// AfterCreate se ejecuta despues de crear el modelo en la base de datos.
	AfterCreate() error

	// BeforeUpdate se ejecuta antes de actualizar el modelo de la base de datos.
	BeforeUpdate() error

	// AfterUpdate se ejecuta despues de actualizar el modelo de la base de datos.
	AfterUpdate() error

	// Fillable establece los atributos que son asignables en masa (mass-assignment).
	// Recibe una lista de nombres de atributos que se pueden llenar de manera masiva.
	Fillable(f ...string)

	// Guarded establece los atributos que no deben ser asignados de manera masiva.
	// Recibe una lista de nombres de atributos que están protegidos de ser asignados masivamente.
	Guarded(g ...string)

	// Load loads a relationship (HasMany, BelongsTo, ManyToMany, etc.).
	//Load(name ...string) any

	// Find encuentra un modelo por el valor de primary key.
	Find(id any, columns ...string) error
}

// ID es una estructura base para modelos que requieran un id con autoincremento
type ID struct {
	Id uint64 `json:"id" db:"BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY"`
}

// TinyID es una estructura base para modelos que requieran un id con autoincremento de tipo TINYINT
type TinyID struct {
	Id uint8 `json:"id" db:"TINYINT UNSIGNED AUTO_INCREMENT PRIMARY KEY"`
}

// SmallID es una estructura base para modelos que requieran un id con autoincremento de tipo SMALLINT
type SmallID struct {
	Id uint16 `json:"id" db:"SMALLINT UNSIGNED AUTO_INCREMENT PRIMARY KEY"`
}

// IntegerID es una estructura base para modelos que requieran un id con autoincremento de tipo INTEGER
type IntegerID struct {
	Id uint32 `json:"id" db:"INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY"`
}

// UUID es una estructura base para modelos que requieran un id unico de tipo UUID
type UUID struct {
	UUID string `json:"uuid" db:"CHAR(36) UNIQUE NOT NULL DEFAULT (UUID())"`
}

// CreateAt agrega el campo create_at para registrar la fecha de creación.
type CreatedAt struct {
	CreatedAt time.Time `json:"created_at" db:"DEFAULT CURRENT_TIMESTAMP"`
}

// UpdatedAt agrega el campo updated_at para registrar la fecha de actualización.
type UpdatedAt struct {
	UpdatedAt time.Time `json:"updated_at" db:"DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

// SoftDelete agrega el campo deleted_at para registrar la fecha de eliminación lógica.
type SoftDelete struct {
	DeletedAt time.Time `json:"deleted_at" db:"INDEX"`
}

// Active agrega el campo active para establecer si un modelo está activo.
type Active struct {
	Active bool `json:"active" db:"BOOLEAN DEFAULT 1"`
}

// ActiveAt agrega el campo active_at para registrar la fecha de activación.
type ActiveAt struct {
	ActiveAt time.Time `json:"active_at" db:"DEFAULT CURRENT_TIMESTAMP"`
}

// Priority agrega el campo priority para establecer la prioridad de un modelo.
type Priority struct {
	Priority int `json:"priority" db:"INT DEFAULT 0"`
}

// Timestamps agrega los campos create_at y updated_at para registrar la fecha de creación y actualización.
type Timestamps struct {
	CreatedAt
	UpdatedAt
}

// AllTimestamps agrega los campos create_at, updated_at y deleted_at para registrar la fecha de creación, actualización y eliminación lógica.
type AllTimestamps struct {
	Timestamps
	SoftDelete
}

type ExtendsModel struct {
	model         *Model
	tableName     string
	columns       []string
	selectColumns []string
	fillable      []string
	guarded       []string
	// selectColumns []string
	// where         [][]string
	// orderBy       []string
	// join          []Join
	// limit         int
}

func (e *ExtendsModel) TableName() string {
	return e.tableName
}

func (e *ExtendsModel) Columns() []string {
	return e.columns
}

// Asigna permite al struct embebido saber cual es el struct modelo !HERMOSA CARACTERISTICA
// se analiza el strcut y almacenan valores relevantes que despues seran usados
// Se hace una validación para ignorar campos que sean:
// Structs (fieldKind == reflect.Struct)
// Slices (fieldKind == reflect.Slice)
// Arrays (fieldKind == reflect.Array)
// Punteros a structs (fieldKind == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct)
func (e *ExtendsModel) SetModel(m Model) {

	// Hacer reflect al modelo
	modelValue := reflect.ValueOf(m)

	// Verificar que el valor sea un puntero
	if modelValue.Kind() != reflect.Ptr || modelValue.IsNil() {
		panic("SetModel requiere un puntero al modelo")
	}

	//guardar el puntero al modelo
	e.model = &m

	// Obtener el tipo del struct
	structType := modelValue.Elem().Type()
	e.tableName = formatter.ToTableName(structType.Name())

	// Usar solo los campos públicos que no sean structs o colecciones
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		// Verificar si el campo es público
		if field.PkgPath != "" {
			continue // Campo privado, ignorar
		}

		// Verificar si el campo es un struct o una colección de structs
		fieldKind := field.Type.Kind()
		if fieldKind == reflect.Struct ||
			fieldKind == reflect.Slice ||
			fieldKind == reflect.Array ||
			(fieldKind == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct) {
			continue // Es un struct o colección, ignorar
		}

		// Añadir el nombre del campo al slice de columnas
		e.columns = append(e.columns, formatter.ToSnakeCase(field.Name))
	}
}

func (e *ExtendsModel) Fillable(f ...string) {
	e.fillable = f
}

func (e *ExtendsModel) Guarded(g ...string) {
	e.guarded = g
}

func (e *ExtendsModel) BeforeSave() error {
	return nil
}

func (e *ExtendsModel) AfterSave() error {
	return nil
}

func (e *ExtendsModel) BeforeDelete() error {
	return nil
}

func (e *ExtendsModel) AfterDelete() error {
	return nil
}

func (e *ExtendsModel) BeforeCreate() error {
	return nil
}

func (e *ExtendsModel) AfterCreate() error {
	return nil
}

func (e *ExtendsModel) BeforeUpdate() error {
	return nil
}

func (e *ExtendsModel) AfterUpdate() error {
	return nil
}
