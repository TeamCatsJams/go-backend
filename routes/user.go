package routes

import (
	"telmed_backend/controllers"
	"telmed_backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func userRouter(c fiber.Router) {
	c.Use(middlewares.CheckAuth)
	c.Get("", controllers.GetUser)
}
