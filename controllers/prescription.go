package controllers

import (
	"telmed_backend/db"
	"telmed_backend/models"
	"telmed_backend/schemas"
	"telmed_backend/views"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreatePrescription(c *fiber.Ctx) error {
	var data schemas.PrescriptionData
	userEmail := c.Locals("userEmail").(string)
	if err := c.BodyParser(&data); err != nil {
		return views.BadRequest(c)
	}

	user, err := db.UserSvc.FetchProfileByEmail(userEmail)
	if err != nil {
		return views.InternalServerError(c, err)
	}
	Prescription := models.Prescription{
		Title:    data.Title,
		Document: data.Document,
		DoctorID: data.DoctorID,
		UserID:   &user.ID,
	}
	newPrescription, err := db.PrescriptionSvc.CreatePrescription(&Prescription)
	if err != nil {
		return views.InternalServerError(c, err)
	}
	return views.StatusOK(c, newPrescription)
}

func FetchAllPrescriptionByDoctorID(c *fiber.Ctx) error {
	doctorId, err := uuid.Parse(c.Params("doctor_id"))
	if err != nil {
		return views.BadRequest(c)
	}
	Prescriptions, err := db.PrescriptionSvc.FetchAllPrescriptionByDoctorID(doctorId.String())
	if err != nil {
		return views.InternalServerError(c, err)
	}
	return views.StatusOK(c, Prescriptions)
}

func FetchPrescriptionByID(c *fiber.Ctx) error {
	prescriptionId, err := uuid.Parse(c.Params("Prescription_id"))
	if err != nil {
		return views.BadRequest(c)
	}
	Prescription, err := db.PrescriptionSvc.FetchPrescriptionByID(prescriptionId.String())
	if err != nil {
		return views.InternalServerError(c, err)
	}
	return views.StatusOK(c, Prescription)
}

func FetchAllPrescriptionByUserID(c *fiber.Ctx) error {
	userEmail := c.Locals("userEmail").(string)
	user, err := db.UserSvc.FetchProfileByEmail(userEmail)
	if err != nil {
		return views.InternalServerError(c, err)
	}
	Prescriptions, err := db.PrescriptionSvc.FetchAllPrescriptionByUserID(user.ID.String())
	if err != nil {
		return views.InternalServerError(c, err)
	}
	return views.StatusOK(c, Prescriptions)
}
