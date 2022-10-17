package model

import (
	"database/sql"
	"fico_ar/domain/shared/constant"
	"time"
)

type AR struct {
	ARID         int64        `db:"ar_id"`
	CompanyID    string       `db:"company_id"`
	DocNumber    int64        `db:"doc_number"`
	DocDate      sql.NullTime `db:"doc_date"`
	PostingDate  sql.NullTime `db:"posting_date"`
	SalesID      int64        `db:"sales_id"`
	OutletID     int64        `db:"outlet_id"`
	CollectiorID int64        `db:"collector_id"`
	BankID       int64        `db:"bank_id"`
	Description  string       `db:"description"`
	Status       int64        `db:"status"`
	CreatedTime  sql.NullTime `db:"created_time"`
	LastUpdate   sql.NullTime `db:"last_update"`
	CreatedBy    string       `db:"created_by"`
	UpdatedBy    string       `db:"updated_by"`
}

type ARRequest struct {
	CompanyID      string  `json:"company_id"`
	DocNumber      int64   `json:"doc_number"`
	DocDate        string  `json:"doc_date" example:"2020-12-19"`
	PostingDate    string  `json:"posting_date" example:"2020-12-19"`
	SalesID        int64   `json:"sales_id"`
	OutletID       int64   `json:"outlet_id"`
	CollectiorID   int64   `json:"collector_id"`
	BankID         int64   `json:"bank_id"`
	Description    string  `json:"description"`
	TransactionID  int64   `json:"transaction_id"`
	Invoice        string  `json:"invoice" example:"2020-12-19"`
	DiscPayment    float64 `json:"disc_payment"`
	CashPayment    float64 `json:"cash_payment"`
	GiroNumber     int64   `json:"giro_number"`
	GiroAmount     float64 `json:"giro_amount"`
	TransferNumber int64   `json:"transfer_number"`
	TransferAmount float64 `json:"transfer_amount"`
	CNDNNumber     int64   `json:"cndn_number"`
	CNDNAmount     float64 `json:"cndn_amount"`
	ReturnNumber   int64   `json:"return_number"`
	ReturnAmount   float64 `json:"return_amount"`
	Status         int64   `json:"status"`
	CreatedBy      string  `json:"created_by" example:"bagas"`
	CreatedTime    string  `db:"created_time"`
	LastUpdate     string  `db:"last_update"`
	UpdatedBy      string  `json:"updated_by"`
}

func NewAR(request ARRequest) (arModel AR, err error) {
	docDate, err := time.Parse(constant.TimeFormat, request.DocDate)
	if err != nil {
		return arModel, err
	}

	postDate, err := time.Parse(constant.TimeFormat, request.PostingDate)
	if err != nil {
		return arModel, err
	}

	arModel.CompanyID = request.CompanyID
	arModel.DocNumber = request.DocNumber
	arModel.DocDate = sql.NullTime{
		Time:  docDate,
		Valid: true,
	}
	arModel.PostingDate = sql.NullTime{
		Time:  postDate,
		Valid: true,
	}
	arModel.SalesID = request.SalesID
	arModel.OutletID = request.OutletID
	arModel.CollectiorID = request.CollectiorID
	arModel.BankID = request.BankID
	arModel.Description = request.Description
	arModel.Status = request.Status

	if arModel.CreatedTime.Time.IsZero() {
		arModel.CreatedTime = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}

	if request.LastUpdate == "" {
		arModel.LastUpdate = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	} else {
		lastUpdate, err := time.Parse(constant.TimeFormat, request.LastUpdate)
		if err != nil {
			return arModel, err
		}
		arModel.LastUpdate = sql.NullTime{
			Time:  lastUpdate,
			Valid: true,
		}
	}
	arModel.CreatedBy = request.CreatedBy
	arModel.UpdatedBy = request.UpdatedBy
	return arModel, nil
}

type ARDetail struct {
	ARDetID        int64        `db:"ar_det_id"`
	ARID           int64        `db:"ar_id"`
	TransactionID  int64        `db:"transaction_id"`
	Invoice        sql.NullTime `db:"invoice"`
	TotalPayment   float64      `db:"total_payment"`
	DiscPayment    float64      `db:"disc_payment"`
	CashPayment    float64      `db:"cash_payment"`
	GiroNumber     int64        `db:"giro_num"`
	GiroAmount     float64      `db:"giro_amount"`
	TransferNumber int64        `db:"transfer_num"`
	TransferAmount float64      `db:"transfer_amount"`
	CNDNNumber     int64        `db:"cndn_num"`
	CNDNAmount     float64      `db:"cndn_amount"`
	ReturnNumber   int64        `db:"return_num"`
	ReturnAmount   float64      `db:"return_amount"`
	Status         int64        `db:"status"`
}

func NewARDetail(request ARRequest) (arDetailModel ARDetail, err error) {
	arDetailModel.TransactionID = request.TransactionID
	invoice, err := time.Parse(constant.TimeFormat, request.Invoice)
	if err != nil {
		return arDetailModel, err
	}
	arDetailModel.Invoice = sql.NullTime{
		Time:  invoice,
		Valid: true,
	}
	totalPayment := request.DiscPayment + request.CashPayment + request.GiroAmount + request.CNDNAmount + request.ReturnAmount
	arDetailModel.TotalPayment = totalPayment
	arDetailModel.DiscPayment = request.DiscPayment
	arDetailModel.CashPayment = request.CashPayment
	arDetailModel.GiroNumber = request.GiroNumber
	arDetailModel.GiroAmount = request.GiroAmount
	arDetailModel.TransferNumber = request.TransferNumber
	arDetailModel.TransferAmount = request.TransferAmount
	arDetailModel.CNDNNumber = request.CNDNNumber
	arDetailModel.CNDNAmount = request.CNDNAmount
	arDetailModel.ReturnNumber = request.ReturnNumber
	arDetailModel.ReturnAmount = request.ReturnAmount
	arDetailModel.Status = request.Status
	return arDetailModel, nil
}

