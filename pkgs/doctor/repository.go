package doctor

import "telmed_backend/models"

type Repository interface {
	FetchProfileByEmail(email string) (*models.Doctor, error)
	CreateDoctor(doctor *models.Doctor) (*models.Doctor, error)
	FetchProfileById(id string) (*models.Doctor, error)
	FetchAllDoctor() ([]models.Doctor, error)
}
