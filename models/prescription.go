package models

import "github.com/google/uuid"

type Prescription struct {
	ID        uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt int        `json:"created_at"`
	UpdatedAt int        `json:"updated_at"`
	Title     string     `json:"title"`
	Document  string     `json:"document"`
	UserID    *uuid.UUID `json:"user_id"`
	DoctorID  *uuid.UUID `json:"doctor_id"`
}
