package controller

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "Test_Go_Monica/models"
	"strconv"
)

func GetAllProducts(c echo.Context) error {
    
    // Mengambil daftar semua produk dari database
    products, err := models.GetAllProducts()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
    }

    return c.JSON(http.StatusOK, products)
}



func AddProduct(c echo.Context) error {
	// Tampung parameter yang diinputkan pada POST form
	Nama_product := c.FormValue("product_name")
	harga := c.FormValue("price")
	qty := c.FormValue("qty")
	brand_id := c.FormValue("brand_id")

	result, err := models.AddProduct(Nama_product, harga,qty, brand_id)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK,result)
}


func UpdateProduct(c echo.Context) error{ 
	// Variable untuk menerima inputan method POST
	Id_author := c.FormValue("product_id")
	Nama_author := c.FormValue("product_name")
	harga := c.FormValue("price")
	// qty := c.FormValue("qty")
	// Brand_Id := c.FormValue("brand_id")

	// Konversi id string to int
	conv_id, err := strconv.Atoi(Id_author)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error()) // tampilkan message error jika error
	}

	result, err := models.UpdateProduct(conv_id, Nama_author, harga) 
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // Jika saat parsing terjadi error maka tampilkan errornya
	}

	return c.JSON(http.StatusOK, result) // Jika sukses maka berikan status 200
	
}

func DeleteProduct(c echo.Context) error{
	// Variable untuk menerima inputan dari method POST
	Id_product := c.FormValue("product_id")

	// Konversi id string to int
	conv_id, err := strconv.Atoi(Id_product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) 
	}

	result, err := models.DeleteProduct(conv_id) 
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error()) // cek apakah saat parsing terjadi error
	}

	return c.JSON(http.StatusOK, result) // jika semua sukses, maka set status menjadi 200 
}