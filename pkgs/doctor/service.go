package doctor

import "telmed_backend/models"

type Service interface {
	FetchProfileByEmail(email string) (*models.Doctor, error)
	CreateDoctor(Doctor *models.Doctor) (*models.Doctor, error)
	FetchProfileById(id string) (*models.Doctor, error)
	FetchAllDoctor() ([]models.Doctor, error)
}

type doctorSvc struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &doctorSvc{repo: r}
}

// CreateDoctor implements Service.
func (s *doctorSvc) CreateDoctor(doctor *models.Doctor) (*models.Doctor, error) {
	return s.repo.CreateDoctor(doctor)
}

// FetchProfileByEmail implements Service.
func (s *doctorSvc) FetchProfileByEmail(email string) (*models.Doctor, error) {
	return s.repo.FetchProfileByEmail(email)
}

// FetchProfileById implements Service.
func (s *doctorSvc) FetchProfileById(id string) (*models.Doctor, error) {
	return s.repo.FetchProfileById(id)
}

// FetchAllDoctor implements Service.
func (s *doctorSvc) FetchAllDoctor() ([]models.Doctor, error) {
	return s.repo.FetchAllDoctor()
}
