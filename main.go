package main

import (
	"CRUD_GoSQL/config"
	"CRUD_GoSQL/controller"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := config.DbConnect()
	if err != nil {
		panic(err)
	}

	var choose int
	fmt.Printf("Pilih satu (1-8) = ")
	fmt.Scanln(&choose)

	switch choose {
	case 1:
		controller.CreateProduct(db)
	case 2:
		controller.UpdateProduct(db)
	case 3:
		controller.GetProductById(db)
	case 4:
		controller.CreateVariant(db)
	case 5:
		controller.UpdateVariantById(db)
	case 6:
		controller.UpdateVariantById(db)
	case 7:
		controller.GetProductWithVariant(db)
	default:
		fmt.Println("Inputan salah, masukkan nomor 1-8")
	}
}
