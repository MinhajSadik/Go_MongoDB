package main

import (
	"fmt"

	"github.com/MinhajSadik/Go-MongoDB/DB"
	"github.com/MinhajSadik/Go-MongoDB/Routes"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	Routes.AuthenticationRoutes(app.Group("/auth"))
}

func main() {
	app := fiber.New()

	initCollections := DB.InitCollections()

	if initCollections {
		fmt.Println("OK")
	} else {
		fmt.Println("ERROR")
	}

	SetupRoutes(app)

	fmt.Println("[OK]")
	app.Listen(":8080")
}
