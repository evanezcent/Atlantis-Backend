package repository

import (
	"Atlantis-Backend/models"

	"gorm.io/gorm"
)

// ItemRepository as interface that cover all function
type ItemRepository interface {
	InsertItem(item models.Item) models.Item
	UpdateItem(item models.Item) models.Item
	UploadImage(item models.ImageItem) models.ImageItem
}

type itemConnection struct {
	connection *gorm.DB
}

// NewItemRepository used to create new Instance of user repository
func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemConnection{
		connection: db,
	}
}

func (db *itemConnection) InsertItem(item models.Item) models.Item {
	db.connection.Save(&item)

	return item
}

func (db *itemConnection) UploadImage(image models.ImageItem) models.ImageItem {
	db.connection.Save(&image)

	return image
}

func (db *itemConnection) UpdateItem(item models.Item) models.Item {
	var tempItem models.Item
	db.connection.Find(&tempItem, item.ID)
	db.connection.Save(&item)

	return item
}
