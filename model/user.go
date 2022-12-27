package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int64           `json:"id"`
	Name      string          `json:"name"`
	Username  string          `json:"username"`
	Password  string          `json:"-"`
	Polls     []Poll          `json:"polls"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"-"`
}
