package user

import (
	"fmt"

	// import database
	database "github.com/azujito/golang-api/config"

	// import service
	UserService "github.com/azujito/golang-api/service/user"

	// import model
	"github.com/azujito/golang-api/service/model"
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
	db, err := database.Connection()
	if err != nil {
		return "", fmt.Errorf("database connection error")
	}
	defer db.Close()

	userData := new(model.UserData)
	err = db.QueryRow(`
        SELECT id, email, password, first_name, last_name  
        FROM tb_users 
        WHERE email = $1
    `, userRequest.Email).Scan(&userData.ID, &userData.Email, &userData.Password, &userData.FirstName, &userData.LastName)

	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	// Verify password first
	if !UserService.CheckPasswordHash(userRequest.Password, userData.Password) {
		return "", fmt.Errorf("invalid credentials")
	}

	// Generate token
	token, err := UserService.GenToken(*userData)
	if err != nil {
		return "", fmt.Errorf("token generation error")
	}

	return token, nil
}
