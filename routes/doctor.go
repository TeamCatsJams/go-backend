package routes

import (
	"telmed_backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func doctorRouter(c fiber.Router) {
	c.Post("", controllers.CreateDoctor)
	c.Get("", controllers.FetchAllDoctor)
	c.Get("/:doctor_id", controllers.FetchDoctor)

}
