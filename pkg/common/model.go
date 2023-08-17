package common

import (
	"time"
)

type Model struct {
	ID        int `json:"id" gorm:"autoIncrement,primaryKey;column:id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
