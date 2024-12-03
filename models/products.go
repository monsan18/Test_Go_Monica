package	models

import (
	"Test_Go_Monica/db"
	"net/http"
	
)

type Products struct {
	ProductId	int `json:"product_id"`
	ProductName	string	`json:"product_name"`
	Price	float64	`json:"price"`
	Qty	int	`json:"qty"`
	BrandId	int	`json:"brand_id"`
}

func GetAllProducts() (Response, error) { // Fungsi untuk mendapatkan Get Methods dari tabel author
	var obj Products
	var arrobj []Products
	var res Response
	// inisialisasi variable yang akan digunakan

	con :=  db.CreateCon() // Inisialisasi awal koneksi model terhadap database

	sqlStatement := "SELECT * FROM products" // query menampilkan seluruh data yang ada pada database iLib pada tabel author

	rows, err := con.Query(sqlStatement) // eksekusi query

	defer rows.Close()                   // tutup koneksi thdp database

	if err != nil {
		return res, err // error akan di handle oleh controller
	}

	for rows.Next() {
		err = rows.Scan(&obj.ProductId, &obj.ProductName) // Scan tiap kolom apakah ada error
		if err != nil {
			return res, err // error akan di handle oleh controller
		}

		arrobj = append(arrobj, obj) // Menampilkan obj yang telah di scan error persatunya dalam bentuk array
	}

	res.Status = http.StatusOK // set status sbg 200 (ok) karena telah dipilah errornya
	res.Message = "Success"    // set pesan berhasil
	res.Data = arrobj          // set data yang ditampilkan pada json dalam bentuk arr yang telah lolos cek error

	return res, nil
}