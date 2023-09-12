package main

import (
	"awesomeProject6/api/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.Createroutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
