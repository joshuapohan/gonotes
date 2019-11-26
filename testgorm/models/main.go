package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func OpenDB() {
	var err error
	dbURL := "host=" + "localhost" +
		" user=" + "postgres" +
		" dbname=testgorm" + " sslmode=" + "disable" +
		" password=" + "dijital2012"

	DB, err = gorm.Open(
		"postgres",
		dbURL,
	)

	fmt.Println(dbURL)

	DB.AutoMigrate(&User{})

	if err != nil {
		panic(err)
	}
}

func CloseDB() {
	DB.Close()
}
