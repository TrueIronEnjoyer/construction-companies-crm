package services

import (
	"construction-companies-crm/internal/database/postgresql"
	"construction-companies-crm/internal/models"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CreateEmployee(e *models.EmployeeRegister) (string, error) {
	const op = "RegisterEmployee"

	password_hash, err := bcrypt.GenerateFromPassword([]byte(e.EmployeeCredentials.Password), bcrypt.DefaultCost)
	if !errors.Is(err, nil) {
		return "", fmt.Errorf("%s: %v", op, err)
	}

	e.EmployeeCredentials.Password = string(password_hash)

	postgresql.CreateEmployee(e)

	return "", nil
}
