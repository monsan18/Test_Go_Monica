package	models

import (
	"Test_Go_Monica/db"
	"net/http"
	"errors"
)

type Brands struct  {
	BrandId	int `json:"brand_id"`
	BrandName	string	`json:"brand_name"`
}

func AddBrand(Brand_Name string) (Response, error){ // Fungsi untuk meng-inputkan data Post Method dari tabel product
	var res Response // variable untuk menampung response nya

	con := db.CreateCon() // inisialisasi koneksi database

	sqlStatement := "INSERT INTO brands (brand_name) VALUES (?)" // query insert data
 
	stmt, err := con.Prepare(sqlStatement) // Variable persiapan awal sebelum eksekusi insert data

	if err != nil {
		return res,err 
	}
	
	//cek jika brand_name kosong.
	if Brand_Name == "" {
        return res, errors.New("Brand Name tidak boleh kosong.")
    }

	result, err := stmt.Exec(Brand_Name) // Eksekusi insert data query


	if err != nil{
		return res, err 
	}

	lastInsertedId, err := result.LastInsertId()

	
	if err !=  nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"Id yang terakhir dimasukan adalah" : lastInsertedId,
	}

	return res, nil

}

func DeleteBrand(Brand_Id int) (Response, error){ 
	var res Response // variable untuk menampung response nya
	//var obj Brands

	con := db.CreateCon() // inisialisasi koneksi database

	//cek jika brand_id tdk diisi.
	if Brand_Id = nil {
		return res, errors.New("Brand ID tidak boleh kosong.")
	}

	sqlStatement2 := "DELETE FROM brands WHERE NOT EXISTS ( SELECT 1 FROM products WHERE brand_id = ?)" // Query hapus data jika brand_id tidak dipakai di product

	stmt, err := con.Prepare(sqlStatement2) 
	if err != nil {
		return res,err 
	}

	result, err := stmt.Exec(Brand_Id) 
	if err != nil {
		return res,err 
	}

	rowsAffected, err := result.RowsAffected() // Meminta return dari berapa banyak rows yang terdampak kena penghapusan data

	if rowsAffected == 0 {
		return res, errors.New("Tidak ada brand yang dihapus.")
	}

	if err != nil {
		return res,err // cek apakah saat penghapusan dari rows yang kena dampak ada error
	}

	// return sukses jika berhasil
	res.Status = http.StatusOK 
	res.Message = "Data berhasil dihapus" 
	res.Data = map[string]int64{
		"Data terhapus sebanyak" : rowsAffected, // tampilkan berapa banyak row yang kena dampak 
	}

	return res, nil

}