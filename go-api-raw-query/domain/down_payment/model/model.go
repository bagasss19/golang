package model

import (
	"database/sql"
	"fico_ar/domain/shared/constant"
	"time"
)

type DownPayment struct {
	DPID             int64        `db:"dp_id"`
	DocNumber        int64        `db:"doc_number"`
	DocDate          sql.NullTime `db:"doc_date"`
	DocType          string       `db:"doc_type"`
	Doc              int64        `db:"doc"`
	Period           int64        `db:"period"`
	PostingDate      sql.NullTime `db:"posting_date"`
	CompanyID        string       `db:"company_id"`
	CurrencyID       int64        `db:"currency_id"`
	Amount           float64      `db:"amount"`
	Reference        string       `db:"reference"`
	HeaderText       string       `db:"header_text"`
	TranslationDate  sql.NullTime `db:"translation_date"`
	TaxreportingDate sql.NullTime `db:"taxreporting_date"`
	TradingPart      string       `db:"trading_part"`
	OutletID         int64        `db:"outlet_id"`
	GLID             int64        `db:"gl_id"`
	TransTypeID      int64        `db:"trans_type_id"`
	Reason           string       `db:"reason"`
	Status           int64        `db:"status"`
	CreatedTime      sql.NullTime `db:"created_time"`
	LastUpdate       sql.NullTime `db:"last_update"`
	CreatedBy        string       `db:"created_by"`
	UpdatedBy        string       `db:"updated_by"`
}

type DownPaymentRequest struct {
	DocNumber        int64   `json:"doc_number"`
	DocDate          string  `json:"doc_date" example:"2020-12-19"`
	DocType          string  `json:"doc_type"`
	Doc              int64   `json:"doc"`
	Period           int64   `json:"period"`
	PostingDate      string  `json:"posting_date" example:"2020-12-19"`
	CompanyID        string  `json:"company_id"`
	CurrencyID       int64   `json:"currency_id"`
	Amount           float64 `json:"amount"`
	Reference        string  `json:"reference"`
	HeaderText       string  `json:"header_text"`
	TranslationDate  string  `json:"translation_date" example:"2020-12-19"`
	TaxreportingDate string  `json:"taxreporting_date" example:"2020-12-19"`
	TradingPart      string  `json:"trading_part"`
	OutletID         int64   `json:"outlet_id"`
	GLID             int64   `json:"gl_id"`
	TransTypeID      int64   `json:"trans_type_id"`
	Reason           string  `json:"reason"`
	Status           int64   `json:"status"`
	CreatedTime      string  `json:"created_time" example:"2020-12-19"`
	LastUpdate       string  `json:"last_update" example:"2020-12-19"`
	CreatedBy        string  `json:"created_by"`
	UpdatedBy        string  `json:"updated_by"`
}

func NewDP(request DownPaymentRequest) (dpModel DownPayment, err error) {
	dpModel.DocNumber = request.DocNumber

	docDate, err := time.Parse(constant.TimeFormat, request.DocDate)
	if err != nil {
		return dpModel, err
	}
	dpModel.DocDate = sql.NullTime{
		Time:  docDate,
		Valid: true,
	}

	dpModel.DocType = request.DocType
	dpModel.Doc = request.Doc
	dpModel.Period = request.Period

	postingDate, err := time.Parse(constant.TimeFormat, request.PostingDate)
	if err != nil {
		return dpModel, err
	}

	dpModel.PostingDate = sql.NullTime{
		Time:  postingDate,
		Valid: true,
	}
	dpModel.CompanyID = request.CompanyID
	dpModel.CurrencyID = request.CurrencyID
	dpModel.Amount = request.Amount
	dpModel.Reference = request.Reference
	dpModel.HeaderText = request.HeaderText

	translationDate, err := time.Parse(constant.TimeFormat, request.TranslationDate)
	if err != nil {
		return dpModel, err
	}
	dpModel.TranslationDate = sql.NullTime{
		Time:  translationDate,
		Valid: true,
	}

	taxreportingDate, err := time.Parse(constant.TimeFormat, request.TaxreportingDate)
	if err != nil {
		return dpModel, err
	}
	dpModel.TaxreportingDate = sql.NullTime{
		Time:  taxreportingDate,
		Valid: true,
	}

	dpModel.TradingPart = request.TradingPart
	dpModel.OutletID = request.OutletID
	dpModel.GLID = request.GLID
	dpModel.TransTypeID = request.TransTypeID
	dpModel.Reason = request.Reason
	dpModel.Status = request.Status
	if dpModel.CreatedTime.Time.IsZero() {
		dpModel.CreatedTime = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}

	if request.LastUpdate == "" {
		dpModel.LastUpdate = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	} else {
		lastUpdate, err := time.Parse(constant.TimeFormat, request.LastUpdate)
		if err != nil {
			return dpModel, err
		}
		dpModel.LastUpdate = sql.NullTime{
			Time:  lastUpdate,
			Valid: true,
		}
	}
	dpModel.CreatedBy = request.CreatedBy
	dpModel.UpdatedBy = request.UpdatedBy
	return dpModel, nil
}

