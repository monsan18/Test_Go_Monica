package main

import (
	
	"github.com/labstack/echo/v4"
	"Test_Go_Monica/db"
	"Test_Go_Monica/controller"
)

func main() {
	//koneksi ke DB
	db, err := db.ConnectDB()
    if err != nil {
        panic("connectionString error")
    }
    defer db.Close()


	//membuat instance echo
	e := echo.New()

	
	e.GET("/products", controller.GetAllProducts)
	//e.POST("/products", controller.CreateProduct)

	e.Logger.Fatal(e.Start(":8080"))
}