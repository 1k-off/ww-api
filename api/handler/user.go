package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ww-api/api/presenter"
	"ww-api/pkg/entities"
	"ww-api/pkg/user"
)

func GetUser(svc user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		u, err := svc.Get(id)
		if err != nil {
			log.Err(err).Msgf("failed to get user with id %s", id)
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(u))
	}
}

func CreateUser(svc user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		u := new(entities.User)
		if err := c.BodyParser(u); err != nil {
			log.Debug().Err(err).Msgf("failed to parse user %v", c.Body())
			log.Err(err).Msg("failed to parse user")
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrorResponse(err))
		}
		u, err := svc.Create(u)
		if err != nil {
			log.Err(err).Msgf("failed to create user %s", u.Login)
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.UserErrorResponse(err))
		}
		log.Info().Msgf("user %s created", u.Login)
		return c.JSON(presenter.UserSuccessResponse(u))
	}
}

func UpdateUser(svc user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		u := new(entities.User)
		if err := c.BodyParser(u); err != nil {
			log.Err(err).Msg("failed to parse user")
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrorResponse(err))
		}
		var err error
		u.ID, err = primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			log.Err(err).Msgf("failed to parse user id %s", c.Params("id"))
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrorResponse(err))
		}
		u, err = svc.Update(u)
		if err != nil {
			log.Err(err).Msgf("failed to update user %s", u.Login)
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.UserErrorResponse(err))
		}
		log.Info().Msgf("user %s updated", u.Login)
		return c.JSON(presenter.UserSuccessResponse(u))
	}
}

func DeleteUser(svc user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := svc.Delete(id)
		if err != nil {
			log.Err(err).Msgf("failed to delete user with id %s", id)
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.UserErrorResponse(err))
		}
		log.Info().Msgf("user with id %s deleted", id)
		return c.SendStatus(fiber.StatusNoContent)
	}
}
