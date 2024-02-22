package controllers

import (
	"telmed_backend/schemas"
	"telmed_backend/views"

	"github.com/gofiber/fiber/v2"
)

func CreateDoctor(c *fiber.Ctx) error {
	var data schemas.DoctorData

	if err := c.BodyParser(&data); err != nil {
		return views.BadRequest(c)
	}

	db.
}
