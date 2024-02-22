package schemas

type DoctorData struct {
	FullName        string   `json:"full_name"`
	Language        string   `json:"language"`
	Email           string   `json:"email" gorm:"unique"`
	Phone           string   `json:"phone"`
	MeetingDuration string   `json:"meeting_duration"`
	Description     string   `json:"description"`
	Avatar          string   `json:"avatar"`
	Country         string   `json:"country"`
	Practice        []string `json:"practice"`
	Location        string   `json:"location"`
	Designation     string   `json:"designation"`
}
