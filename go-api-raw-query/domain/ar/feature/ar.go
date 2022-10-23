package feature

import (
	"context"
	"errors"
	"fico_ar/domain/ar/model"
	"fico_ar/domain/shared/response"
	"math"

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

func (ar arFeature) GetOneData(ctx context.Context, arID int64) (resp model.AR, err error) {
	resp, err = ar.arRepository.GetOneData(ctx, arID)
	if err != nil {
		return resp, err
	}

	return resp, nil
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
		case "invoice_type":
			newValue, err := cast.ToStringE(val)
			if err != nil {
				continue
			}
			updatedAR.InvoiceType = newValue
			fields["invoice_type"] = "InvoiceType"

		case "invoice_value":
			newValue, err := cast.ToFloat64E(val)
			if err != nil {
				continue
			}
			updatedAR.InvoiceValue = newValue
			fields["invoice_value"] = "InvoiceValue"

		case "total_payment":
			newValue, err := cast.ToFloat64E(val)
			if err != nil {
				continue
			}
			updatedAR.TotalPayment = newValue
			fields["total_payment"] = "TotalPayment"

		case "discount_payment":
			newValue, err := cast.ToFloat64E(val)
			if err != nil {
				continue
			}
			updatedAR.DiscountPayment = newValue
			fields["discount_payment"] = "DiscountPayment"

		case "cash_payment":
			newValue, err := cast.ToFloat64E(val)
			if err != nil {
				continue
			}
			updatedAR.CashPayment = newValue
			fields["cash_payment"] = "CashPayment"

		case "giro_number":
			newValue, err := cast.ToInt64E(val)
			if err != nil {
				continue
			}
			updatedAR.GiroNumber = newValue
			fields["giro_number"] = "GiroNumber"

		case "cndn_payment":
			newValue, err := cast.ToFloat64E(val)
			if err != nil {
				continue
			}
			updatedAR.CNDNPayment = newValue
			fields["cndn_payment"] = "CNDNPayment"

		case "cndn_number":
			newValue, err := cast.ToInt64E(val)
			if err != nil {
				continue
			}
			updatedAR.CNDNNumber = newValue
			fields["cndn_number"] = "CNDNNumber"

		case "return_payment":
			newValue, err := cast.ToFloat64E(val)
			if err != nil {
				continue
			}
			updatedAR.ReturnPayment = newValue
			fields["return_payment"] = "ReturnPayment"

		case "return_number":
			newValue, err := cast.ToInt64E(val)
			if err != nil {
				continue
			}
			updatedAR.ReturnNumber = newValue
			fields["return_number"] = "ReturnNumber"

		case "down_payment":
			newValue, err := cast.ToFloat64E(val)
			if err != nil {
				continue
			}
			updatedAR.DownPayment = newValue
			fields["down_payment"] = "DownPayment"

		case "down_payment_number":
			newValue, err := cast.ToInt64E(val)
			if err != nil {
				continue
			}
			updatedAR.DownPaymentNumber = newValue
			fields["down_payment_number"] = "DownPaymentNumber"

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
		if data.Status != 1 {
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

func (ar arFeature) GetAllCompanyCode(ctx context.Context) (resp []model.ARSales, err error) {
	data, err := ar.arRepository.GetAllCompanyCode(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (ar arFeature) DeleteData(ctx context.Context, arID int64) (err error) {
	data, err := ar.arRepository.GetOneData(ctx, arID)
	if err != nil {
		return err
	}

	if data.Status != 1 {
		return errors.New("[DeleteData] you can only delete draft data")
	}
	err = ar.arRepository.DeleteData(ctx, arID)
	if err != nil {
		return err
	}

	return nil
}
