package user

import (
	"telmed_backend/models"
	"telmed_backend/pkgs/errors"

	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

// CreateUser implements Repository.
func (r *repo) CreateUser(user *models.User) (*models.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// FetchOrCreateUser implements Repository.
func (r *repo) FetchOrCreateUser(user *models.User) (*models.User, error) {
	err := r.DB.FirstOrCreate(user, "email = ?", user.Email).Error
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrRecordNotFound
		}
		return nil, err
	}
	return user, nil
}

// FetchProfileByEmail implements Repository.
func (r *repo) FetchProfileByEmail(email string) (*models.User, error) {
	users := &models.User{}
	err := r.DB.First(users, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// FetchProfileById implements Repository.
func (r *repo) FetchProfileById(id string) (*models.User, error) {
	user := &models.User{}
	if err := r.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func NewPostgresRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}
