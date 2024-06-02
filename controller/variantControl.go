package controller

import (
	"CRUD_GoSQL/entities"
	"database/sql"
	"fmt"
	"time"
)

var variant = entities.Variant{}

func CreateVariant(db *sql.DB) {
	result, err := db.Exec(`insert into variant (variant_name, quantity, product_id, created_at, updated_at) values (?,?,?,?,?)`,
		"Samsung A23", 999, 3, time.Now(), time.Now())

	if err != nil {
		panic(err)
	}

	lastInsertID, _ := result.LastInsertId()

	sqlRetrieve := `select * from variant where id = ?`

	err = db.QueryRow(sqlRetrieve, lastInsertID).Scan(&variant.ID, &variant.Variant_name, &variant.Quantity, &variant.Product_id, &variant.Created_at, &variant.Updated_at)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New Product Data : %+v\n", variant)
}

func UpdateVariantById(db *sql.DB) {
	variant.ID = 1
	result, err := db.Exec(`update variant set variant_name = ?, quantity = ?, updated_at = ? where id = ?`,
		"Redmi Note 9", 666, time.Now(), variant.ID)

	if result != nil {
		fmt.Println("Update Data Variant Berhasil!")
	} else {
		panic(err)
	}

	row, err := db.Query(`select * from variant where id = ?`, variant.ID)
	if err != nil {
		panic(err)
	}

	defer row.Close()
	for row.Next() {
		if err := row.Scan(&variant.ID, &variant.Variant_name, &variant.Quantity, &variant.Product_id, &variant.Created_at, &variant.Updated_at); err != nil {
			panic(err)
		}
	}

	fmt.Println("Data Variant Terupdate")
	fmt.Println("ID :", variant.ID)
	fmt.Println("Name :", variant.Variant_name)
	fmt.Println("Quantity :", variant.Quantity, "buah")
	fmt.Println("Product_id :", variant.Product_id)
	fmt.Println("Created At :", variant.Created_at)
	fmt.Println("Updated At :", variant.Updated_at)
}

func DeleteVariantById(db *sql.DB) {
	variant.ID = 2
	result, err := db.Exec(`delete from variant where id = ?`, variant.ID)
	if result != nil {
		fmt.Println("Variant dengan ID", variant.ID, "Terhapus")
	} else {
		panic(err)
	}
}
