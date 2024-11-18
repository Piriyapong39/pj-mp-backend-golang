package main

import (
	"fmt"
	"os"

	// import database
	database "github.com/azujito/golang-api/config"

	// import modules
	"github.com/azujito/golang-api/modules/user"

	// import package from dev.go
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// connect database
	db, err := database.Connection()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// import env
	if err = godotenv.Load("./config/.env"); err != nil {
		fmt.Println(err)
	}
	PORT := os.Getenv("PORT")

	// Start instance
	app := fiber.New()

	// import routes
	user.UserRoute(app)

	if err := app.Listen(":" + PORT); err != nil {
		fmt.Println(err)
	}

}
