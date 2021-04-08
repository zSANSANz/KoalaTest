package db

import (
	"retailStore/config"
	"retailStore/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetProducts() (interface{}, interface{}) {
	products := []models.Product{}

	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := models.Product{}

	if err := config.DB.Find(&product, id).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func CreateProduct(c echo.Context) (interface{}, interface{}) {
	product := models.Product{}
	c.Bind(&product)
	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func UpdateProductById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := models.Product{}
	c.Bind(&product)
	product.ID = uint(id)
	if err := config.DB.Save(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func DeleteProductById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := models.Product{}

	if err := config.DB.Delete(&product, id).Error; err != nil {
		return nil, err
	}
	return product, nil
}



