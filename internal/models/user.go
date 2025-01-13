package models

import (
	"time"
)

// User modelo base para autenticación
type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Email       string    `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password    string    `json:"-" gorm:"type:varchar(255);not null"` // "-" excluye el campo de las respuestas JSON
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(255);unique;default:null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"index"`
}

// Profile información personal y del juego
type Profile struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	UserID       uint      `json:"user_id" gorm:"not null"`
	FirstName    string    `json:"first_name" gorm:"type:varchar(255)"`
	LastName     string    `json:"last_name" gorm:"type:varchar(255)"`
	GameNickname string    `json:"game_nickname" gorm:"type:varchar(255)"`
	GameRole     string    `json:"game_role" gorm:"type:varchar(255)"`
	GameID       string    `json:"game_id" gorm:"type:varchar(255)"`
	AccountType  string    `json:"account_type" gorm:"type:varchar(255)"` // F2P, LowSpender, MidSpender, WhaleSpender
	ClanID       uint      `json:"clan_id" gorm:"default:null"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at" gorm:"index"`
}

// Role modelo para roles
type Role struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"type:varchar(50);unique;not null"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"index"`
}

// Permission modelo para permisos
type Permission struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"type:varchar(50);unique;not null"`
	Description string    `json:"description" gorm:"type:text"`
	Module      string    `json:"module" gorm:"type:varchar(50)"` // Para agrupar permisos por módulo
	Action      string    `json:"action" gorm:"type:varchar(50)"` // create, read, update, delete
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"index"`
}

// UserRole tabla pivote para la relación many-to-many entre User y Role
type UserRole struct {
	UserID    uint      `json:"user_id" gorm:"primaryKey"`
	RoleID    uint      `json:"role_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
}

// UserPermission tabla pivote para la relación many-to-many entre User y Permission
type UserPermission struct {
	UserID       uint      `json:"user_id" gorm:"primaryKey"`
	PermissionID uint      `json:"permission_id" gorm:"primaryKey"`
	CreatedAt    time.Time `json:"created_at"`
}

// RolePermission tabla pivote para la relación many-to-many entre Role y Permission
type RolePermission struct {
	RoleID       uint      `json:"role_id" gorm:"primaryKey"`
	PermissionID uint      `json:"permission_id" gorm:"primaryKey"`
	CreatedAt    time.Time `json:"created_at"`
}

// UserToken modelo para manejar tokens JWT
type UserToken struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Token     string    `json:"token" gorm:"type:text;not null"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
	LastUsed  time.Time `json:"last_used"`
	UserAgent string    `json:"user_agent" gorm:"type:varchar(255)"`
	IP        string    `json:"ip" gorm:"type:varchar(45)"` // Para IPv6
	IsRevoked bool      `json:"is_revoked" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
