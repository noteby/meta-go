package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// 网关中间件
func Gate() fiber.Handler {
	return func(c *fiber.Ctx) error {

		err := c.Next()

		// 404
		if c.Response().StatusCode() == http.StatusNotFound {
			return c.Status(http.StatusNotFound).Render("404", fiber.Map{})
		}
		return err
	}
}
