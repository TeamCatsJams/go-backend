package schemas

import "github.com/google/uuid"

type PrescriptionData struct {
	Title    string     `json:"title"`
	Document string     `json:"document"`
	DoctorID *uuid.UUID `json:"doctor_id"`
}
