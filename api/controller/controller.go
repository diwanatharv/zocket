package controller

import (
	"awesomeProject6/api/handler"
	"github.com/labstack/echo/v4"
)

func Createroutes(e *echo.Echo) {
	e.PUT("/products/:id", handler.Updateproduct)
	e.POST("/users", handler.Createuser)
	e.POST("/products", handler.CreateProduct)
}
