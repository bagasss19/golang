package feature

import (
	"context"
	"fico_ar/config"
	"fico_ar/domain/giro/model"
	"fico_ar/domain/giro/repository"
	"fico_ar/domain/shared/response"
)

type GiroFeature interface {
	GetAllData(ctx context.Context, payload *model.GetGiroListPayload) (resp response.Data, err error)
	GetOneData(ctx context.Context, giroID int64) (resp model.Giro, err error)
	CreateData(ctx context.Context, request model.GiroRequest) (err error)
	DeleteData(ctx context.Context, giroID int64) (err error)
	UpdateData(ctx context.Context, request model.GiroUpdatePayload, giroID int64) (resp bool, err error)
}

type giroFeature struct {
	config         config.EnvironmentConfig
	giroRepository repository.GiroRepository
}

func NewGiroFeature(config config.EnvironmentConfig, giroRepo repository.GiroRepository) GiroFeature {
	return &giroFeature{
		config:         config,
		giroRepository: giroRepo,
	}
}
