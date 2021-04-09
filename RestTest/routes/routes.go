package routes

import (
	"os"
	"retailStore/controllers"
	"retailStore/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	middlewares.LogMiddlewares(e)
	eJWT := e.Group("")
	eJWT.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eJWT.GET("/customers", controllers.GetCustomersController)
	eJWT.GET("/customers/:id", controllers.GetCustomerByIdController)
	eJWT.POST("/customers", controllers.CreateCustomerController)
	eJWT.PUT("/customers/:id", controllers.UpdateCustomerByIdController)
	eJWT.DELETE("/customers/:id", controllers.DeleteCustomerByIdController)

	e.POST("/login", controllers.LoginCustomerController)

	eJWT.GET("/orders", controllers.GetOrdersController)
	eJWT.GET("/orders/:id", controllers.GetOrderByIdController)
	eJWT.POST("/orders", controllers.CreateOrderController)
	eJWT.PUT("/orders/:id", controllers.UpdateOrderByIdController)
	eJWT.DELETE("/orders/:id", controllers.DeleteOrderByIdController)

	eJWT.GET("/orderDetails", controllers.GetOrderDetailsController)
	eJWT.GET("/orderDetails/:id", controllers.GetOrderDetailByIdController)
	eJWT.POST("/orderDetails", controllers.CreateOrderDetailController)
	eJWT.PUT("/orderDetails/:id", controllers.UpdateOrderDetailByIdController)
	eJWT.DELETE("/orderDetails/:id", controllers.DeleteOrderDetailByIdController)

	eJWT.GET("/paymentMethods", controllers.GetPaymentMethodsController)
	eJWT.GET("/paymentMethods/:id", controllers.GetPaymentMethodByIdController)
	eJWT.POST("/paymentMethods", controllers.CreatePaymentMethodController)
	eJWT.PUT("/paymentMethods/:id", controllers.UpdatePaymentMethodByIdController)
	eJWT.DELETE("/paymentMethods/:id", controllers.DeletePaymentMethodByIdController)

	e.GET("/products", controllers.GetProductsController)
	e.GET("/products/:id", controllers.GetProductByIdController)
	eJWT.POST("/products", controllers.CreateProductController)
	eJWT.PUT("/products/:id", controllers.UpdateProductByIdController)
	eJWT.DELETE("/products/:id", controllers.DeleteProductByIdController)

	e.POST("/register", controllers.RegisterCustomerController)

	return e
}
