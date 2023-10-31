package model

import (
	"database/sql"
	"time"
)

type Model struct {
	ID        uint         `gorm:"primarykey;" json:"_"`
	CreatedAt time.Time    `json:"_"`
	UpdatedAt time.Time    `json:"_"`
	DeletedAt sql.NullTime `gorm:"index;"json:"_"`
}
