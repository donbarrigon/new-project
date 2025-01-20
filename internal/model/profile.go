package model

// Profile información personal y del juego
type Profile struct {
	ID
	AllTimestamps
	UserID       uint   `json:"user_id" db:"NOT NULL"`
	Name         string `json:"name" db:"VARCHAR(255)"`
	GameNickname string `json:"game_nickname" db:"VARCHAR(255)"`
	GameRole     string `json:"game_role" db:"VARCHAR(255)"`
	GameID       string `json:"game_id" db:"VARCHAR(255)"`
	AccountType  string `json:"account_type" db:"VARCHAR(255)"` // F2P, LowSpender, MidSpender, WhaleSpender
	ClanID       uint   `json:"clan_id" db:"DEFAULT NULL"`
}
