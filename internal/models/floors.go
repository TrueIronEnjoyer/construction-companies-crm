package models

type Floor struct {
	Id          *int    `json:"id,omitempty"`
	HouseId     *int    `json:"house_id,omitempty"`
	Number      *int    `json:"number,omitempty"`
	Description *string `json:"description,omitempty"`
}
