package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Hace la conexion a la BD
func GetConnectionBd() (*gorm.DB, error) {

	db, err := gorm.Open("mysql", "root:jhon0128@/mydb")

	if err != nil {
		return nil, err
	} else {
		return db, nil
	}
}

func ChequeoConnectionDb() int {
	db, err := GetConnectionBd()

	if err != nil {
		fmt.Printf("Error obteniendo base de dato: %v", err)
	}

	defer db.Close()

	err = db.DB().Ping()
	if err != nil {
		return 0
	}

	return 1
}
