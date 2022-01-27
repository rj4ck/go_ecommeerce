package main

import (
	"github.com/gofiber/fiber/v2"
	"server/src/configs"
	"server/src/routes"
)

func main() {

	ServerPort := configs.GetEnv("PORT")

	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.Categories(app)

	app.Listen(":" + ServerPort)
}
