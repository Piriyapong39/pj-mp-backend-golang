package UserService

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"

	"os"
	"time"

	"github.com/azujito/golang-api/service/model"
)

func GenToken(userData model.UserData) (string, error) {
	// import env
	if err := godotenv.Load("./config/.env"); err != nil {
		return "", err
	}
	JWT_SECRET_KEY := os.Getenv("JWT_SECRET_KEY")

	claims := jwt.MapClaims{
		"id":        userData.ID,
		"email":     userData.Email,
		"firstName": userData.FirstName,
		"lastName":  userData.LastName,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Use the actual JWT_SECRET_KEY instead of hardcoded "secret"
	t, err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}
	return t, nil
}