type DownPaymentDetail struct {
	DPDetailID   int64        `db:"dp_detail_id"`
	DPID         int64        `db:"dp_id"`
	AmountInLoc  float64      `db:"amount_in_loc"`
	AmountInDoc  float64      `db:"amount_in_doc"`
	PPNCode      int64        `db:"ppn_code"`
	TaxAmount    float64      `db:"tax_amount"`
	PONumber     int64        `db:"po_number"`
	POItem       int64        `db:"po_item"`
	Assign       string       `db:"assign"`
	PaymentBlock int64        `db:"payment_block"`
	PaymentMet   int64        `db:"payment_met"`
	PaymentMeet  int64        `db:"payment_meet"`
	ProfitID     string       `db:"profit_id"`
	DueOn        sql.NullTime `db:"due_on"`
	Orders       string       `db:"orders"`
	Reason       string       `db:"reason"`
	Status       int64        `db:"status"`
	CreatedTime  sql.NullTime `db:"created_time"`
	LastUpdate   sql.NullTime `db:"last_update"`
	CreatedBy    string       `db:"created_by"`
	UpdatedBy    string       `db:"updated_by"`
}

type DownPaymentDetailRequest struct {
	DPID         int64   `json:"dp_id"`
	AmountInLoc  float64 `json:"amount_in_loc"`
	AmountInDoc  float64 `json:"amount_in_doc"`
	PPNCode      int64   `json:"ppn_code"`
	TaxAmount    float64 `json:"tax_amount"`
	PONumber     int64   `json:"po_number"`
	POItem       int64   `json:"po_item"`
	Assign       string  `json:"assign"`
	PaymentBlock int64   `json:"payment_block"`
	PaymentMet   int64   `json:"payment_met"`
	PaymentMeet  int64   `json:"payment_meet"`
	ProfitID     string  `json:"profit_id"`
	DueOn        string  `json:"due_on" example:"2020-12-19"`
	Orders       string  `json:"orders"`
	Reason       string  `json:"reason"`
	Status       int64   `json:"status"`
	CreatedTime  string  `json:"created_time" example:"2020-12-19"`
	LastUpdate   string  `json:"last_update" example:"2020-12-19"`
	CreatedBy    string  `josn:"created_by"`
	UpdatedBy    string  `json:"updated_by"`
}

func NewDPDetail(request DownPaymentDetailRequest) (dpModel DownPaymentDetail, err error) {
	dpModel.DPID = request.DPID
	dpModel.AmountInDoc = request.AmountInDoc
	dpModel.AmountInLoc = request.AmountInLoc
	dpModel.PPNCode = request.PPNCode
	dpModel.TaxAmount = request.TaxAmount
	dpModel.PONumber = request.PONumber
	dpModel.POItem = request.POItem
	dpModel.Assign = request.Assign
	dpModel.PaymentBlock = request.PaymentBlock
	dpModel.PaymentMet = request.PaymentMet
	dpModel.PaymentMeet = request.PaymentMeet
	dpModel.ProfitID = request.ProfitID

	dueOn, err := time.Parse(constant.TimeFormat, request.DueOn)
	if err != nil {
		return dpModel, err
	}
	dpModel.DueOn = sql.NullTime{
		Time:  dueOn,
		Valid: true,
	}

	dpModel.Orders = request.Orders
	dpModel.Reason = request.Reason
	dpModel.Status = request.Status
	if dpModel.CreatedTime.Time.IsZero() {
		dpModel.CreatedTime = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}

	if request.LastUpdate == "" {
		dpModel.LastUpdate = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	} else {
		lastUpdate, err := time.Parse(constant.TimeFormat, request.LastUpdate)
		if err != nil {
			return dpModel, err
		}
		dpModel.LastUpdate = sql.NullTime{
			Time:  lastUpdate,
			Valid: true,
		}
	}
	dpModel.CreatedBy = request.CreatedBy
	dpModel.UpdatedBy = request.UpdatedBy
	return dpModel, nil
}

type DPUpdatePayload struct {
	Data map[string]interface{}
}

type GetAllDPResponse struct {
	Data      []DownPayment
	TotalPage int64
	TotalItem int64
	Page      int64
	Size      int64
	First     bool
	Last      bool
}

type GetDPListPayload struct {
	Page   int64 `query:"page"`
	Limit  int64 `query:"limit"`
	Offset int64 //payload for repository
}
