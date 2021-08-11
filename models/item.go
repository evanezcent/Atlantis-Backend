package models

// Item model for database
type Item struct {
	ID            uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title         string `gorm:"type:varchar(255)" json:"title"`
	SpesificDate  string `gorm:"not null" json:"spesific_date"`
	SpesificPlace string `gorm:"type:text" json:"spesific_place"`
	Description   string `gorm:"type:text" json:"description"`
	UserID        string `gorm:"not null" json:"-"`
	User          User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	IsDone        bool   `gorm:"default:false" json:"is_done"`
	DeletedAt     string `gorm:"null" json:"deleted_at"`
}

type Combined struct {
	Item   Item
	Images []ImageItem
}
