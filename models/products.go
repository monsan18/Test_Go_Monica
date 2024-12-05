package models

import (
	"Test_Go_Monica/db"
	"net/http"
)

type Products struct {
	ProductId   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Qty         int     `json:"qty"`
	BrandId     int     `json:"brand_id"`
}

func GetAllProducts() (Response, error) {

	// inisialisasi variable yang akan digunakan
	var obj Products
	var arrobj []Products
	var res Response

	con := db.CreateCon() // Inisialisasi awal koneksi model terhadap database

	sqlStatement := "SELECT * FROM products" // query menampilkan seluruh data yang ada pada tabel product

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err // error akan di handle oleh controller
	}

	for rows.Next() {
		err = rows.Scan(&obj.ProductId, &obj.ProductName, &obj.Price, &obj.Qty, &obj.BrandId) // Scan tiap kolom apakah ada error
		if err != nil {
			return res, err // error akan di handle oleh controller
		}

		arrobj = append(arrobj, obj) // Menampilkan obj yang telah di scan error persatunya dalam bentuk array
	}

	res.Status = http.StatusOK // set status sbg 200 (ok)
	res.Message = "Success"    // set pesan berhasil
	res.Data = arrobj          // set data yang ditampilkan pada json dalam bentuk arr yang telah lolos cek error

	return res, nil
}

func AddProduct(Product_Name string, Price string, Qty string, Brand_Id string) (Response, error) { // Fungsi untuk meng-inputkan data Post Method dari tabel product
	var res Response // variable untuk menampung response nya

	con := db.CreateCon() // inisialisasi koneksi database

	sqlStatement := "INSERT INTO products (product_name,price, qty, brand_id) VALUES (?,?,?,?)" // query insert data

	stmt, err := con.Prepare(sqlStatement) // Variable persiapan awal sebelum eksekusi insert data

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Product_Name, Price, Qty, Brand_Id) // Eksekusi insert data query

	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"Id yang terakhir dimasukan adalah": lastInsertedId,
	}

	return res, nil

}

func UpdateProduct(Id_product int, Product_Name string, Price string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE products SET product_name = ? , price = ? WHERE product_id = ?" // Query update

	stmt, err := con.Prepare(sqlStatement) // Variable persiapan awal sebelum eksekusi update data

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Product_Name, Price, Id_product) // eksekusi dari query

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK // set status OK karena berhasil
	res.Message = "Success"
	res.Data = map[string]int64{
		"Data terupdate sebanyak": rowsAffected, // tampilkan data banyaknya rows yang terkena pembaruan data
	}

	return res, nil
}

func DeleteProduct(Id_product int) (Response, error) {
	var res Response

	con := db.CreateCon() // Inisialisasi database

	sqlStatement := "DELETE FROM products WHERE product_id = ?" // Query hapus data tablenya

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Id_product)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected() // Meminta return dari berapa banyak rows yang terdampak kena penghapusan data
	if err != nil {
		return res, err // cek apakah saat penghapusan dari rows yang kena dampak ada error
	}

	// return sukses jika berhasil
	res.Status = http.StatusOK
	res.Message = "Data berhasil dihapus"
	res.Data = map[string]int64{
		"Data terhapus sebanyak": rowsAffected, // tampilkan berapa banyak row yang kena dampak
	}

	return res, nil
}

func GetProductUsingPagination(page, pageSize int) ([]Products, error) {
	con := db.CreateCon() // Inisialisasi database

	offset := (page - 1) * pageSize
	rows, err := con.Query("SELECT product_id, product_name, price, qty, brand_id FROM products LIMIT ? OFFSET ?", pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Products
	for rows.Next() {
		var product Products
		if err := rows.Scan(&product.ProductId, &product.ProductName, &product.Price, &product.Qty, &product.BrandId); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
