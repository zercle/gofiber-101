package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zercle/gofiber-101/internal/services"
)

// we can manage router outside main()
// by parse fiber's app pointer to other function
// that we can use app variable as same as inside main()
// https://docs.gofiber.io/guide/routing
func SetupRoutes(app *fiber.App) {
	// group request `your.web/api` to apiGroup variable
	apiGroup := app.Group("/api")

	// group request `your.web/api/v1` to v1Group variable
	v1Group := apiGroup.Group("/v1")
	{
		// GET request to `your.web/api/v1/book`
		// will process by book.GetBooks handler
		v1Group.Get("/book", services.GetBooks)
		// GET request to `your.web/api/v1/book/{id}`
		// will process by book.GetBook handler
		// `id`` will parse into fiber's context Params("id")
		// https://docs.gofiber.io/guide/routing#parameters
		v1Group.Get("/book/:id", services.GetBook)
		// POST request to `your.web/api/v1/book`
		// will process by book.NewBook handler
		v1Group.Post("/book", services.NewBook)
		// DELETE request to `your.web/api/v1/book/{id}`
		// will process by book.DeleteBook handler
		// `id`` will parse into fiber's context Params("id")
		v1Group.Delete("/book/:id", services.DeleteBook)
	}

	// group request `your.web/api/v2` to v2Group variable
	v2Group := apiGroup.Group("/v2")
	{
		v2Group.Get("/book", services.GetBooks)
		v2Group.Get("/book/:id", services.GetBook)
		v2Group.Post("/book", services.NewBook)
		v2Group.Delete("/book/:id", services.DeleteBook)
	}
}
