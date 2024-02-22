package db

import (
	"telmed_backend/pkgs/appointment"
	"telmed_backend/pkgs/doctor"
	"telmed_backend/pkgs/prescription"
	"telmed_backend/pkgs/user"
)

var (
	UserSvc         user.Service         = nil
	DoctorSvc       doctor.Service       = nil
	AppointmentSvc  appointment.Service  = nil
	PrescriptionSvc prescription.Service = nil
	// services
)

func InitServices() {
	db := GetDB()
	usersRepo := user.NewPostgresRepo(db)
	UserSvc = user.NewService(usersRepo)

	doctorRepo := doctor.NewPostgresRepo(db)
	DoctorSvc = doctor.NewService(doctorRepo)

	appointmentRepo := appointment.NewPostgresRepo(db)
	AppointmentSvc = appointment.NewService(appointmentRepo)
}
