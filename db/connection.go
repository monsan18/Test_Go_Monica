package db

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
     db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test_golang_monica")
    if err != nil {
        return nil, err
    }

    // Memeriksa koneksi ke database
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    fmt.Println("Berhasil terkoneksi ke database.")
    return db, nil
}