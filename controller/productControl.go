package controller

import (
	"CRUD_GoSQL/entities"
	"database/sql"
	"fmt"
	"time"
)

var product = entities.Product{}

func CreateProduct(db *sql.DB) {
	result, err := db.Exec(`insert into product (name, created_at, updated_at) values (?,?,?)`,
		"Samsung", time.Now(), time.Now())

	if err != nil {
		panic(err)
	}

	lastInsertID, _ := result.LastInsertId()

	sqlRetrieve := `select * from product where id_product = ?`

	err = db.QueryRow(sqlRetrieve, lastInsertID).Scan(&product.ID_Product, &product.Name, &product.Created_at, &product.Updated_at)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New Product Data : %+v\n", product)
}

func UpdateProduct(db *sql.DB) {
	product.ID_Product = 5
	result, err := db.Exec(`update product set name = ?, updated_at = ? where id_product = ?`,
		"Infinix", time.Now(), product.ID_Product)
	if err != nil {
		panic(err)
	}

	results, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	if results == 0 {
		fmt.Println("DATA TIDAK ADA")
	} else {
		fmt.Println("Update Data Product Berhasil!")
		row, err := db.Query(`select * from product where id_product = ?`, product.ID_Product)
		if err != nil {
			panic(err)
		}

		defer row.Close()
		for row.Next() {
			if err := row.Scan(&product.ID_Product, &product.Name, &product.Created_at, &product.Updated_at); err != nil {
				panic(err)
			}
		}
		fmt.Println("Data Product Terupdate")
		fmt.Println("ID :", product.ID_Product)
		fmt.Println("Name :", product.Name)
		fmt.Println("Created At :", product.Created_at)
		fmt.Println("Updated At :", product.Updated_at)
	}
}

func GetProductById(db *sql.DB) {
	product.ID_Product = 4
	row, err := db.Query(`select * from product where id_product = ?`, product.ID_Product)
	if err != nil {
		panic(err)
	}

	defer row.Close()
	for row.Next() {
		if err := row.Scan(&product.ID_Product, &product.Name, &product.Created_at, &product.Updated_at); err != nil {
			panic(err)
		}
	}

	fmt.Println("Product by ID :", product.ID_Product)
	fmt.Println("Product Name :", product.Name)
	fmt.Println("Created At :", product.Created_at)
	fmt.Println("Updated At :", product.Updated_at)
}

func getProduct(db *sql.DB) {
	var products = []entities.Product{}
	row, err := db.Query("select * from product")
	if err != nil {
		panic(err)
	}

	defer row.Close()

	var i int
	fmt.Println("Data Product")
	for row.Next() {
		err = row.Scan(&product.ID_Product, &product.Name, &product.Created_at, &product.Updated_at)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
		fmt.Println("Product ID:", products[i].ID_Product)
		fmt.Println("Product Name :", products[i].Name)
		fmt.Println("-------------------------------")
		i++
	}
}

func GetProductWithVariant(db *sql.DB) {
	getProduct(db)
	type newVariant struct {
		Name     string
		Quantity int
	}

	var newVariantt = newVariant{}
	var newVariants = []newVariant{}
	// var variants = []entities.Variant{}

	var choose int
	var pilih string
	var j int

	fmt.Println("Cari variant dengan?")
	fmt.Println("1. Dengan nama product")
	fmt.Println("2. Dengan ID product")
	fmt.Printf("Pilih ( 1 / 2 ) = ")
	fmt.Scanln(&choose)

	switch choose {
	case 1:
		fmt.Printf("Variant berdasarkan Nama Product apa? = ")
		fmt.Scanln(&pilih)
		rowSearch, err := db.Query(`select variant.variant_name, quantity from variant join product on variant.product_id = product.id_product where product.name = ?`,
			pilih)
		if err != nil {
			panic(err)
		}

		defer rowSearch.Close()

		fmt.Println("Data Variant berdasarkan product", pilih)
		for rowSearch.Next() {
			err = rowSearch.Scan(&newVariantt.Name, &newVariantt.Quantity)
			if err != nil {
				panic(err)
			}

			newVariants := append(newVariants, newVariantt)

			fmt.Println("Nama Variant :", newVariants[j].Name)
			fmt.Println("Quantity :", newVariants[j].Quantity)
			fmt.Println("-------------------------------")
		}

	case 2:
		fmt.Printf("Variant berdasarkan ID Product apa = ")
		fmt.Scanln(&choose)
		rowSearch, err := db.Query(`select variant.variant_name, quantity from variant join product on variant.product_id = product.id_product where product.id_product = ?`, choose)
		if err != nil {
			panic(err)
		}

		defer rowSearch.Close()

		fmt.Println("Data Variant berdasarkan ID Product", choose)
		for rowSearch.Next() {
			err = rowSearch.Scan(&newVariantt.Name, &newVariantt.Quantity)
			if err != nil {
				panic(err)
			}

			newVariants := append(newVariants, newVariantt)

			fmt.Println("Nama Variant :", newVariants[j].Name)
			fmt.Println("Quantity :", newVariants[j].Quantity)
			fmt.Println("-------------------------------")
		}

	default:
		fmt.Println("Salah Inputan")
	}

}
