package config

/**
 * @file app.go
 * @author Original author: Free Code Camp
 *		   Changes made by 0xChristopher for learning purposes
 *
 * This file attempts to connect to a database using Gorm, and can then return that connection to handle
 * incoming API calls.
 */

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// The Connect() function connects to the MySQL database.
func Connect() {
	d, err := gorm.Open("mysql", "ChristopherTest:Test@tcp(127.0.0.1:3306)/gobookstore?charset=utf8&parseTime=true&loc=Local")

	if err != nil {
		panic(err)
	}

	db = d
}

// The GetDB() function returns the database instance.
func GetDB() *gorm.DB {
	return db
}
