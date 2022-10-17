package repository

import (
	"context"
	"fico_ar/domain/giro/model"
	"fico_ar/infrastructure/database"
)

type GiroRepository interface {
	GetAllData(ctx context.Context, payload *model.GetGiroListPayload) (resp model.GetAllGiroResponse, err error)
	CreateData(ctx context.Context, request model.Giro) (err error)
	GetOneData(ctx context.Context, giroID int64) (data model.Giro, err error)
	DeleteData(ctx context.Context, giroID int64) (err error)
	UpdateData(ctx context.Context, request model.GiroRequest, columns map[string]interface{}, giroID int64) (resp bool, err error)
}

type giroRepository struct {
	database *database.Database
}

func NewGiroFeature(db *database.Database) GiroRepository {
	return &giroRepository{
		database: db,
	}
}
