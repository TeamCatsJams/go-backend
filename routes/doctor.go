package routes

import (
	"telmed_backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func doctorRouter(c fiber.Router) {
	c.Get("", controllers.FetchAllDoctor)
	c.Get("/:doctor_id", controllers.FetchDoctor)
	c.Post("", controllers.CreateDoctor)
}
