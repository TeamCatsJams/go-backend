package appointment

import "telmed_backend/models"

type Service interface {
	CreateAppointment(appointment *models.Appointment) (*models.Appointment, error)
	FetchAppointmenByID(appointmentId string) (*models.Appointment, error)
	FetchAllAppointmentByUserID(userId string) ([]models.Appointment, error)
	FetchAllAppointmentByDoctorID(doctorId string) ([]models.Appointment, error)
}

type appointmentSvc struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &appointmentSvc{repo: r}
}

// CreateAppointment implements Service.
func (s *appointmentSvc) CreateAppointment(appointment *models.Appointment) (*models.Appointment, error) {
	return s.repo.CreateAppointment(appointment)
}

// FetchAppointmenByID implements Service.
func (s *appointmentSvc) FetchAppointmenByID(appointmentId string) (*models.Appointment, error) {
	return s.repo.FetchAppointmenByID(appointmentId)
}

// FetchAllAppointmentByDoctorID implements Service.
func (s *appointmentSvc) FetchAllAppointmentByDoctorID(doctorId string) ([]models.Appointment, error) {
	return s.repo.FetchAllAppointmentByDoctorID(doctorId)
}

// FetchAllAppointmentByUserID implements Service.
func (s *appointmentSvc) FetchAllAppointmentByUserID(userId string) ([]models.Appointment, error) {
	return s.repo.FetchAllAppointmentByUserID(userId)
}
