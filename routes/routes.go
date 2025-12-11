package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sourabh-sharma3435/rest-inventory/handlers"
	"github.com/sourabh-sharma3435/rest-inventory/middlewares"
)

func SetupRoutes(app *fiber.App) {

	var publicRoutes fiber.Router = app.Group("/api/v1")

	publicRoutes.Post("/signup", handlers.Signup)
	publicRoutes.Post("/login", handlers.Login)
	publicRoutes.Get("/items", handlers.GetAllItems)
	publicRoutes.Get("/items/:id", handlers.GetItemByID)

	var privateRoutes fiber.Router = app.Group("/api/v1", middlewares.CreateMiddleware())

	privateRoutes.Post("/items", handlers.CreateItem)
	privateRoutes.Put("/items/:id", handlers.UpdateItem)
	privateRoutes.Delete("/items/:id", handlers.DeleteItem)
}
