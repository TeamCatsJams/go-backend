package routes

import (
	"telmed_backend/controllers"
	"telmed_backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func appointmentRouter(c fiber.Router) {
	c.Use(middlewares.CheckAuth)
	c.Post("", controllers.CreateAppointment)
	c.Get(":appointment_id", controllers.FetchAppointmentByID)
	c.Get("/user/:user_id", controllers.FetchAllAppointmentByUserID)
	c.Get("/doctor/:doctor_id", controllers.FetchAllAppointmentByDoctorID)
}
