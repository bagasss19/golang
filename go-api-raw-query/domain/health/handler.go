package health

import (
	"fico_ar/domain/health/feature"
	"fico_ar/domain/shared/context"
	"fico_ar/domain/shared/response"

	"github.com/gofiber/fiber/v2"
)

type HealthHandler interface {
	ServiceHealth(c *fiber.Ctx) error
	Ping(c *fiber.Ctx) error
}

type healthHandler struct {
	healthFeature feature.HealthFeature
}

func NewHealthHandler(healthFeature feature.HealthFeature) HealthHandler {
	return &healthHandler{
		healthFeature: healthFeature,
	}
}

func (hh healthHandler) ServiceHealth(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	resp, err := hh.healthFeature.HealthCheck(ctx)
	if err != nil {
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOK(c, "", resp)
}

func (hh healthHandler) Ping(c *fiber.Ctx) error {
	return c.JSON("pong!")
}
