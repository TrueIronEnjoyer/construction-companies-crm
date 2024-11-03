package models

type Blocks struct {
	Id          int    `json:"id,omitempty"`
	ComplexId   int    `json:"complex_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Descriptiom string `json:"descriptiom,omitempty"`
}
