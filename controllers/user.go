package controllers

import (
	"errors"
	"telmed_backend/db"
	"telmed_backend/views"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUser(c *fiber.Ctx) error {

	userEmail := c.Locals("userEmail").(string)
	data, err := db.UserSvc.FetchProfileByEmail(userEmail)
	if err != nil {
		//TODO: fix this!
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err != nil {

				return views.InternalServerError(c, err)
			}
			//insert data into db
			user, err := db.UserSvc.CreateUser(data)
			if err != nil {
				return views.InternalServerError(c, err)
			}
			data = user
		} else {
			return views.InternalServerError(c, err)
		}

	}
	return views.StatusOK(c, data)
}
