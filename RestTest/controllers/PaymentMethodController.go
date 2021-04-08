package controllers

import (
	"net/http"
	"retailStore/lib/db"

	"github.com/labstack/echo"
)

func GetPaymentMethodsController(c echo.Context) error {
	paymentMethods, err := db.GetPaymentMethods()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success getting paymentMethods",
		"data":   paymentMethods,
	})
}

func GetPaymentMethodByIdController(c echo.Context) error {
	paymentMethod, err := db.GetPaymentMethodById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success getting paymentMethod by id",
		"data":   paymentMethod,
	})
}

func CreatePaymentMethodController(c echo.Context) error {
	paymentMethod, err := db.CreatePaymentMethod(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success paymentMethod created",
		"data":    paymentMethod,
	})
}

func UpdatePaymentMethodByIdController(c echo.Context) error {
	paymentMethod, err := db.UpdatePaymentMethodById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success deleting paymentMethod by id",
		"data":    paymentMethod,
	})
}

func DeletePaymentMethodByIdController(c echo.Context) error {
	paymentMethod, err := db.DeletePaymentMethodById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success deleting paymentMethod by id",
		"data":   paymentMethod,
	})
}



