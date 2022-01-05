package router

import (
	"meta-go/router/view/auth"
	"meta-go/router/view/star"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	auth.NewRouter(app.Group("auth/"))
	star.NewRouter(app.Group("star/"))
}
