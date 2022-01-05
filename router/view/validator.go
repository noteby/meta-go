package view

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateRequestBody(c *fiber.Ctx, requestStruct interface{}) error {
	err := c.BodyParser(requestStruct)
	if err != nil {
		return err
	}
	err = validate.Struct(requestStruct)
	return err

}

func ValidateRequestQuery(c *fiber.Ctx, requestStruct interface{}) error {
	err := c.QueryParser(requestStruct)
	if err != nil {
		return err
	}
	err = validate.Struct(requestStruct)
	return err
}
