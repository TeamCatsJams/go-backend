package controllers

import (
	"telmed_backend/db"
	"telmed_backend/models"
	"telmed_backend/schemas"
	"telmed_backend/views"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateAppointment(c *fiber.Ctx) error {
	var data schemas.AppointmentData
	userEmail := c.Locals("userEmail").(string)
	if err := c.BodyParser(&data); err != nil {
		return views.BadRequest(c)
	}

	user, err := db.UserSvc.FetchProfileByEmail(userEmail)
	if err != nil {
		return views.InternalServerError(c, err)
	}
	appointment := models.Appointment{
		AppointmentAt: data.AppointmentAt,
		MeetLink:      data.MeetLink,
		DoctorID:      &data.DoctorID,
		UserID:        &user.ID,
	}
	newAppointment, err := db.AppointmentSvc.CreateAppointment(&appointment)
	if err != nil {
		return views.InternalServerError(c, err)
	}
	return views.StatusOK(c, newAppointment)
}

func FetchAllAppointmentByDoctorID(c *fiber.Ctx) error {
	doctorId, err := uuid.Parse(c.Params("doctor_id"))
	if err != nil {
		return views.BadRequest(c)
	}
	appointments, err := db.AppointmentSvc.FetchAllAppointmentByDoctorID(doctorId.String())
	if err != nil {
		return views.InternalServerError(c, err)
	}
	return views.StatusOK(c, appointments)
}

func FetchAppointmentByID(c *fiber.Ctx) error {
	appointmentId, err := uuid.Parse(c.Params("appointment_id"))
	if err != nil {
		return views.BadRequest(c)
	}
	appointment, err := db.AppointmentSvc.FetchAppointmenByID(appointmentId.String())
	if err != nil {
		return views.InternalServerError(c, err)
	}
	return views.StatusOK(c, appointment)
}

func FetchAllAppointmentByUserID(c *fiber.Ctx) error {
	userEmail := c.Locals("userEmail").(string)
	user, err := db.UserSvc.FetchProfileByEmail(userEmail)
	if err != nil {
		return views.InternalServerError(c, err)
	}
	appointments, err := db.AppointmentSvc.FetchAllAppointmentByUserID(user.ID.String())
	if err != nil {
		return views.InternalServerError(c, err)
	}
	return views.StatusOK(c, appointments)
}
