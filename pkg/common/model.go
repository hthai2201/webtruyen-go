package common

import (
	"time"
)

type Model struct {
	ID        int       `json:"id" gorm:"autoIncrement,primaryKey;column:id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoIncrement,primaryKey;column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoIncrement,primaryKey;column:updated_at"`
}
