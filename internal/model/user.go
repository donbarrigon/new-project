package model

import (
	. "github.com/donbarrigon/new-project/internal/orm"
)

// NewUser función estática (la más eficiente) para crear la instancia a un struct de tipo usuario
// func NewUser() *User {
// 	model := &User{}
// 	model.SetModel(model)
// 	model.Fillable()
// 	model.Guarded("password")
// 	return model
// }

// // Role modelo para roles
// type Role struct {
// 	ExtendsModel
// 	ID
// 	AllTimestamps
// 	Name        string `json:"name" db:"VARCHAR(50) UNIQUE NOT NULL"`
// 	Description string `json:"description" db:"TEXT"`
// }

// // Permission modelo para permisos
// type Permission struct {
// 	ExtendsModel
// 	ID
// 	AllTimestamps
// 	Name        string `json:"name" db:"VARCHAR(50) UNIQUE NOT NULL"`
// 	Description string `json:"description" db:"TEXT"`
// 	Module      string `json:"module" db:"VARCHAR(50)"` // Para agrupar permisos por módulo
// 	Action      string `json:"action" db:"VARCHAR(50)"` // create, read, update, delete
// }

// // UserRole tabla pivote para la relación many-to-many entre User y Role
// type UserRole struct {
// 	ExtendsModel
// 	UserId uint `json:"user_id" db:"PRIMARY KEY"`
// 	RoleId uint `json:"role_id" db:"PRIMARY KEY"`
// 	CreatedAt
// }

// // UserPermission tabla pivote para la relación many-to-many entre User y Permission
// type UserPermission struct {
// 	ExtendsModel
// 	UserId       uint `json:"user_id" db:"PRIMARY KEY"`
// 	PermissionID uint `json:"permission_id" db:"PRIMARY KEY"`
// 	CreatedAt
// }

// // RolePermission tabla pivote para la relación many-to-many entre Role y Permission
// type RolePermission struct {
// 	ExtendsModel
// 	RoleId       uint `json:"role_id" db:"PRIMARY KEY"`
// 	PermissionID uint `json:"permission_id" db:"PRIMARY KEY"`
// 	CreatedAt
// }

// // UserToken modelo para manejar tokens JWT
// type Token struct {
// 	ExtendsModel
// 	ID
// 	Timestamps
// 	UserId    uint      `json:"user_id" db:"NOT NULL"`
// 	Token     string    `json:"token" db:"TEXT NOT NULL"`
// 	ExpiresAt time.Time `json:"expires_at" db:"NOT NULL"`
// 	LastUsed  time.Time `json:"last_used" db:"DEFAULT CURRENT_TIMESTAMP"`
// 	UserAgent string    `json:"user_agent" db:"VARCHAR(255)"`
// 	IP        string    `json:"ip" db:"VARCHAR(45)"` // Para IPv6
// 	IsRevoked bool      `json:"is_revoked" db:"DEFAULT FALSE"`
// }
