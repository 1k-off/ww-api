package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ww-api/api/presenter"
	"ww-api/pkg/entities"
	"ww-api/pkg/target"
)

func TargetGet(svc target.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		t, err := svc.Get(id)
		if err != nil {
			log.Err(err).Msgf("failed to get target with id %s", id)
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.TargetErrorResponse(err))
		}
		return c.JSON(presenter.TargetSuccessResponse(t))
	}
}

func TargetAdd(svc target.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		t := new(entities.Target)
		if err := c.BodyParser(t); err != nil {
			log.Debug().Err(err).Msgf("failed to parse target %v", c.Body())
			log.Err(err).Msg("failed to parse target")
			return c.Status(fiber.StatusBadRequest).JSON(presenter.TargetErrorResponse(err))
		}
		t, err := svc.Create(t)
		if err != nil {
			log.Debug().Err(err).Msgf("failed to create target %v", t)
			log.Err(err).Msgf("failed to create target %s", t.URL)
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.TargetErrorResponse(err))
		}
		log.Info().Msgf("target %s created", t.URL)
		return c.JSON(presenter.TargetSuccessResponse(t))
	}
}

func TargetDelete(svc target.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := svc.Delete(id)
		if err != nil {
			log.Err(err).Msgf("failed to delete target with id %s", id)
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.TargetErrorResponse(err))
		}
		log.Info().Msgf("target with id %s deleted", id)
		return c.SendStatus(fiber.StatusNoContent)
	}
}

func TargetUpdate(svc target.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		t := new(entities.Target)
		if err := c.BodyParser(t); err != nil {
			log.Debug().Err(err).Msgf("failed to parse target %v", c.Body())
			log.Err(err).Msg("failed to parse target")
			return c.Status(fiber.StatusBadRequest).JSON(presenter.TargetErrorResponse(err))
		}
		var err error
		t.ID, err = primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			log.Err(err).Msgf("failed to parse target id %s", c.Params("id"))
			return c.Status(fiber.StatusBadRequest).JSON(presenter.TargetErrorResponse(err))
		}
		t, err = svc.Update(t)
		if err != nil {
			log.Debug().Err(err).Msgf("failed to update target %v", t)
			log.Err(err).Msgf("failed to update target %s", t.URL)
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.TargetErrorResponse(err))
		}
		log.Info().Msgf("target %s updated", t.URL)
		return c.JSON(presenter.TargetSuccessResponse(t))
	}
}

func TargetsGetAll(svc target.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		t, err := svc.GetAll()
		if err != nil {
			log.Err(err).Msg("failed to get all targets")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(t)
	}
}
