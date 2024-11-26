package user

import (
	"fmt"

	// import database
	database "github.com/azujito/golang-api/config"

	// import service
	UserService "github.com/azujito/golang-api/service/user"
)

func _userRegister(user User) (string, error) {
	// connect database
	db, err := database.Connection()
	if err != nil {
		return "", err
	}
	defer db.Close()

	var exists bool
	if err = db.QueryRow(`
        SELECT EXISTS(SELECT 1 FROM tb_users WHERE email = $1)
    `, user.Email).Scan(&exists); err != nil {
		return "", err
	}

	if exists {
		return "", fmt.Errorf("email already exists")
	}

	hashPassword, err := UserService.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	if _, err = db.Exec(`
        INSERT INTO tb_users (email, password, first_name, last_name)
        VALUES ($1, $2, $3, $4)
    `, user.Email, hashPassword, user.FirstName, user.LastName); err != nil {
		return "", err
	}

	return "registered successfully", nil
}

func _userLogin(userRequest User) (string, error) {
	return "1234", nil
}
