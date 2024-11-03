package services

import (
	"construction-companies-crm/internal/config"
	"construction-companies-crm/internal/models"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthStorage struct {
	EmployeesId map[string]int
}

var authStorage AuthStorage

func GetEmployeeId(jwt string) int {
	return authStorage.EmployeesId[jwt]
}

func GenerateJWR(e *models.EmployeeCredentials, id int) (string, error) {
	const op = "GenerateJWR"

	conf := config.GetConfig()

	payload := jwt.MapClaims{
		"sub": e.UserName,
		"exp": time.Now().Add(time.Hour * time.Duration(conf.TokenTimeoutHoursCount)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(conf.JwtSecret)
	if !errors.Is(err, nil) {
		return "", fmt.Errorf("%s: %v", op, err)
	}

	authStorage.EmployeesId[tokenStr] = id

	return tokenStr, nil
}
