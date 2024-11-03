package models

type House struct {
	Id          int    `json:"id,omitempty"`
	BlockId     int    `json:"block_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Address     string `json:"address,omitempty"`
	YearBuilt   int    `json:"year_built,omitempty"`
	Description string `json:"description,omitempty"`
}
