package datasources

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// function call for init database connection
// param: none
// return: error if occur
func InitSqLite() (dbConn *gorm.DB, err error) {
	// open connecttion to database and store connection object pointer to database.DBConn and any error to err variable
	// https://gorm.io/docs/connecting_to_the_database.html
	dbConn, err = gorm.Open(sqlite.Open("store.db"), &gorm.Config{})
	// if err variable not nil show log message and exit app
	if err != nil {
		// show log message and exit app
		log.Fatal("failed to connect database")
	}
	// just show log message
	log.Println("Connection Opened to Database")
	return
}
