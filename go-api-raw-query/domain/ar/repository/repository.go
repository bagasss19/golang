package repository

import (
	"context"
	"fico_ar/domain/ar/model"
	"fico_ar/infrastructure/database"
)

type ArRepository interface {
	GetAllData(ctx context.Context, payload *model.ARFilterList) (resp model.GetAllARResponse, err error)
	GetAllCompanyCode(ctx context.Context) (data []model.ARSales, err error)
	DeleteData(ctx context.Context, arID int64) (err error)
	GetOneData(ctx context.Context, arID int64) (data model.AR, err error)
	UpdateData(ctx context.Context, request model.ARUpdate, columns map[string]interface{}, arID int64) (resp bool, err error)
	UpdateStatusData(ctx context.Context, status int64, arID int64) (err error)
}

type arRepository struct {
	database *database.Database
}

func NewArFeature(db *database.Database) ArRepository {
	return &arRepository{
		database: db,
	}
}
