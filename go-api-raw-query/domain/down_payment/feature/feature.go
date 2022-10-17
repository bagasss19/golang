package feature

import (
	"context"
	"fico_ar/config"
	"fico_ar/domain/down_payment/model"
	"fico_ar/domain/down_payment/repository"
	"fico_ar/domain/shared/response"
)

type DPFeature interface {
	GetAllData(ctx context.Context, payload *model.GetDPListPayload) (resp response.Data, err error)
	GetOneData(ctx context.Context, giroID int64) (resp model.DownPayment, err error)
	CreateData(ctx context.Context, request model.DownPaymentRequest) (giroID int64, err error)
	DeleteData(ctx context.Context, arID int64) (err error)
	UpdateData(ctx context.Context, request model.DPUpdatePayload, giroID int64) (resp bool, err error)
}

type dpFeature struct {
	config       config.EnvironmentConfig
	dpRepository repository.DPRepository
}

func NewDPFeature(config config.EnvironmentConfig, dpRepo repository.DPRepository) DPFeature {
	return &dpFeature{
		config:       config,
		dpRepository: dpRepo,
	}
}
