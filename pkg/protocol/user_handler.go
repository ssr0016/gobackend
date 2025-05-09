package protocol

import (
	"backend/pkg/identity/account"
	"backend/pkg/identity/user"
	"backend/pkg/infra/api/middleware"
	"backend/pkg/infra/api/response"
	"backend/pkg/infra/api/routing"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var (
	reqOnlyUser = middleware.ForUser(account.Admin)
	// reqReadUser   = middleware.ForPermission(accesscontrol.ActionReadRole)
	// reqCreateUser = middleware.ForPermission(accesscontrol.ActionCreateUser)
	// reqUpdateUser = middleware.ForPermission(accesscontrol.ActionUpdateUser)
)

func (s *Server) NewUserHandler(r *routing.Router) {
	admin := r.Group("/api/user")

	// 🚀 No authentication required for user creation
	admin.Handle(fiber.MethodPost, "/", s.createUser)

	// ✅ Authentication required for other operations
	admin.Handle(fiber.MethodGet, "/", s.searchUser, reqOnlyUser)
	admin.Handle(fiber.MethodGet, "/:id", s.getUserDetail, reqOnlyUser)
	admin.Handle(fiber.MethodPut, "/:id", s.updateUser, reqOnlyUser)
	admin.Handle(fiber.MethodPut, "/:id/status", s.updateStatus, reqOnlyUser)
	admin.Handle(fiber.MethodPut, "/:id/password", s.updatePassword, reqOnlyUser)
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

func (s *Server) updateStatus(c *fiber.Ctx) error {
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

	err = cmd.Validate()
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err)
	}

	err = s.Dependencies.UserSvc.UpdateStatus(c.Context(), &cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, err)
	}

	return response.SuccessMessage(c, "Update status successfully")
}

func (s *Server) updatePassword(c *fiber.Ctx) error {
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

	err = s.Dependencies.UserSvc.UpdatePassword(c.Context(), &cmd)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, err)
	}

	return response.SuccessMessage(c, "Update password successfully")
}
