package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DbConnect() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPassword := ""
	dbName := "item_product"

	db, err = sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Succesfully connected to database")

	return
}
