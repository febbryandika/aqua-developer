package main

import (
	"user-management/config"
	"user-management/handler"
	"user-management/repository"
	"user-management/routes"
	"user-management/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Database()
	config.AutoMigrate()

	app := fiber.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	routes.Routes(app, userHandler)

	app.Listen(":3000")
}
