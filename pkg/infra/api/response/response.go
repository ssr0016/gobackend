package response

import (
	"backend/pkg/infra/api/errors"

	"github.com/gofiber/fiber/v2"
)

func SendError(c *fiber.Ctx, status int, err error) error {
	var result errors.ErrorStatus

	switch e := err.(type) {
	case errors.ErrorStatus:
		result = e
	default:
		result = errors.ErrorStatus{
			Code:    errors.GeneralErrorCode,
			Message: e.Error(),
		}
	}

	return c.Status(status).JSON(result)
}

func SuccessMessage(c *fiber.Ctx, message string) error {
	return c.JSON(map[string]string{"message": message})
}

func Result(c *fiber.Ctx, data interface{}) error {
	return c.JSON(data)
}
