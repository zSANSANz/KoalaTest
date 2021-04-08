package controllers

import (
	"net/http"
	"retailStore/lib/db"

	"github.com/labstack/echo"
)

func GetCustomersController(c echo.Context) error {
	customers, err := db.GetCustomers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success getting customers",
		"data":   customers,
	})
}

func GetCustomerByIdController(c echo.Context) error {
	customer, err := db.GetCustomerById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success getting customer by id",
		"data":   customer,
	})
}

func CreateCustomerController(c echo.Context) error {
	customer, err := db.CreateCustomer(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success customer created",
		"data":    customer,
	})
}

func UpdateCustomerByIdController(c echo.Context) error {
	customer, err := db.UpdateCustomerById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success deleting customer by id",
		"data":    customer,
	})
}

func DeleteCustomerByIdController(c echo.Context) error {
	customer, err := db.DeleteCustomerById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success deleting customer by id",
		"data":   customer,
	})
}



