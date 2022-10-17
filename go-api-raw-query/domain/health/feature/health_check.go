package feature

import (
	"context"
	"fico_ar/domain/health/constant"
	"fico_ar/domain/health/model"
)

func (hf healthFeature) HealthCheck(ctx context.Context) (resp model.HealthCheck, err error) {

	var (
		status   = constant.HEALTHY
		dbstatus = constant.CONNECTED
	)

	db, err := hf.healthRepository.DatabaseHealth()
	if !db {
		status = constant.NOT_HEALTHY
		dbstatus = constant.DISCONECTED
	} else if err != nil {
		status = constant.NOT_HEALTHY
	}

	resp = model.HealthCheck{
		AppDetail: model.AppDetail{
			Name:    hf.config.App.Name,
			Version: hf.config.App.Version,
		},
		DatabaseDetail: model.DatabaseDetail{
			Dialect: hf.config.Database.Dialect,
			Status:  dbstatus,
		},
		Status: status,
	}

	return
}
