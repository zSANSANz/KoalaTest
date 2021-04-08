package controllers

import (
	"net/http"
	"retailStore/lib/db"

	"github.com/labstack/echo"
)

func GetOrderDetailsController(c echo.Context) error {
	orderDetails, err := db.GetOrderDetails()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success getting orderDetails",
		"data":   orderDetails,
	})
}

func GetOrderDetailByIdController(c echo.Context) error {
	orderDetail, err := db.GetOrderDetailById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success getting orderDetail by id",
		"data":   orderDetail,
	})
}

func CreateOrderDetailController(c echo.Context) error {
	orderDetail, err := db.CreateOrderDetail(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success orderDetail created",
		"data":    orderDetail,
	})
}

func UpdateOrderDetailByIdController(c echo.Context) error {
	orderDetail, err := db.UpdateOrderDetailById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success deleting orderDetail by id",
		"data":    orderDetail,
	})
}

func DeleteOrderDetailByIdController(c echo.Context) error {
	orderDetail, err := db.DeleteOrderDetailById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success deleting orderDetail by id",
		"data":   orderDetail,
	})
}



