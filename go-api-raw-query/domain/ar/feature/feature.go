package feature

import (
	"context"
	"fico_ar/config"
	"fico_ar/domain/ar/model"
	"fico_ar/domain/ar/repository"
	"fico_ar/domain/shared/response"
)

type ArFeature interface {
	GetAllData(ctx context.Context, payload *model.ARFilterList) (resp response.Data, err error)
	GetOneDataSales(ctx context.Context, salesID int64) (resp model.ARSales, err error)
	GetAllCompanyCode(ctx context.Context) (resp []model.ARSales, err error)
	CreateData(ctx context.Context, request model.ARRequest) (arID int64, err error)
	DeleteData(ctx context.Context, arID int64) (err error)
	UpdateData(ctx context.Context, request model.ARUpdatePayload, arID int64) (resp bool, err error)
	UpdateDataStatus(ctx context.Context, status int64, arID int64) (err error)
	GetOneData(ctx context.Context, arID int64) (resp model.ARResponse, err error)
}

type arFeature struct {
	config       config.EnvironmentConfig
	arRepository repository.ArRepository
}

func NewArFeature(config config.EnvironmentConfig, arRepo repository.ArRepository) ArFeature {
	return &arFeature{
		config:       config,
		arRepository: arRepo,
	}
}
