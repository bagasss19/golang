package feature

import (
	"context"
	"errors"
	"fico_ar/domain/down_payment/model"
	"fico_ar/domain/shared/response"
	"math"

	"github.com/spf13/cast"
)

func (d dpFeature) GetAllData(ctx context.Context, payload *model.GetDPListPayload) (resp response.Data, err error) {
	if payload.Limit == 0 {
		payload.Limit = 5
	}

	if payload.Page == 0 || payload.Page == 1 {
		payload.Offset = 0
	} else {
		payload.Offset = (payload.Page - 1) * payload.Limit
	}

	data, err := d.dpRepository.GetAllData(ctx, payload)
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

func (d dpFeature) GetAllDataDetail(ctx context.Context, payload *model.GetDPDetailListPayload) (resp response.Data, err error) {
	if payload.Limit == 0 {
		payload.Limit = 5
	}

	if payload.Page == 0 || payload.Page == 1 {
		payload.Offset = 0
	} else {
		payload.Offset = (payload.Page - 1) * payload.Limit
	}

	data, err := d.dpRepository.GetAllDataDetail(ctx, payload)
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

func (d dpFeature) GetOneData(ctx context.Context, dpID int64) (resp model.DownPayment, err error) {
	resp, err = d.dpRepository.GetOneData(ctx, dpID)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (d dpFeature) GetOneDataDetail(ctx context.Context, dpDetailID int64) (resp model.DownPaymentDetail, err error) {
	resp, err = d.dpRepository.GetOneDataDetail(ctx, dpDetailID)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (d dpFeature) CreateData(ctx context.Context, request model.DownPaymentRequest) (dpID int64, err error) {
	request.Status = 0
	dpPayload, err := model.NewDP(request)
	if err != nil {
		return 0, err
	}

	err = d.dpRepository.CreateData(ctx, dpPayload)
	if err != nil {
		return 0, err
	}

	return dpID, nil
}

func (d dpFeature) CreateDataDetail(ctx context.Context, request model.DownPaymentDetailRequest) (dpDetailID int64, err error) {
	request.Status = 0
	dpDetailPayload, err := model.NewDPDetail(request)
	if err != nil {
		return 0, err
	}

	err = d.dpRepository.CreateDataDetail(ctx, dpDetailPayload)
	if err != nil {
		return 0, err
	}

	return dpDetailID, nil
}

func (d dpFeature) DeleteData(ctx context.Context, dpID int64) (err error) {
	data, err := d.dpRepository.GetOneData(ctx, dpID)
	if err != nil {
		return err
	}

	if data.Status != 0 {
		return errors.New("[DeleteData] you can only delete draft data")
	}
	err = d.dpRepository.DeleteData(ctx, dpID)
	if err != nil {
		return err
	}

	return nil
}

func (d dpFeature) DeleteDataDetail(ctx context.Context, dpDetailID int64) (err error) {
	data, err := d.dpRepository.GetOneDataDetail(ctx, dpDetailID)
	if err != nil {
		return err
	}

	if data.Status != 0 {
		return errors.New("[DeleteDataDetail] you can only delete draft data")
	}
	err = d.dpRepository.DeleteDataDetail(ctx, dpDetailID)
	if err != nil {
		return err
	}
	
	return nil
}

func (d dpFeature) UpdateData(ctx context.Context, request model.DPUpdatePayload, dpID int64) (resp bool, err error) {
	var (
		updatedDP model.DownPaymentRequest
	)

	data, err := d.dpRepository.GetOneData(ctx, dpID)
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
			updatedDP.CompanyID = newValue
			fields["company_id"] = "CompanyID"

		case "outlet_id":
			newValue, err := cast.ToInt64E(val)
			if err != nil {
				continue
			}
			updatedDP.OutletID = newValue
			fields["outlet_id"] = "OutletID"

		case "status":
			newValue, err := cast.ToInt64E(val)
			if err != nil {
				continue
			}
			updatedDP.Status = newValue
			fields["status"] = "Status"
		}
	}

	if len(fields) > 0 {
		if data.Status != 0 && updatedDP.Status != 0 {
			return false, errors.New("[UpdateData] you can only edit draft data")
		}
		_, err = d.dpRepository.UpdateData(ctx, updatedDP, fields, dpID)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
