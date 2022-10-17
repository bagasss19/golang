package repository

import (
	"context"
	"fico_ar/domain/ar/model"
	"fico_ar/infrastructure/database"
)

type ArRepository interface {
	GetAllData(ctx context.Context, payload *model.ARFilterList) (resp model.GetAllARResponse, err error)
	GetAllCompanyCode(ctx context.Context) (data []model.ARSales, err error)
	GetOneDataSales(ctx context.Context, salesID int64) (data model.ARSales, err error)
	CreateData(ctx context.Context, request model.AR) (arID int64, err error)
	CreateDataDetail(ctx context.Context, request model.ARDetail) (status bool, err error)
	DeleteData(ctx context.Context, arID int64) (err error)
	GetOneData(ctx context.Context, arID int64) (data model.ARResponse, err error)
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
