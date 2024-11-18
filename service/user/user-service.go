package UserService

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	// Load .env file
	err := godotenv.Load("./config/.env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	// Get salt round from environment variable
	saltRound := os.Getenv("saltRound")
	if saltRound == "" {
		return nil, fmt.Errorf("saltRound not set in .env file")
	}

	// Convert salt round to integer
	s, err := strconv.Atoi(saltRound)
	if err != nil {
		return nil, fmt.Errorf("invalid saltRound value: %v", err)
	}

	// Hash the password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), s)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %v", err)
	}

	// Return hashed password as a byte slice
	return bytes, nil
}

func ComparePassword() (bool, error) {
	return true, nil
}
