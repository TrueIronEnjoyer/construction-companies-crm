package models

type Complex struct {
	Id          *int    `json:"id,omitempty"`
	DeveloperId *int    `json:"developer_id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Location    *string `json:"location,omitempty"`
	Description *string `json:"description,omitempty"`
}
