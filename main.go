package main

import (
	"log"
	"pretest-golang-tdi/config"
	"pretest-golang-tdi/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.ConnectDB()
	defer config.DB.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:4200",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
