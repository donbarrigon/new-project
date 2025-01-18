package model

import "time"

// Model ID es una estructura base para todos los modelos que requieren un ID autoincremental.
type ID struct {
	ID uint64 `json:"id" db:"BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY"`
}

// CreatedAt agrega el campo CreatedAt para registrar la fecha de creación.
type CreatedAt struct {
	CreatedAt time.Time `json:"created_at" db:"DEFAULT CURRENT_TIMESTAMP"`
}

// UpdatedAt agrega el campo UpdatedAt para registrar las últimas modificaciones.
type UpdatedAt struct {
	UpdatedAt time.Time `json:"updated_at" db:"DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

// SoftDelete agrega el campo DeletedAt para registros eliminados lógicamente.
type SoftDelete struct {
	DeletedAt time.Time `json:"deleted_at" db:"INDEX"`
}

// Timestamps representa los campos de fecha de creación y actualización.
type Timestamps struct {
	CreatedAt
	UpdatedAt
}

// AllTimestamps agrega los campos de timestamps y soft delete a la estructura base.
type AllTimestamps struct {
	Timestamps
	SoftDelete
}

// BaseModel es la estructura base para todos los modelos que necesitan un ID y timestamps.
type BaseModel struct {
	ID
	Timestamps
}

// ExtendedModel es la estructura base para todos los modelos que necesitan un ID, timestamps y soft delete.
type ExtendedModel struct {
	ID
	AllTimestamps
}
