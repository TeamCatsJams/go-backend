package appointment

import (
	"telmed_backend/models"
)

type Repository interface {
	CreateAppointment(appointment *models.Appointment) (*models.Appointment, error)
	FetchAppointmenByID(appointmentId string) (*models.Appointment, error)
	FetchAllAppointmentByUserID(userId string) ([]models.Appointment, error)
	FetchAllAppointmentByDoctorID(doctorId string) ([]models.Appointment, error)
}
