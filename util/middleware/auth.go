package middleware

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"meta-go/model"
	"meta-go/router/view"
	"meta-go/service/authservice"
)

// 用户鉴权
func UserAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user model.User
		login_info := c.Cookies("login_info")
		if login_info != "" {
			user = authservice.GetUserByUsername(login_info)
		}

		ctx := c.UserContext()
		c.SetUserContext(context.WithValue(ctx, "user", user))

		err := c.Next()

		return err
	}
}

// 需要登录才能访问
func LoginRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.UserContext().Value("user").(model.User)
		if user.ID == 0 {
			return c.Render("base", view.RespWithWarn(c, "没有权限访问"))
		}

		err := c.Next()
		return err
	}
}
