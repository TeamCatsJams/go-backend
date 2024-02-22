package models

import "github.com/google/uuid"

type Appointment struct {
	ID            uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt     int        `json:"created_at"`
	UpdatedAt     int        `json:"updated_at"`
	AppointmentAt int        `json:"appointment_at"`
	MeetLink      string     `json:"meet_link"`
	UserID        *uuid.UUID `json:"user_id"`
	DoctorID      *uuid.UUID `json:"doctor_id"`
}
