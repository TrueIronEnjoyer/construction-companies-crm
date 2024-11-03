package models

import "time"

type Developer struct {
	Id              *int       `json:"id,omitempty"`
	Name            *string    `json:"name,omitempty"`
	Address         *string    `json:"address,omitempty"`
	Phone           *string    `json:"phone,omitempty"`
	Email           *string    `json:"email,omitempty"`
	ContactInfo     *string    `json:"contact_info,omitempty"`
	EstablishedDate *time.Time `json:"established_date,omitempty"`
	Description     *string    `json:"description,omitempty"`
}
