package feature

import (
	"context"
	"errors"
	"fico_ar/domain/giro/model"
	"fico_ar/domain/shared/response"
	"math"

	"github.com/spf13/cast"
)

func (g giroFeature) GetAllData(ctx context.Context, payload *model.GetGiroListPayload) (resp response.Data, err error) {
	if payload.Limit == 0 {
		payload.Limit = 5
	}

	if payload.Page == 0 || payload.Page == 1 {
		payload.Offset = 0
	} else {
		payload.Offset = (payload.Page - 1) * payload.Limit
	}

	data, err := g.giroRepository.GetAllData(ctx, payload)
	if err != nil {
		return resp, err
	}

	pagination := response.Pagination{
		LimitPerPage: payload.Limit,
		CurrentPage:  payload.Page,
		TotalPage:    int64(math.Ceil(float64(data.TotalItem) / float64(payload.Limit))),
		TotalRows:    int64(len(data.Data)),
		TotalItems:   data.TotalItem,
	}

	if payload.Page == 0 {
		pagination.CurrentPage = 1
	}

	if pagination.CurrentPage == 1 {
		pagination.First = true
	} else {
		pagination.First = false
	}

	if pagination.CurrentPage == pagination.TotalPage {
		pagination.Last = true
	} else {
		pagination.Last = false
	}

	resp.Items = data.Data
	resp.Pagination = pagination

	return resp, nil
}

func (g giroFeature) GetOneData(ctx context.Context, giroID int64) (resp model.Giro, err error) {
	resp, err = g.giroRepository.GetOneData(ctx, giroID)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (g giroFeature) CreateData(ctx context.Context, request model.GiroRequest) (err error) {
	request.Status = 0
	giroPayload, err := model.NewGiro(request)
	if err != nil {
		return err
	}

	err = g.giroRepository.CreateData(ctx, giroPayload)
	if err != nil {
		return err
	}

	return nil
}

func (g giroFeature) DeleteData(ctx context.Context, giroID int64) (err error) {
	data, err := g.giroRepository.GetOneData(ctx, giroID)
	if err != nil {
		return err
	}

	if data.Status != 0 {
		return errors.New("[DeleteData] you can only delete draft data")
	}
	err = g.giroRepository.DeleteData(ctx, giroID)
	if err != nil {
		return err
	}

	return nil
}

func (g giroFeature) UpdateData(ctx context.Context, request model.GiroUpdatePayload, giroID int64) (resp bool, err error) {
	var (
		updatedAR model.GiroRequest
	)

	data, err := g.giroRepository.GetOneData(ctx, giroID)
	if err != nil {
		return false, err
	}

	fields := make(map[string]interface{})

	for field, val := range request.Data {
		switch field {
		case "company_id":
			newValue, err := cast.ToStringE(val)
			if err != nil {
				continue
			}
			updatedAR.CompanyID = newValue
			fields["company_id"] = "CompanyID"

		case "giro_date":
			newValue, err := cast.ToStringE(val)
			if err != nil {
				continue
			}
			updatedAR.GiroDate = newValue
			fields["giro_date"] = "GiroDate"

		case "giro_num":
			newValue, err := cast.ToInt64E(val)
			if err != nil {
				continue
			}
			updatedAR.GiroNum = newValue
			fields["giro_num"] = "GiroNum"

		case "account_id":
			newValue, err := cast.ToStringE(val)
			if err != nil {
				continue
			}
			updatedAR.AccountID = newValue
			fields["account_id"] = "AccountID"

		case "account_name":
			newValue, err := cast.ToStringE(val)
			if err != nil {
				continue
			}
			updatedAR.AccountName = newValue
			fields["account_name"] = "AccountName"

		case "giro_amount":
			newValue, err := cast.ToFloat64E(val)
			if err != nil {
				continue
			}
			updatedAR.GiroAmount = newValue
			fields["giro_amount"] = "GiroAmount"

		case "status":
			newValue, err := cast.ToInt64E(val)
			if err != nil {
				continue
			}
			updatedAR.Status = newValue
			fields["status"] = "Status"
		}
	}

	if len(fields) > 0 {
		if data.Status != 0 {
			return false, errors.New("[UpdateData] you can only edit draft data")
		}
		_, err = g.giroRepository.UpdateData(ctx, updatedAR, fields, giroID)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
