package controllers

import (
	"telmed_backend/db"
	"telmed_backend/models"
	"telmed_backend/schemas"
	"telmed_backend/views"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateDoctor(c *fiber.Ctx) error {
	var data schemas.DoctorData

	if err := c.BodyParser(&data); err != nil {
		return views.BadRequest(c)
	}
	doctor := models.Doctor{
		FullName:        data.FullName,
		Language:        data.Language,
		Email:           data.Email,
		Phone:           data.Phone,
		MeetingDuration: data.MeetingDuration,
		Description:     data.Description,
		Country:         data.Country,
		Avatar:          data.Avatar,
		Practice:        data.Practice,
		Location:        data.Location,
		Designation:     data.Description,
	}
	newDoctor, err := db.DoctorSvc.CreateDoctor(&doctor)
	if err != nil {
		return views.InternalServerError(c, err)
	}
	return views.StatusOK(c, newDoctor)
}

func FetchAllDoctor(c *fiber.Ctx) error {

	doctors, err := db.DoctorSvc.FetchAllDoctor()
	if err != nil {
		return views.InternalServerError(c, err)
	}
	return views.StatusOK(c, doctors)
}

func FetchDoctor(c *fiber.Ctx) error {
	doctorId, err := uuid.Parse(c.Params("doctor_id"))
	if err != nil {
		return views.BadRequest(c)
	}
	doctor, err := db.DoctorSvc.FetchProfileById(doctorId.String())
	if err != nil {
		return views.InternalServerError(c, err)
	}
	return views.StatusOK(c, doctor)
}
