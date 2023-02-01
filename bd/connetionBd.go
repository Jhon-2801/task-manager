package bd

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetConnetionBd() (*gorm.DB, error) {

	db, err := gorm.Open("mysql", "root:jhon0128@/mydb")

	if err != nil {
		return nil, err
	} else {
		log.Printf("ConectadoDB")
	}

	return db, nil
}

func ChequeoConnectioDb() int {
	db, err := GetConnetionBd()

	if err != nil {
		fmt.Printf("Error obteniendo base de datos: %v", err)
	}

	defer db.Close()

	err = db.DB().Ping()
	if err != nil {
		return 0
	}

	return 1
}
