package models

import (
	"time"
)

// Item model for database
type Item struct {
	ID             uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title          string    `gorm:"type:varchar(255)" json:"title"`
	DateLostFound  time.Time `gorm:"not null" json:"date_lostfound"`
	PlaceLostFound string    `gorm:"type:text" json:"place_lostfound"`
	Description    string    `gorm:"type:text" json:"description"`
	UserID         string    `gorm:"not null" json:"-"`
	User           User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	DeletedAt      time.Time `gorm:"null" json:"deleted_at"`
}
