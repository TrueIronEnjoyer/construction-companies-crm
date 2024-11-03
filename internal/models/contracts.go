package models

import "time"

type Contracts struct {
	Id            *int           `json:"id,omitempty"`
	EmployeeId    *int           `json:"employee_id,omitempty"`
	Date          *time.Time     `json:"date,omitempty"`
	Cost          *float32       `json:"cost,omitempty"`
	Status        *string        `json:"status,omitempty"`
	Description   *string        `json:"description,omitempty"`
	ContractBuyer *ContractBuyer `json:"contract_buyer,omitempty"`
}

type ContractBuyer struct {
	BuyerId        *int     `json:"buyer_id,omitempty"`
	OwnershipShare *float32 `json:"ownership_share,omitempty"`
}
