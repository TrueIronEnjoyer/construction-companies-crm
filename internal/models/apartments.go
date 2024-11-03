package models

type Apartment struct {
	Id          int     `json:"id,omitempty"`
	HouseID     int     `json:"house_id,omitempty"`
	FloorId     int     `json:"floor_id,omitempty"`
	Entrance    int     `json:"entrance,omitempty"`
	Number      int     `json:"number,omitempty"`
	Area        float32 `json:"area,omitempty"`
	Rooms       int     `json:"rooms,omitempty"`
	Status      string  `json:"status,omitempty"`
	Description string  `json:"description,omitempty"`
}
