package models

import "github.com/google/uuid"

type User struct {
	ID            uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt     int            `json:"created_at"`
	UpdatedAt     int            `json:"updated_at"`
	FullName      string         `json:"full_name"`
	Email         string         `json:"email" gorm:"unique"`
	Phone         string         `json:"phone"`
	Avatar        string         `json:"avatar"`
	Appointments  []Appointment  `json:"appointments" gorm:"foreignKey:UserID;default:null"`
	Prescriptions []Prescription `json:"prescriptions" gorm:"foreignKey:UserID;default:null"`
}
