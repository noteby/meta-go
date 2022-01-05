package view

import (
	"meta-go/model"

	"github.com/gofiber/fiber/v2"
)

type alert struct {
	warn string
	info string
}

type user struct {
	id       uint
	username string
	isLogin  bool
}

func response(c *fiber.Ctx, warn string, info string, results ...map[string]interface{}) fiber.Map {
	var result map[string]interface{}
	if len(results) > 0 {
		result = results[0]
	}

	u := c.UserContext().Value("user").(model.User)
	var isLogin bool
	if u.ID != 0 {
		isLogin = true
	}
	return fiber.Map{
		"alert": alert{
			warn: warn,
			info: info,
		},
		"user": user{
			id:       u.ID,
			username: u.Username,
			isLogin:  isLogin,
		},
		"result": result,
	}
}

func RespWithWarn(c *fiber.Ctx, str string, results ...map[string]interface{}) fiber.Map {
	return response(c, str, "", results...)
}

func RespWithInfo(c *fiber.Ctx, str string, results ...map[string]interface{}) fiber.Map {
	return response(c, "", str, results...)
}

func Resp(c *fiber.Ctx, results ...map[string]interface{}) fiber.Map {
	return response(c, "", "", results...)
}
