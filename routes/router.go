package routes

import (
	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {
	app.Get("/api", func(c *fiber.Ctx) error { return c.SendString("Hello world, routes mounted!!") }) // logged in browers /api,if MountRoutes run
	//api routeObject group
	api := app.Group("/api") // api group

	//version group
	v1 := api.Group("/v1") // /api/v1 group

	//main routes
	_ = v1.Group("/user") //api/v1/user

	doctorRoute := v1.Group("/doctor")
	doctorRouter(doctorRoute)

	appointmentRoute := v1.Group("/appointment")
	appointmentRouter(appointmentRoute)

	prescriptionRoute := v1.Group("/prescription")
	prescriptionRouter(prescriptionRoute)
}
