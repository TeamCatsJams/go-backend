package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Doctor struct {
	ID              uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt       int            `json:"created_at"`
	UpdatedAt       int            `json:"updated_at"`
	FullName        string         `json:"full_name"`
	Language        string         `json:"language"`
	Email           string         `json:"email" gorm:"unique"`
	Phone           string         `json:"phone"`
	MeetingDuration string         `json:"meeting_duration"`
	Description     string         `json:"description"`
	Avatar          string         `json:"avatar"`
	Country         string         `json:"country"`
	Practice        pq.StringArray `gorm:"type:text[]" json:"practice"`
	Location        string         `json:"location"`
	Designation     string         `json:"designation"`
	Appointments    []Appointment  `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Prescriptions   []Prescription `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
