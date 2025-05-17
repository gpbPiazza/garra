package response

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorInfo struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

type Response struct {
	Data   any         `json:"data,omitempty"`
	Errors []ErrorInfo `json:"errors,omitempty"`
}

// OK will return http StatusOK and API Response as body with Data.
// Content type JSON
func OK(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(Response{Data: data})
}

func BadRequest(c *fiber.Ctx, errInfos ...ErrorInfo) error {
	return c.Status(fiber.StatusBadRequest).JSON(Response{Errors: errInfos})
}

func InternalError(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusInternalServerError)
}
