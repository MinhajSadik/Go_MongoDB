package Routes

import (
	"github.com/MinhajSadik/Go-MongoDB/Controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthenticationRoutes(route fiber.Router) {
	route.Post("/register", Controllers.Register)

}
