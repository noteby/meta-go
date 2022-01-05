package auth

import (
	"github.com/gofiber/fiber/v2"
)

func NewRouter(router fiber.Router) {
	router.Get("register", authHandler.getRegister())
	router.Post("register", authHandler.postRegister())
	router.Get("login", authHandler.getLogin())
	router.Post("login", authHandler.postLogin())
	router.Get("logout", authHandler.getLogout())
}
