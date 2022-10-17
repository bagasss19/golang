package feature

import (
	"context"
	"database/sql"
	"errors"
	"fico_ar/domain/ar/model"
	"fico_ar/domain/shared/constant"
	"fico_ar/domain/shared/response"
	"math"
	"time"

	"github.com/spf13/cast"
)

func (ar arFeature) GetAllData(ctx context.Context, payload *model.ARFilterList) (resp response.Data, err error) {
	if payload.Limit == 0 {
		payload.Limit = 5
	}

	if payload.Page == 0 || payload.Page == 1 {
		payload.Offset = 0
	} else {
		payload.Offset = (payload.Page - 1) * payload.Limit
	}

	data, err := ar.arRepository.GetAllData(ctx, payload)
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

func (ar arFeature) GetOneData(ctx context.Context, arID int64) (resp model.ARResponse, err error) {
	resp, err = ar.arRepository.GetOneData(ctx, arID)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (ar arFeature) GetOneDataSales(ctx context.Context, salesID int64) (resp model.ARSales, err error) {
	resp, err = ar.arRepository.GetOneDataSales(ctx, salesID)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (ar arFeature) GetAllCompanyCode(ctx context.Context) (resp []model.ARSales, err error) {
	data, err := ar.arRepository.GetAllCompanyCode(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (ar arFeature) CreateData(ctx context.Context, request model.ARRequest) (arID int64, err error) {
	request.Status = 0
	arPayload, err := model.NewAR(request)
	if err != nil {
		return 0, err
	}

	invoice, err := time.Parse(constant.TimeFormat, request.Invoice)
	if err != nil {
		return 0, err
	}

	arID, err = ar.arRepository.CreateData(ctx, arPayload)
	if err != nil {
		return 0, err
	}

	totalPayment := request.DiscPayment + request.CashPayment + request.GiroAmount + request.CNDNAmount + request.ReturnAmount
	arDetailPayload := model.ARDetail{
		ARID:          arID,
		TransactionID: request.TransactionID,
		Invoice: sql.NullTime{
			Time:  invoice,
			Valid: true,
		},
		TotalPayment:   totalPayment,
		DiscPayment:    request.DiscPayment,
		CashPayment:    request.CashPayment,
		GiroNumber:     request.GiroNumber,
		GiroAmount:     request.GiroAmount,
		TransferNumber: request.TransferNumber,
		TransferAmount: request.TransferAmount,
		CNDNNumber:     request.CNDNNumber,
		CNDNAmount:     request.CNDNAmount,
		ReturnNumber:   request.ReturnNumber,
		ReturnAmount:   request.ReturnAmount,
		Status:         0,
	}
	_, err = ar.arRepository.CreateDataDetail(ctx, arDetailPayload)
	if err != nil {
		return 0, err
	}

	return arID, nil
}

func (ar arFeature) DeleteData(ctx context.Context, arID int64) (err error) {
	data, err := ar.arRepository.GetOneData(ctx, arID)
	if err != nil {
		return err
	}

	if data.Status != 0 {
		return errors.New("[DeleteData] you can only delete draft data")
	}
	err = ar.arRepository.DeleteData(ctx, arID)
	if err != nil {
		return err
	}

	return nil
}

func (ar arFeature) UpdateData(ctx context.Context, request model.ARUpdatePayload, arID int64) (resp bool, err error) {
	var (
		updatedAR model.ARUpdate
	)

	data, err := ar.arRepository.GetOneData(ctx, arID)
	if err != nil {
		return false, err
	}

	fields := make(map[string]interface{})

	for field, val := range request.Data {
		switch field {
		case "total_payment":
			newValue, err := cast.ToFloat64E(val)
			if err != nil {
				continue
			}
			updatedAR.TotalPayment = newValue
			fields["total_payment"] = "TotalPayment"

		case "disc_payment":
			newValue, err := cast.ToFloat64E(val)
			if err != nil {
				continue
			}
			updatedAR.DiscPayment = newValue
			fields["disc_payment"] = "DiscPayment"

		case "cash_payment":
			newValue, err := cast.ToFloat64E(val)
			if err != nil {
				continue
			}
			updatedAR.CashPayment = newValue
			fields["cash_payment"] = "CashPayment"

		case "giro_amount":
			newValue, err := cast.ToFloat64E(val)
			if err != nil {
				continue
			}
			updatedAR.GiroAmount = newValue
			fields["giro_amount"] = "GiroAmount"

		case "giro_num":
			newValue, err := cast.ToInt64E(val)
			if err != nil {
				continue
			}
			updatedAR.GiroNumber = newValue
			fields["giro_num"] = "GiroNumber"

		case "transfer_amount":
			newValue, err := cast.ToFloat64E(val)
			if err != nil {
				continue
			}
			updatedAR.TransferAmount = newValue
			fields["transfer_amount"] = "TransferAmount"

		case "transfer_num":
			newValue, err := cast.ToInt64E(val)
			if err != nil {
				continue
			}
			updatedAR.TransferNumber = newValue
			fields["transfer_num"] = "TransferNumber"

		case "cndn_amount":
			newValue, err := cast.ToFloat64E(val)
			if err != nil {
				continue
			}
			updatedAR.CNDNAmount = newValue
			fields["cndn_amount"] = "CNDNAmount"

		case "cndn_num":
			newValue, err := cast.ToInt64E(val)
			if err != nil {
				continue
			}
			updatedAR.CNDNNumber = newValue
			fields["cndn_num"] = "CNDNNumber"

		case "return_amount":
			newValue, err := cast.ToFloat64E(val)
			if err != nil {
				continue
			}
			updatedAR.ReturnAmount = newValue
			fields["return_amount"] = "ReturnAmount"

		case "return_num":
			newValue, err := cast.ToInt64E(val)
			if err != nil {
				continue
			}
			updatedAR.ReturnNumber = newValue
			fields["return_num"] = "ReturnNumber"

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
		_, err = ar.arRepository.UpdateData(ctx, updatedAR, fields, arID)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (ar arFeature) UpdateDataStatus(ctx context.Context, status int64, arID int64) (err error) {
	err = ar.arRepository.UpdateStatusData(ctx, status, arID)
	if err != nil {
		return err
	}

	return nil
}
