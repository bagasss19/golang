package model

import (
	"database/sql"
	"fico_ar/domain/shared/constant"
	"time"
)

type Giro struct {
	GiroCekID    int64        `db:"girocek_id"`
	CompanyID    string       `db:"company_id"`
	GiroDate     sql.NullTime `db:"giro_date"`
	GiroNum      int64        `db:"giro_num"`
	AccountID    string       `db:"account_id"`
	AccountName  string       `db:"account_name"`
	GiroAmount   float64      `db:"giro_amount"`
	DueDate      sql.NullTime `db:"due_date"`
	ProfitCenter string       `db:"profit_center"`
	BankName     string       `db:"bank_name"`
	Type         string       `db:"type"`
	Status       int64        `db:"status"`
	CreatedTime  sql.NullTime `db:"created_time"`
	LastUpdate   sql.NullTime `db:"last_update"`
	CreatedBy    string       `db:"created_by"`
	UpdatedBy    string       `db:"updated_by"`
}

func NewGiro(request GiroRequest) (giroModel Giro, err error) {
	giroDate, err := time.Parse(constant.TimeFormat, request.GiroDate)
	if err != nil {
		return giroModel, err
	}

	dueDate, err := time.Parse(constant.TimeFormat, request.DueDate)
	if err != nil {
		return giroModel, err
	}

	giroModel.CompanyID = request.CompanyID
	giroModel.GiroDate = sql.NullTime{
		Time:  giroDate,
		Valid: true,
	}
	giroModel.GiroNum = request.GiroNum
	giroModel.AccountID = request.AccountID
	giroModel.AccountName = request.AccountName
	giroModel.GiroAmount = request.GiroAmount
	giroModel.DueDate = sql.NullTime{
		Time:  dueDate,
		Valid: true,
	}
	giroModel.ProfitCenter = request.ProfitCenter
	giroModel.BankName = request.BankName
	giroModel.Type = request.Type
	giroModel.Status = request.Status
	if giroModel.CreatedTime.Time.IsZero() {
		giroModel.CreatedTime = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}

	if request.LastUpdate == "" {
		giroModel.LastUpdate = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	} else {
		lastUpdate, err := time.Parse(constant.TimeFormat, request.LastUpdate)
		if err != nil {
			return giroModel, err
		}
		giroModel.LastUpdate = sql.NullTime{
			Time:  lastUpdate,
			Valid: true,
		}
	}
	giroModel.CreatedBy = request.CreatedBy
	giroModel.UpdatedBy = request.UpdatedBy
	return giroModel, nil
}

type GiroRequest struct {
	CompanyID    string  `json:"company_id"`
	GiroDate     string  `json:"giro_date" example:"2020-12-19"`
	GiroNum      int64   `json:"giro_num"`
	AccountID    string  `json:"account_id"`
	AccountName  string  `json:"account_name"`
	GiroAmount   float64 `json:"giro_amount"`
	ProfitCenter string  `json:"profit_center"`
	BankName     string  `json:"bank_name"`
	Type         string  `json:"type"`
	DueDate      string  `json:"due_date" example:"2020-12-19"`
	Status       int64   `json:"status"`
	CreatedTime  string  `json:"created_time" example:"2020-12-19"`
	LastUpdate   string  `json:"last_update" example:"2020-12-19"`
	CreatedBy    string  `json:"created_by"`
	UpdatedBy    string  `json:"updated_by"`
}

type GiroUpdatePayload struct {
	Data map[string]interface{}
}

type GetAllGiroResponse struct {
	Data      []Giro
	TotalPage int64
	TotalItem int64
	Page      int64
	Size      int64
	First     bool
	Last      bool
}

type GetGiroListPayload struct {
	Page   int64 `query:"page"`
	Limit  int64 `query:"limit"`
	Offset int64 //payload for repository
}
