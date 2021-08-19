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
	GetAllItem() []models.Combined
	GetAllItemImage(itemID uint64) []models.ImageItem
	FindItemByID(id uint64) models.Combined
	FindItemByUser(id uint64) []models.Combined
	FindItemByQuery(query string) []models.Combined
	ConfirmItem(id string) models.Item
}

type itemConnection struct {
	connection *gorm.DB
} 

// NewItemRepository used to create new Instance of item repository
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

func (db *itemConnection) ConfirmItem(id string) models.Item {
	var item models.Item
	db.connection.Find(&item, id)
	item.IsDone = !item.IsDone
	db.connection.Preload("ImageItem").Save(&item)

	return item
}

func (db *itemConnection) GetAllItem() []models.Combined {
	var items []models.Item
	var result []models.Combined

	db.connection.Preload("User").Find(&items)

	for _, element := range items {
		images := db.GetAllItemImage(element.ID) 
		result = append(result, models.Combined{Item: element, Images: images})
	}

	return result
}

func (db *itemConnection) GetAllItemImage(itemID uint64) []models.ImageItem {
	var images []models.ImageItem
	db.connection.Where("item_id = ?", itemID).Find(&images) 

	return images
}

func (db *itemConnection) FindItemByID(itemID uint64) models.Combined {
	var item models.Item
	var result models.Combined

	db.connection.Preload("User").Find(&item, itemID)

	images := db.GetAllItemImage(item.ID) 
	result = models.Combined{Item: item, Images: images} 

	return result
}

func (db *itemConnection) FindItemByUser(userID uint64) []models.Combined {
	var items []models.Item
	var result []models.Combined

	db.connection.Where("user_id = ?", userID).Preload("User").Find(&items)

	for _, element := range items {
		images := db.GetAllItemImage(element.ID) 
		result = append(result, models.Combined{Item: element, Images: images})
	}

	return result
}

func (db *itemConnection) FindItemByQuery(query string) []models.Combined {
	var items []models.Item
	var result []models.Combined
	db.connection.Where("title LIKE ?", ("%"+query+"%")).Or("spesific_place LIKE ?", ("%"+query+"%")).Or("description LIKE ?", ("%"+query+"%")).Preload("User").Find(&items)

	for _, element := range items {
		images := db.GetAllItemImage(element.ID) 
		result = append(result, models.Combined{Item: element, Images: images})
	}

	return result
}