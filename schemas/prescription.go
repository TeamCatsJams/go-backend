package schemas

import "github.com/google/uuid"

type PrescriptionData struct {
	Title    string     `json:"title"`
	Document string     `json:"document"`
	UserID   *uuid.UUID `json:"user_id"`
}
