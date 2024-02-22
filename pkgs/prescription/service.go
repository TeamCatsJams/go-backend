package prescription

import "telmed_backend/models"

type Service interface {
	CreatePrescription(prescription *models.Prescription) (*models.Prescription, error)
	FetchPrescriptionByID(prescriptionId string) (*models.Prescription, error)
	FetchAllPrescriptionByUserID(userId string) ([]models.Prescription, error)
	FetchAllPrescriptionByDoctorID(doctorId string) ([]models.Prescription, error)
}

type prescriptionSvc struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &prescriptionSvc{repo: r}
}

// CreatePrescription implements Service.
func (s *prescriptionSvc) CreatePrescription(prescription *models.Prescription) (*models.Prescription, error) {
	return s.repo.CreatePrescription(prescription)
}

// FetchAllPrescriptionByDoctorID implements Service.
func (s *prescriptionSvc) FetchAllPrescriptionByDoctorID(doctorId string) ([]models.Prescription, error) {
	return s.repo.FetchAllPrescriptionByDoctorID(doctorId)
}

// FetchAllPrescriptionByUserID implements Service.
func (s *prescriptionSvc) FetchAllPrescriptionByUserID(userId string) ([]models.Prescription, error) {
	return s.repo.FetchAllPrescriptionByUserID(userId)
}

// FetchPrescriptionByID implements Service.
func (s *prescriptionSvc) FetchPrescriptionByID(prescriptionId string) (*models.Prescription, error) {
	return s.repo.FetchPrescriptionByID(prescriptionId)
}
