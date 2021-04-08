package controllers

import (
	"net/http"
	"retailStore/lib/db"

	"github.com/labstack/echo"
)

func GetOrdersController(c echo.Context) error {
	orders, err := db.GetOrders()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success getting orders",
		"data":   orders,
	})
}

func GetOrderByIdController(c echo.Context) error {
	order, err := db.GetOrderById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success getting order by id",
		"data":   order,
	})
}

func CreateOrderController(c echo.Context) error {
	order, err := db.CreateOrder(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success order created",
		"data":    order,
	})
}

func UpdateOrderByIdController(c echo.Context) error {
	order, err := db.UpdateOrderById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success deleting order by id",
		"data":    order,
	})
}

func DeleteOrderByIdController(c echo.Context) error {
	order, err := db.DeleteOrderById(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success deleting order by id",
		"data":   order,
	})
}



