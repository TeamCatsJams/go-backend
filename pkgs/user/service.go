package user

import "telmed_backend/models"

type Service interface {
	FetchProfileByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	FetchOrCreateUser(user *models.User) (*models.User, error)
	FetchProfileById(id string) (*models.User, error)
}

type userSvc struct {
	repo Repository
}

func (s *userSvc) FetchProfileByEmail(email string) (*models.User, error) {
	panic("unimplemented")
}

func (s *userSvc) CreateUser(user *models.User) (*models.User, error) {
	panic("unimplemented")
}
func (s *userSvc) FetchOrCreateUser(user *models.User) (*models.User, error) {
	panic("unimplemented")
}

func (s *userSvc) FetchProfileById(id string) (*models.User, error) {
	panic("unimplemented")
}

func NewService(r Repository) Service {
	return &userSvc{repo: r}
}
