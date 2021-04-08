package db

import (
	"retailStore/config"
	"retailStore/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetOrders() (interface{}, interface{}) {
	orders := []models.Order{}

	if err := config.DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrderById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	order := models.Order{}

	if err := config.DB.Find(&order, id).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func CreateOrder(c echo.Context) (interface{}, interface{}) {
	order := models.Order{}
	c.Bind(&order)
	if err := config.DB.Create(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func UpdateOrderById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	order := models.Order{}
	c.Bind(&order)
	order.ID = uint(id)
	if err := config.DB.Save(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func DeleteOrderById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	order := models.Order{}

	if err := config.DB.Delete(&order, id).Error; err != nil {
		return nil, err
	}
	return order, nil
}



