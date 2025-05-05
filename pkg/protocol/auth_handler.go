package protocol

import (
	"backend/pkg/identity/auth"
	"backend/pkg/infra/api/response"
	"backend/pkg/infra/api/routing"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) NewAuthHandler(r *routing.Router) {
	admin := r.Group("/api/auth")

	admin.POST("/login", s.loginUser)
}

func (s *Server) loginUser(c *fiber.Ctx) error {
	var cmd auth.LoginCommand

	err := c.BodyParser(&cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err)
	}

	err = cmd.Validate()
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err)
	}

	result, err := s.Dependencies.AuthSvc.Login(c.Context(), &cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, err)
	}

	return response.Result(c, result)
}
