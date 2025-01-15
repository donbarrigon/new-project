package models

import (
	"time"
)

type User struct {
	Model
	Email       string `json:"email" db:"VARCHAR(255) UNIQUE DEFAULT NULL"`    //unico pero opcional
	Password    string `json:"-" db:"VARCHAR(255) NOT NULL"`                   // Contraseñas no deben ser nulas
	PhoneNumber string `json:"phone_number" db:"VARCHAR(255) NOT NULL UNIQUE"` //obligatorio y unico
}

// Role modelo para roles
type Role struct {
	Model
	Name        string `json:"name" db:"VARCHAR(50) UNIQUE NOT NULL"`
	Description string `json:"description" db:"TEXT"`
}

// Permission modelo para permisos
type Permission struct {
	Model
	Name        string `json:"name" db:"VARCHAR(50) UNIQUE NOT NULL"`
	Description string `json:"description" db:"TEXT"`
	Module      string `json:"module" db:"VARCHAR(50)"` // Para agrupar permisos por módulo
	Action      string `json:"action" db:"VARCHAR(50)"` // create, read, update, delete
}

// UserRole tabla pivote para la relación many-to-many entre User y Role
type UserRole struct {
	UserID    uint      `json:"user_id" db:"PRIMARY KEY"`
	RoleID    uint      `json:"role_id" db:"PRIMARY KEY"`
	CreatedAt time.Time `json:"created_at" db:"DEFAULT CURRENT_TIMESTAMP"`
}

// UserPermission tabla pivote para la relación many-to-many entre User y Permission
type UserPermission struct {
	UserID       uint      `json:"user_id" db:"PRIMARY KEY"`
	PermissionID uint      `json:"permission_id" db:"PRIMARY KEY"`
	CreatedAt    time.Time `json:"created_at" db:"DEFAULT CURRENT_TIMESTAMP"`
}

// RolePermission tabla pivote para la relación many-to-many entre Role y Permission
type RolePermission struct {
	RoleID       uint      `json:"role_id" db:"PRIMARY KEY"`
	PermissionID uint      `json:"permission_id" db:"PRIMARY KEY"`
	CreatedAt    time.Time `json:"created_at" db:"DEFAULT CURRENT_TIMESTAMP"`
}

// UserToken modelo para manejar tokens JWT
type UserToken struct {
	Model
	UserID    uint      `json:"user_id" db:"NOT NULL"`
	Token     string    `json:"token" db:"TEXT NOT NULL"`
	ExpiresAt time.Time `json:"expires_at" db:"NOT NULL"`
	LastUsed  time.Time `json:"last_used" db:"DEFAULT CURRENT_TIMESTAMP"`
	UserAgent string    `json:"user_agent" db:"VARCHAR(255)"`
	IP        string    `json:"ip" db:"VARCHAR(45)"` // Para IPv6
	IsRevoked bool      `json:"is_revoked" db:"DEFAULT FALSE"`
}
