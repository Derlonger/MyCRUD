package models

type User struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
