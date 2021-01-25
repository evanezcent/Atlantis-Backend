package models

import (
	"time"
)

// User model for database
type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Image     string    `gorm:"type:varchar(255)" json:"image"`
	Phone     int8      `gorm:"uniqueIndex" json:"phone"`
	Email     string    `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password  string    `gorm:"->;<-; not null" json:"-"`
	Token     string    `gorm:"-" json:"token,omitempty"`
	DeletedAt time.Time `gorm:"null" json:"deleted_at"`
}
