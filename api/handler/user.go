package handler

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ww-api-gateway/api/presenter"
	"ww-api-gateway/pkg/entities"
	"ww-api-gateway/pkg/user"
)

func GetUser(svc user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		u, err := svc.Get(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(u))
	}
}

func CreateUser(svc user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		u := new(entities.User)
		if err := c.BodyParser(u); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrorResponse(err))
		}
		u, err := svc.Create(u)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(u))
	}
}

func UpdateUser(svc user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		u := new(entities.User)
		if err := c.BodyParser(u); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrorResponse(err))
		}
		var err error
		u.ID, err = primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrorResponse(err))
		}
		u, err = svc.Update(u)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(u))
	}
}

func DeleteUser(svc user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := svc.Delete(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.UserErrorResponse(err))
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}
