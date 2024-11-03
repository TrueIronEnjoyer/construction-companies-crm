package postgresql

import (
	"construction-companies-crm/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

func CreateEmployee(e *models.EmployeeRegister) (int, error) {
	const op = "CreateEmployee"

	db := GetDB()

	id := 0

	if err := db.QueryRow(`SELECT id FROM employee_credentials WHERE username = $1;`,
		e.EmployeeCredentials.UserName).Scan(&id); errors.Is(err, nil) {
		return 0, fmt.Errorf("username \"%s\" is already taken", e.EmployeeCredentials.UserName)
	} else if !errors.Is(err, sql.ErrNoRows) {
		return 0, fmt.Errorf("%s: %v", op, err)
	}

	_, err := db.Exec(`INSERT INTO employee_credentials (username, password_hash)
		VALUES ($1, $2)`, e.EmployeeCredentials.UserName, e.EmployeeCredentials.Password)
	if !errors.Is(err, nil) {
		return 0, fmt.Errorf("%s: %v", op, err)
	}

	if err := db.QueryRow(`SELECT id FROM employee_credentials WHERE username = $1;`,
		e.EmployeeCredentials.UserName).Scan(&id); !errors.Is(err, nil) {
		return 0, fmt.Errorf("%s: %v", op, err)
	}

	return id, nil
}
