package main

import (
	"SQLFiberApi/database"
	"SQLFiberApi/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

var apiBaseUrl string = "/api/v1/users"

func welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello World!")
}

func setupRoutes(server *fiber.App) {
	server.Get("/", welcome)

	// ? User endpoints
	// Get method(s)
	server.Get(apiBaseUrl+"/users", routes.GetUsers)
	server.Get(apiBaseUrl+"/users/:id", routes.GetUserById)
	// Post method(s)
	server.Post(apiBaseUrl+"/users", routes.CreateUser)
	// Put method(s)
	server.Put(apiBaseUrl+"/users/:id", routes.UpdateUserById)
	// Delete method(s)
	server.Delete(apiBaseUrl+"/users/:id", routes.DeleteUserById)
}

func main() {
	// * Open SQLite DB file.
	database.ConnectDatabase()

	// ? Setup the server app.
	server := fiber.New()
	setupRoutes(server)

	log.Fatal(server.Listen(":8080"))
}
