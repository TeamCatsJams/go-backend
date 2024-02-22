package prescription

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

// CreatePrescription implements Repository.
func (r *repo) CreatePrescription(prescription *models.Prescription) (*models.Prescription, error) {
	if err := r.DB.Create(prescription).Error; err != nil {
		return nil, err
	}
	return prescription, nil
}

// FetchAllPrescriptionByDoctorID implements Repository.
func (r *repo) FetchAllPrescriptionByDoctorID(doctorId string) ([]models.Prescription, error) {
	prescriptions := make([]models.Prescription, 0)
	if err := r.DB.Find(prescriptions, "doctor_id = ?", doctorId).Error; err != nil {
		return nil, err
	}
	return prescriptions, nil
}

// FetchAllPrescriptionByUserID implements Repository.
func (r *repo) FetchAllPrescriptionByUserID(userId string) ([]models.Prescription, error) {
	prescriptions := make([]models.Prescription, 0)
	if err := r.DB.Find(prescriptions, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}
	return prescriptions, nil
}

// FetchPrescriptionByID implements Repository.
func (r *repo) FetchPrescriptionByID(prescriptionId string) (*models.Prescription, error) {
	prescription := &models.Prescription{}
	if err := r.DB.First(prescription, "id = ?", prescriptionId).Error; err != nil {
		return nil, err
	}
	return prescription, nil
}
