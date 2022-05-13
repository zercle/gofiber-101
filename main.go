package main

// imported package that will use in this file
import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/zercle/gofiber-101/internal/datasources"
	"github.com/zercle/gofiber-101/internal/routes"
	"github.com/zercle/gofiber-101/pkg/models"
)

// go will start here
func main() {
	var err error
	// init Fiber's app with default config
	// https://docs.gofiber.io/api/fiber
	app := fiber.New()
	// basically we'll use fiber's app to handle request and response
	// https://docs.gofiber.io/api/app

	// all request into this app will handle by CORS middleware before to next handler
	// https://docs.gofiber.io/api/middleware
	app.Use(cors.New())

	// call initDatabase function
	datasources.DBConn, err = datasources.InitSqLite()
	if datasources.DBConn == nil || err != nil {
		log.Fatal("failed to connect database")
	}

	// migrate table into database by model
	// https://gorm.io/docs/migration.html
	err = datasources.DBConn.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}
	log.Println("Database Migrated")

	// serve static files from directory `./public` to request / path
	// https://docs.gofiber.io/api/app#static
	app.Static("/", "./web")

	// parse app's pointer
	routes.SetupRoutes(app)

	// app listen to port 3000/tcp (http://localhost:3000) if error app will exit by log.Fatal
	log.Fatal(app.Listen(":3000"))
}
