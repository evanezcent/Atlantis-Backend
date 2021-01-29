package service

import (
	"Atlantis-Backend/dto"
	"Atlantis-Backend/models"
	"Atlantis-Backend/repository"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
)

// ItemService interface for Item service
type ItemService interface {
	InsertUpdate(item dto.ItemCreateDTO, tipe string) models.Item
	GetAll() []models.Item
	AuthorizeForEdit(userID string, ItemID uint64) bool
}

type itemService struct {
	itemRepository repository.ItemRepository
}

// NewItemService is new Instance
func NewItemService(repo repository.ItemRepository) ItemService {
	return &itemService{
		itemRepository: repo,
	}
}

func (service *itemService) InsertUpdate(item dto.ItemCreateDTO, tipe string) models.Item {
	newItem := models.Item{}
	err := smapping.FillStruct(&newItem, smapping.MapFields(&item))

	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	var res models.Item
	if tipe == "insert" {

		res = service.itemRepository.InsertItem(newItem)
	} else {
		res = service.itemRepository.UpdateItem(newItem)
	}
	return res
}

func (service *itemService) GetAll() []models.Item {
	return service.itemRepository.GetAllItem()
}

func (service *itemService) AuthorizeForEdit(userID string, itemID uint64) bool {
	Item := service.itemRepository.FindItemByID(itemID)
	id := fmt.Sprintf(Item.UserID)

	return userID == id
}
