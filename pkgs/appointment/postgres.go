package appointment

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

func (r *repo) CreateAppointment(appointment *models.Appointment) (*models.Appointment, error) {
	if err := r.DB.Create(appointment).Error; err != nil {
		return nil, err
	}
	return appointment, nil
}

func (r *repo) FetchAppointmenByID(appointmentId string) (*models.Appointment, error) {
	appointment := &models.Appointment{}
	if err := r.DB.First(appointment, "id = ?", appointmentId).Error; err != nil {
		return nil, err
	}
	return appointment, nil
}

func (r *repo) FetchAllAppointmentByUserID(userId string) ([]models.Appointment, error) {
	appointments := make([]models.Appointment, 0)
	if err := r.DB.Find(appointments, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *repo) FetchAllAppointmentByDoctorID(doctorId string) ([]models.Appointment, error) {
	appointments := make([]models.Appointment, 0)
	if err := r.DB.Find(appointments, "doctor_id = ?", doctorId).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}
