package db

import (
	"retailStore/config"
	"retailStore/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetOrderDetails() (interface{}, interface{}) {
	orderDetails := []models.OrderDetail{}

	if err := config.DB.Find(&orderDetails).Error; err != nil {
		return nil, err
	}
	return orderDetails, nil
}

func GetOrderDetailById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	orderDetail := models.OrderDetail{}

	if err := config.DB.Find(&orderDetail, id).Error; err != nil {
		return nil, err
	}
	return orderDetail, nil
}

func CreateOrderDetail(c echo.Context) (interface{}, interface{}) {
	orderDetail := models.OrderDetail{}
	c.Bind(&orderDetail)
	if err := config.DB.Create(&orderDetail).Error; err != nil {
		return nil, err
	}
	return orderDetail, nil
}

func UpdateOrderDetailById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	orderDetail := models.OrderDetail{}
	c.Bind(&orderDetail)
	orderDetail.ID = uint(id)
	if err := config.DB.Save(&orderDetail).Error; err != nil {
		return nil, err
	}
	return orderDetail, nil
}

func DeleteOrderDetailById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	orderDetail := models.OrderDetail{}

	if err := config.DB.Delete(&orderDetail, id).Error; err != nil {
		return nil, err
	}
	return orderDetail, nil
}



