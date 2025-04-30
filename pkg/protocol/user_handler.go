package protocol

import (
	"backend/pkg/identity/user"
	"backend/pkg/infra/api/response"
	"backend/pkg/infra/api/routing"
	"strconv"

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
	var cmd user.CreateUserCommand

	err := c.BodyParser(&cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err)
	}

	err = cmd.Validate()
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err)
	}

	err = s.Dependencies.UserSvc.Create(c.Context(), &cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, err)
	}

	return response.SuccessMessage(c, "Created successfully")
}

func (s *Server) searchUser(c *fiber.Ctx) error {
	var (
		query       user.SearchUserQuery
		queryValues = c.Queries()
	)

	if len(queryValues) > 0 {
		err := c.QueryParser(&query)
		if err != nil {
			return response.SendError(c, fiber.StatusBadRequest, err)
		}
	}

	result, err := s.Dependencies.UserSvc.Search(c.Context(), &query)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, err)
	}

	return response.Result(c, result)
}

func (s *Server) getUserDetail(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is not valid")
	}

	result, err := s.Dependencies.UserSvc.GetByID(c.Context(), id)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, err)
	}

	return response.Result(c, result)
}

func (s *Server) updateUser(c *fiber.Ctx) error {
	var cmd user.UpdateUserCommand

	idStr := c.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is not valid")
	}

	err = c.BodyParser(&cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err)
	}

	cmd.ID = id
	err = cmd.Validate()
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err)
	}

	err = s.Dependencies.UserSvc.Update(c.Context(), &cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, err)
	}

	return response.SuccessMessage(c, "Updated successfully")
}

func (r *Server) updateStatus(c *fiber.Ctx) error {
	var cmd user.UpdateStatusCommand

	idStr := c.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is not valid")
	}

	err = c.BodyParser(&cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err)
	}

	cmd.ID = id
	err = cmd.Status.Validate()
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err)
	}

	err = r.Dependencies.UserSvc.UpdateStatus(c.Context(), &cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, err)
	}

	return response.SuccessMessage(c, "Updated successfully")
}

func (r *Server) updatePassword(c *fiber.Ctx) error {
	var cmd user.UpdatePasswordCommand

	idStr := c.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is not valid")
	}

	err = c.BodyParser(&cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err)
	}

	cmd.ID = id
	err = cmd.Validate()
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err)
	}

	err = r.Dependencies.UserSvc.UpdatePassword(c.Context(), &cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, err)
	}

	return response.SuccessMessage(c, "Updated password successfully")
}
