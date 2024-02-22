package doctor

import (
	"telmed_backend/models"

	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

// CreateDoctor implements Repository.
func (r *repo) CreateDoctor(doctor *models.Doctor) (*models.Doctor, error) {
	if err := r.DB.Create(doctor).Error; err != nil {
		return nil, err
	}
	return doctor, nil
}

// FetchProfileByEmail implements Repository.
func (r *repo) FetchProfileByEmail(email string) (*models.Doctor, error) {
	doctors := &models.Doctor{}
	err := r.DB.First(doctors, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return doctors, nil
}

func (r *repo) FetchAllDoctor() ([]models.Doctor, error) {
	doctors := make([]models.Doctor, 0)
	err := r.DB.Find(&doctors).Error
	if err != nil {
		return nil, err
	}
	return doctors, nil
}

// FetchProfileById implements Repository.
func (r *repo) FetchProfileById(id string) (*models.Doctor, error) {
	doctor := &models.Doctor{}
	if err := r.DB.First(&doctor, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return doctor, nil
}
