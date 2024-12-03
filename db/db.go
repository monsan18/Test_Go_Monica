package db

import (
	"database/sql"
	"Test_Go_Monica/config"
	_"github.com/go-sql-driver/mysql"
)

 var db *sql.DB
 var err error

 func Init(){
	conf := config.GetConfig()

	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp("+ conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err = sql.Open("mysql", connectionString)

	if err != nil {
		panic("connectionString error") // cek apakah connection string ada salah input ga
	}

	err = db.Ping() // cek koneksi jalan  ada error ga
	if err != nil {
		panic("DSN invalid")
	}
 } // Fungsi Inisialisasi db awal


 func CreateCon() *sql.DB{
	 return db
 }