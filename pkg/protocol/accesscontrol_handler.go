package protocol

import (
	"backend/pkg/infra/api/response"
	"backend/pkg/infra/api/routing"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) NewAccessControlHandler(r *routing.Router) {
	admin := r.Group("/api/accesscontrol")

	admin.GET("/permission", s.getPermission)

}
func (s *Server) getPermission(c *fiber.Ctx) error {
	result := s.Dependencies.AccessControlSvc.GetPermissions(c.Context())

	return response.Result(c, result)
}
