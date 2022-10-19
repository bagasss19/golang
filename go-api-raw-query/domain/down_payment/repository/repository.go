package repository

import (
	"context"
	"fico_ar/domain/down_payment/model"
	"fico_ar/infrastructure/database"
)

type DPRepository interface {
	GetAllData(ctx context.Context, payload *model.GetDPListPayload) (resp model.GetAllDPResponse, err error)
	CreateData(ctx context.Context, request model.DownPayment) (err error)
	CreateDataDetail(ctx context.Context, request model.DownPaymentDetail) (err error)
	GetOneData(ctx context.Context, dpID int64) (data model.DownPayment, err error)
	DeleteData(ctx context.Context, dpID int64) (err error)
	UpdateData(ctx context.Context, request model.DownPaymentRequest, columns map[string]interface{}, arID int64) (resp bool, err error)
}

type dpRepository struct {
	database *database.Database
}

func NewDPFeature(db *database.Database) DPRepository {
	return &dpRepository{
		database: db,
	}
}
