package controller

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "Test_Go_Monica/models"
	"Test_Go_Monica/db"
)

func GetAllProducts(c echo.Context) error {
    // koneksi ke DB
	db, err := db.ConnectDB()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal koneksi ke database"})
    }
    defer db.Close()

    // Mengambil daftar semua produk dari database
    products, err := models.GetAllProducts()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil list produk"})
    }

    return c.JSON(http.StatusOK, products)
}


// func GetProductByID(c echo.Context) error {
//     // Your code to fetch a book by ID from the database and return as JSON
// }


// func CreateProduct(c echo.Context) error {
// 	//koneksi ke db
// 	db, err := db.ConnectDB()
//     if err != nil {
//         return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal koneksi ke database"})
//     }
//     defer db.Close()

    
//     return c.JSON(http.StatusCreated, product)
// }
