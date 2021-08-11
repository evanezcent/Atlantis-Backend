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
	Insert(item dto.ItemCreateDTO) models.Item
	Update(item dto.ItemUpdateDTO) models.Item
	UploadImage(item dto.ItemImageCreateDTO) models.ImageItem
	GetAll() []models.Combined
	AuthorizeForEdit(userID string, ItemID uint64) bool
	ConfirmItem(ItemID string) models.Item
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

func (service *itemService) Insert(item dto.ItemCreateDTO) models.Item {
	newItem := models.Item{}
	err := smapping.FillStruct(&newItem, smapping.MapFields(&item))

	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	res := service.itemRepository.InsertItem(newItem)
	return res
}

func (service *itemService) UploadImage(item dto.ItemImageCreateDTO) models.ImageItem {
	newImage := models.ImageItem{}
	err := smapping.FillStruct(&newImage, smapping.MapFields(&item))

	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	res := service.itemRepository.UploadImage(newImage)
	return res
}

func (service *itemService) Update(item dto.ItemUpdateDTO) models.Item {
	newItem := models.Item{}
	err := smapping.FillStruct(&newItem, smapping.MapFields(&item))

	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	res := service.itemRepository.UpdateItem(newItem)
	return res
}

func (service *itemService) GetAll() []models.Combined {
	res := service.itemRepository.GetAllItem()
	return res
}

func (service *itemService) AuthorizeForEdit(userID string, itemID uint64) bool {
	item := service.itemRepository.FindItemByID(itemID)
	id := fmt.Sprintf(item.UserID)

	return userID == id
}

func (service *itemService) ConfirmItem(itemID string) models.Item {
	item := service.itemRepository.ConfirmItem(itemID)
	return item
}
