package routes

import (
	"user-management/handler"

	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router, userHandler *handler.UserHandler) {
	r := app.Group("api/v1")

	r.Get("/users", userHandler.GetList)
	r.Post("/user", userHandler.CreateUser)
	r.Put("/user/edit/:id", userHandler.UpdateUser)
	r.Delete("/user/delete/:id", userHandler.DeleteUser)
}
