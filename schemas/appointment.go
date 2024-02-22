package schemas

import "github.com/google/uuid"

type AppointmentData struct {
	AppointmentAt int       `json:"appointment_at"`
	MeetLink      string    `json:"meet_link"`
	DoctorID      uuid.UUID `json:"doctor_id"`
}
