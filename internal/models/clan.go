package models

// Clan informacion de los clanes
type Clan struct {
	Model
	Name        string `json:"name" db:"VARCHAR(255) NOT NULL"`
	Description string `json:"description" db:"TEXT NULL"`
	IsMain      bool   `json:"is_main" db:"BOOLEAN DEFAULT FALSE"`
}
