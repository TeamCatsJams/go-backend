package controllers

import (
	"telmed_backend/db"
	"telmed_backend/models"
	"telmed_backend/views"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUser(c *fiber.Ctx) error {
	userEmail := c.Locals("userEmail").(string)
	userName := c.Locals("userName").(string)
	userAvatar := c.Locals("userAvatar").(string)
	data, err := db.UserSvc.FetchProfileByEmail(userEmail)
	if err != nil {
		//TODO: fix this!
		if err == gorm.ErrRecordNotFound {
			//insert data into db
			user, err := db.UserSvc.CreateUser(&models.User{Email: userEmail, FullName: userName, Avatar: userAvatar})
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
