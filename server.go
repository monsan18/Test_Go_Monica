package main

import (
	
	"github.com/labstack/echo/v4"
	"Test_Go_Monica/db"
	"Test_Go_Monica/controller"
)

func main() {
	//koneksi ke DB
	db.Init()


	//membuat instance echo
	e := echo.New()

	
	e.GET("/products", controller.GetAllProducts)
	e.POST("/products", controller.AddProduct)
	e.PUT("/products", controller.UpdateProduct)
	e.DELETE("/products", controller.DeleteProduct)

	e.Logger.Fatal(e.Start(":8080"))
}
