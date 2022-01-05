package auth

import (
	"meta-go/router/view"
	"meta-go/service/authservice"
	"time"

	"github.com/gofiber/fiber/v2"
)

type handler struct{}

var authHandler handler

func (handler) getRegister() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("auth/register", view.Resp(c))
	}
}

func (handler) postRegister() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := &RegisterRequest{}
		if err := view.ValidateRequestBody(c, req); err != nil {
			return c.Render("auth/register", view.RespWithWarn(c, err.Error()))
		}
		if err := authservice.Register(req.Username, req.Password); err != nil {
			return c.Render("auth/register", view.RespWithWarn(c, err.Error()))
		}
		return c.Redirect("/star/index")
	}
}

func (handler) getLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("auth/login", view.Resp(c))
	}
}

func (handler) postLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := &LoginRequest{}
		if err := view.ValidateRequestBody(c, req); err != nil {
			return c.Render("auth/login", view.RespWithWarn(c, err.Error()))
		}
		user, err := authservice.Login(req.Username, req.Password)
		if err != nil {
			return c.Render("auth/login", view.RespWithWarn(c, err.Error()))
		}
		// 设置Cookie
		c.Cookie(&fiber.Cookie{
			Name:     "login_info",
			Value:    user.Username,
			Expires:  time.Now().Add(1 * time.Hour),
			SameSite: "Lax",
		})

		return c.Redirect("/star/index")
	}
}

func (handler) getLogout() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "login_info",
			Expires:  time.Now().Add(-1 * time.Hour),
			SameSite: "Lax",
		})
		return c.Redirect("login")
	}
}
