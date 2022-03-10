package main

import (
	"database/sql"

	db2 "github.com/alancrist/nft-marktplace-service/adapters/db"
	"github.com/alancrist/nft-marktplace-service/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product", 30)

	productService.Enable(product)
}
