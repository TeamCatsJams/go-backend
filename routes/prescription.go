package routes

import (
	"telmed_backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func prescriptionRouter(c fiber.Router) {
	c.Post("", controllers.CreatePrescription)
	c.Get(":prescription_id", controllers.FetchAppointmentByID)
	c.Get("/user/:user_id", controllers.FetchAllPrescriptionByUserID)
	c.Get("/doctor/:doctor_id", controllers.FetchAllPrescriptionByDoctorID)
}
