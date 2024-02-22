package routes

import (
	"telmed_backend/controllers"
	"telmed_backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func prescriptionRouter(c fiber.Router) {
	c.Use(middlewares.CheckAuth)
	c.Post("", controllers.CreatePrescription)
	c.Get(":prescription_id", controllers.FetchAppointmentByID)
	c.Get("/user/:user_id", controllers.FetchAllPrescriptionByUserID)
	c.Get("/doctor/:doctor_id", controllers.FetchAllPrescriptionByDoctorID)
}
