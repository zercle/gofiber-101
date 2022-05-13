package datasources

import "gorm.io/gorm"

var (
	// global variable pointer to gorm.DB database connection object
	DBConn *gorm.DB
)
