package schemas

import "github.com/google/uuid"

type AppointmentData struct {
	AppointmentAt int       `json:"appointment_at"`
	DoctorID      uuid.UUID `json:"doctor_id"`
}
