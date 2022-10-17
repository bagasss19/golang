package feature

import (
	"context"
	"fico_ar/config"
	"fico_ar/domain/health/model"
	"fico_ar/domain/health/repository"
)

type HealthFeature interface {
	HealthCheck(ctx context.Context) (resp model.HealthCheck, err error)
}

type healthFeature struct {
	config           config.EnvironmentConfig
	healthRepository repository.HealthRepository
}

func NewHealthFeature(config config.EnvironmentConfig, healthRepo repository.HealthRepository) HealthFeature {
	return &healthFeature{
		config:           config,
		healthRepository: healthRepo,
	}
}
