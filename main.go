package main

// imported package that will use in this file
import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/zercle/gofiber-101/book"
	"github.com/zercle/gofiber-101/database"
	"github.com/zercle/gofiber-101/routers"
)

// function call for init database connection
// param: none
// return: error if occur
func initDatabase() (err error) {
	// open connecttion to database and store connection object pointer to database.DBConn and any error to err variable
	// https://gorm.io/docs/connecting_to_the_database.html
	database.DBConn, err = gorm.Open(sqlite.Open("store.db"), &gorm.Config{})
	// if err variable not nil show log message and exit app
	if err != nil {
		// show log message and exit app
		log.Fatal("failed to connect database")
	}
	// just show log message
	log.Println("Connection Opened to Database")

	// migrate table into database by model
	// https://gorm.io/docs/migration.html
	err = database.DBConn.AutoMigrate(&book.Book{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}
	log.Println("Database Migrated")
	return
}

// go will start here
func main() {
	// init Fiber's app with default config
	// https://docs.gofiber.io/api/fiber
	app := fiber.New()
	// basically we'll use fiber's app to handle request and response
	// https://docs.gofiber.io/api/app

	// all request into this app will handle by CORS middleware before to next handler
	// https://docs.gofiber.io/api/middleware
	app.Use(cors.New())

	// call initDatabase function
	initDatabase()

	// serve static files from directory `./public` to request / path
	// https://docs.gofiber.io/api/app#static
	app.Static("/", "./public")

	// parse app's pointer
	routers.SetupRoutes(app)

	// app listen to port 3000/tcp (http://localhost:3000) if error app will exit by log.Fatal
	log.Fatal(app.Listen(":3000"))
}
