package models

import (
	"time"
)

// ImageItem model for database
type ImageItem struct {
	ID     uint64 `gorm:"primary_key;auto_increment" json:"id"`
	URL    string `gorm:"type:text" json:"url"`
	ItemID string `gorm:"not null,foreignkey:ItemID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"item_id"`
	// Item      Item      `gorm:"foreignkey:ItemID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"item"`
	DeletedAt time.Time `gorm:"null" json:"deleted_at"`
}
