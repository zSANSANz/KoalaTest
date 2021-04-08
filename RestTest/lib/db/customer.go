package db

import (
	"retailStore/config"
	"retailStore/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetCustomers() (interface{}, interface{}) {
	customers := []models.Customer{}

	if err := config.DB.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func GetCustomerById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("customer_id"))
	customer := models.Customer{}

	if err := config.DB.Find(&customer, id).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func CreateCustomer(c echo.Context) (interface{}, interface{}) {
	customer := models.Customer{}
	c.Bind(&customer)
	if err := config.DB.Create(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func UpdateCustomerById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("customer_id"))
	customer := models.Customer{}
	c.Bind(&customer)
	customer.CustomerId = uint(id)
	if err := config.DB.Save(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func DeleteCustomerById(c echo.Context) (interface{}, interface{}) {
	id, _ := strconv.Atoi(c.Param("customer_id"))
	customer := models.Customer{}

	if err := config.DB.Delete(&customer, id).Error; err != nil {
		return nil, err
	}
	return customer, nil
}



