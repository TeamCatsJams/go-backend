package user

import "telmed_backend/models"

type Repository interface {
	FetchProfileByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	FetchOrCreateUser(user *models.User) (*models.User, error)
	FetchProfileById(id string) (*models.User, error)
}
