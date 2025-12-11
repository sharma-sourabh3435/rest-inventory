package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/sourabh-sharma3435/rest-inventory/database"
	"github.com/sourabh-sharma3435/rest-inventory/models"
)

var storage []models.Item = []models.Item{}

func GetAllItems() []models.Item {
	var items []models.Item = []models.Item{}

	database.DB.Order("created_at desc").Find(&items)
	return items
}

func GetItemByID(id string) (models.Item, error) {
	// for _, item := range storage {
	// 	if item.ID == id {
	// 		return item, nil
	// 	}
	// }
	// return models.Item{}, errors.New("item not found")
	var item models.Item

	result := database.DB.First(&item, "id = ?", id)

	if result.RowsAffected == 0 {
		return models.Item{}, errors.New("item not found")
	}

	return item, nil
}

func CreateItem(itemRequest models.ItemRequest) models.Item {
	var newItem models.Item = models.Item{
		ID:        uuid.New().String(),
		Name:      itemRequest.Name,
		Price:     itemRequest.Price,
		Quantity:  itemRequest.Quantity,
		CreatedAt: time.Now(),
	}
	database.DB.Create(&newItem)
	return newItem
}

func UpdateItem(itemRequest models.ItemRequest, id string) (models.Item, error) {
	// for index, item := range storage {
	// 	if item.ID == id {
	// 		item.Name = itemRequest.Name
	// 		item.Price = itemRequest.Price
	// 		item.Quantity = itemRequest.Quantity
	// 		item.UpdatedAt = time.Now()

	// 		storage[index] = item
	// 		return item, nil
	// 	}
	// }
	// return models.Item{}, errors.New("item update failed, item not found")
	item, err := GetItemByID(id)
	if err != nil {
		return models.Item{}, err
	}

	item.Name = itemRequest.Name
	item.Price = itemRequest.Price
	item.Quantity = itemRequest.Quantity
	item.UpdatedAt = time.Now()

	database.DB.Save(&item)
	return item, nil
}

func DeleteItem(id string) bool {
	// var newItems []models.Item = []models.Item{}

	// for _, item := range storage {
	// 	if item.ID != id {
	// 		newItems = append(newItems, item)
	// 	}
	// }
	// storage = newItems
	// return true
	item, err := GetItemByID(id)

	if err != nil {
		return false
	}

	database.DB.Delete(&item)
	return true
}
