// file to intercat with the database
package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//variable to connect to database
var db*gorm.DB

//Connect is a function to establish a connection to the database
//
//It opens a connection to a mysql database using the provided dsn (data source name)
//The DSN is a string with the following format: "username:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local"
//
//The function returns an error if it is unable to connect to the database
//
//Parameters:
//  None
//
//Returns:
//  An error if unable to connect to the database
func Connect() {
	//Open a connection to the database using the provided DSN
	d, err := gorm.Open("mysql", "Sam:Sam@12@/simple_rest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		//panic if unable to connect to the database
		panic(err)
	}
	//Assign the database connection to the global variable db
	db = d
}

//Retrieves the global database connection variable
//This function is useful when the database connection needs to be used in other files.
func GetDB() *gorm.DB {
	//Returns the global variable db, which is the database connection
	//This variable is created when the Connect function is called
	return db
}
