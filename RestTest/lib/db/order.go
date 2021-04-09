package db

import (
	"retailStore/config"
	"retailStore/models"
)

func GetOrders() (interface{}, error) {
	var order []models.Order

	if err := config.DB.Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func GetOrderById(id string) (interface{}, error) {
	var order []models.Order

	if err := config.DB.First(&order, "order_id=?", id).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func CreateOrder(order *models.Order) (interface{}, error) {
	if err := config.DB.Save(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func UpdateOrder(id string, order *models.Order) (interface{}, error) {
	var existingOrder models.Order
	if err := config.DB.First(&existingOrder, "order_id=?", id).Error; err != nil {
		return nil, err
	}
	existingOrder.CustomerId = order.CustomerId
	existingOrder.OrderNumber = order.OrderNumber
	if updateErr := config.DB.Save(&existingOrder).Where("order_id=?", id).Error; updateErr != nil {
		return nil, updateErr
	}
	return existingOrder, nil
}

func DeleteOrder(id string) (interface{}, error) {
	var order models.Order
	if err := config.DB.First(&order, "order_id=?", id).Error; err != nil {
		return nil, err
	}
	if deleteErr := config.DB.Delete(&order).Where("order_id=?", id).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return order, nil
}
