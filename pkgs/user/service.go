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
	return s.repo.FetchProfileByEmail(email)
}

func (s *userSvc) CreateUser(user *models.User) (*models.User, error) {
	return s.repo.CreateUser(user)
}
func (s *userSvc) FetchOrCreateUser(user *models.User) (*models.User, error) {
	return s.repo.FetchOrCreateUser(user)
}

func (s *userSvc) FetchProfileById(id string) (*models.User, error) {
	return s.repo.FetchProfileById(id)
}

func NewService(r Repository) Service {
	return &userSvc{repo: r}
}
