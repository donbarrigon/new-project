package models

import "time"

type Model struct {
	ID uint64 `json:"id" db:"BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY"`
	SoftDelete
}

type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SoftDelete struct {
	Timestamps
	DeletedAt time.Time `json:"deleted_at" db:"INDEX"`
}
