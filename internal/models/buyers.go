package models

type Buyers struct {
	Id          int    `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	MiddleName  string `json:"middle_name,omitempty"`
	Email       string `json:"email,omitempty"`
	Phone       string `json:"phone,omitempty"`
	ContactInfo string `json:"contact_info,omitempty"`
	Description string `json:"description,omitempty"`
}
