package protocol

import (
	"backend/pkg/infra/api/routing"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) NewUserHandler(r *routing.Router) {
	admin := r.Group("/api/user")

	admin.POST("/", s.createUser)
	admin.GET("/", s.searchUser)
	admin.GET("/:id", s.getUserDetail)
	admin.PUT("/:id", s.updateUser)
	admin.PUT("/:id/status", s.updateStatus)
	admin.PUT("/:id/password", s.updatePassword)
}

func (s *Server) createUser(c *fiber.Ctx) error {
	return c.SendString("Create User")
}

func (s *Server) searchUser(c *fiber.Ctx) error {
	return c.SendString("Search Users")
}

func (s *Server) getUserDetail(c *fiber.Ctx) error {
	return c.SendString("Get User detail")
}

func (s *Server) updateUser(c *fiber.Ctx) error {
	return c.SendString("Update User")
}

func (r *Server) updateStatus(c *fiber.Ctx) error {
	return c.SendString("Update User Status")
}

func (r *Server) updatePassword(c *fiber.Ctx) error {
	return c.SendString("Update User Password")
}
