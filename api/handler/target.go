package handler

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ww-api-gateway/api/presenter"
	"ww-api-gateway/pkg/entities"
	"ww-api-gateway/pkg/target"
)

func TargetGet(svc target.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		t, err := svc.Get(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.TargetErrorResponse(err))
		}
		return c.JSON(presenter.TargetSuccessResponse(t))
	}
}

func TargetAdd(svc target.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		t := new(entities.Target)
		if err := c.BodyParser(t); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.TargetErrorResponse(err))
		}
		t, err := svc.Create(t)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.TargetErrorResponse(err))
		}
		return c.JSON(presenter.TargetSuccessResponse(t))
	}
}

func TargetDelete(svc target.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := svc.Delete(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.TargetErrorResponse(err))
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}

func TargetUpdate(svc target.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		t := new(entities.Target)
		if err := c.BodyParser(t); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.TargetErrorResponse(err))
		}
		var err error
		t.ID, err = primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.TargetErrorResponse(err))
		}
		t, err = svc.Update(t)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.TargetErrorResponse(err))
		}
		return c.JSON(presenter.TargetSuccessResponse(t))
	}
}

func TargetsGetAll(svc target.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		t, err := svc.GetAll()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(t)
	}
}