type ARResponse struct {
	ARID           int64        `db:"ar_id"`
	CompanyID      string       `db:"company_id"`
	DocNumber      int64        `db:"doc_number"`
	DocDate        sql.NullTime `db:"doc_date"`
	PostingDate    sql.NullTime `db:"posting_date"`
	SalesID        int64        `db:"sales_id"`
	OutletID       int64        `db:"outlet_id"`
	CollectiorID   int64        `db:"collector_id"`
	BankID         int64        `db:"bank_id"`
	Description    string       `db:"description"`
	ARDetID        int64        `db:"ar_det_id"`
	TransactionID  int64        `db:"transaction_id"`
	Invoice        sql.NullTime `db:"invoice"`
	TotalPayment   float64      `db:"total_payment"`
	DiscPayment    float64      `db:"disc_payment"`
	CashPayment    float64      `db:"cash_payment"`
	GiroNumber     int64        `db:"giro_num"`
	GiroAmount     float64      `db:"giro_amount"`
	TransferNumber int64        `db:"transfer_num"`
	TransferAmount float64      `db:"transfer_amount"`
	CNDNNumber     int64        `db:"cndn_num"`
	CNDNAmount     float64      `db:"cndn_amount"`
	ReturnNumber   int64        `db:"return_num"`
	ReturnAmount   float64      `db:"return_amount"`
	CreatedTime    sql.NullTime `db:"created_time"`
	LastUpdate     sql.NullTime `db:"last_update"`
	CreatedBy      string       `db:"created_by"`
	UpdatedBy      string       `db:"updated_by"`
	Status         int64        `db:"status"`
}

type ARFilterList struct {
	CompanyID   string `query:"company_id"`
	DocDate     string `query:"doc_date"`
	PostingDate string `query:"posting_date"`
	SalesID     int64  `query:"sales_id"`
	OutletID    int64  `query:"outlet_id"`
	CollectorID int64  `query:"collector_id"`
	BankID      int64  `query:"bank_id"`
	Description string `query:"description"`
	Page        int64  `query:"page"`
	Limit       int64  `query:"limit"`
	Offset      int64  //payload for repository
}

type GetAllARResponse struct {
	Data      []ARResponse
	TotalPage int64
	TotalItem int64
	Page      int64
	Size      int64
	First     bool
	Last      bool
}

type ARUpdate struct {
	TotalPayment   float64 `json:"total_payment"`
	DiscPayment    float64 `json:"disc_payment"`
	CashPayment    float64 `json:"cash_payment"`
	GiroNumber     int64   `json:"giro_number"`
	GiroAmount     float64 `json:"giro_amount"`
	TransferNumber int64   `json:"transfer_number"`
	TransferAmount float64 `json:"transfer_amount"`
	CNDNNumber     int64   `json:"cndn_number"`
	CNDNAmount     float64 `json:"cndn_amount"`
	ReturnNumber   int64   `json:"return_number"`
	ReturnAmount   float64 `json:"return_amount"`
	Status         int64   `json:"status"`
	UpdatedBy      string  `json:"updated_by"`
}

type ARUpdatePayload struct {
	Data map[string]interface{}
}

type ARSales struct {
	SalesID             int64        `db:"sales_id"`
	CompanyCode         string       `db:"company_code"`
	GLAccount           string       `db:"gl_account"`
	ClearingDate        sql.NullTime `db:"clearing_date"`
	ClearingDocument    string       `db:"clearing_document"`
	Assignment          string       `db:"assignment"`
	FiscalYear          int64        `db:"fiscal_year"`
	DocumentNumber      string       `db:"document_number"`
	LineItem            int64        `db:"line_item"`
	PostingDate         sql.NullTime `db:"posting_date"`
	DocumentDate        sql.NullTime `db:"document_date"`
	Currency            float64      `db:"currency"` //money?
	Reference           string       `db:"reference"`
	DocumentType        string       `db:"document_type"`
	PostingPeriod       int64        `db:"posting_period"`
	PostingKey          string       `db:"posting_key"`
	DebitCredit         string       `db:"debit_credit"`
	AmountInLocal       float64      `db:"amount_in_local"` //money?
	Amount              float64      `db:"amount"`          //money?
	Text                string       `db:"text"`
	CostCenter          string       `db:"cost_center"`
	BaselinePaymentDate sql.NullTime `db:"baseline_payment_date"`
	OpenItem            string       `db:"open_item"`
	ValueDate           sql.NullTime `db:"value_date"`
	DocumentStatus      string       `db:"document_status"`
	GLCurrency          string       `db:"gl_currency"` //chuck?
	ProfitCenter        string       `db:"profit_center"`
	GLAmount            float64      `db:"gl_amount"` //money?
	ClearingYear        int64        `db:"clearing_year"`
}
