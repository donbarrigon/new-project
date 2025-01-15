package models

import "time"

// Model es una estructura base para todos los modelos que requieren un ID autoincremental.
type Model struct {
	ID uint64 `json:"id" db:"BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY"`
	SoftDelete
}

// Timestamps representa los campos de fecha de creación y actualización.
type Timestamps struct {
	CreatedAt time.Time `json:"created_at" db:"DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" db:"DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

// SoftDelete agrega el campo DeletedAt para registros eliminados lógicamente.
type SoftDelete struct {
	Timestamps
	DeletedAt time.Time `json:"deleted_at" db:"INDEX"`
}
