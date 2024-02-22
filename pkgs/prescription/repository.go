package prescription

import "telmed_backend/models"

type Repository interface {
	CreatePrescription(prescription *models.Prescription) (*models.Prescription, error)
	FetchPrescriptionByID(prescriptionId string) (*models.Prescription, error)
	FetchAllPrescriptionByUserID(userId string) ([]models.Prescription, error)
	FetchAllPrescriptionByDoctorID(doctorId string) ([]models.Prescription, error)
}
