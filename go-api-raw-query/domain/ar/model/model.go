package model

import (
	"database/sql"
	"fico_ar/domain/shared/constant"
	"time"
)

type AR struct {
	ARID              int64        `db:"ar_id"`
	CompanyCode       string       `db:"company_code"`
	DocDate           sql.NullTime `db:"doc_date"`
	PostingDate       sql.NullTime `db:"posting_date"`
	SalesID           int64        `db:"sales_id"`
	OutletID          int64        `db:"outlet_id"`
	CollectorID       int64        `db:"collector_id"`
	BankID            int64        `db:"bank_id"`
	Description       string       `db:"description"`
	InvoiceNumber     int64        `db:"invoice_number"`
	InvoiceType       string       `db:"invoice_type"`
	InvoiceDate       sql.NullTime `db:"Invoice_date"`
	InvoiceValue      float64      `db:"invoice_value"`
	TotalPayment      float64      `db:"total_payment"`
	DiscountPayment   float64      `db:"discount_payment"`
	CashPayment       float64      `db:"cash_payment"`
	GiroNumber        int64        `db:"giro_number"`
	GiroPayment       float64      `db:"giro_payment"`
	TransferNumber    int64        `db:"transfer_number"`
	TransferPayment   float64      `db:"transfer_payment"`
	CNDNNumber        int64        `db:"cndn_number"`
	CNDNPayment       float64      `db:"cndn_payment"`
	ReturnNumber      int64        `db:"return_number"`
	ReturnPayment     float64      `db:"return_payment"`
	DownPaymentNumber int64        `db:"down_payment_number"`
	DownPayment       float64      `db:"down_payment"`
	Status            int64        `db:"status"`
	CreatedTime       sql.NullTime `db:"created_time"`
	LastUpdate        sql.NullTime `db:"last_update"`
	CreatedBy         string       `db:"created_by"`
	UpdatedBy         string       `db:"updated_by"`
}

type ARRequest struct {
	ARID              int64   `db:"ar_id"`
	CompanyCode       string  `db:"company_code"`
	DocDate           string  `db:"doc_date"`
	PostingDate       string  `db:"posting_date"`
	SalesID           int64   `db:"sales_id"`
	OutletID          int64   `db:"outlet_id"`
	CollectorID       int64   `db:"collector_id"`
	BankID            int64   `db:"bank_id"`
	Description       string  `db:"description"`
	InvoiceNumber     int64   `db:"invoice_number"`
	InvoiceType       string  `db:"invoice_type"`
	InvoiceDate       string  `db:"Invoice_date"`
	InvoiceValue      float64 `db:"invoice_value"`
	TotalPayment      float64 `db:"total_payment"`
	DiscountPayment   float64 `db:"discount_payment"`
	CashPayment       float64 `db:"cash_payment"`
	GiroNumber        int64   `db:"giro_number"`
	GiroPayment       float64 `db:"giro_payment"`
	TransferNumber    int64   `db:"transfer_number"`
	TransferPayment   float64 `db:"transfer_payment"`
	CNDNNumber        int64   `db:"cndn_number"`
	CNDNPayment       float64 `db:"cndn_payment"`
	ReturnNumber      int64   `db:"return_number"`
	ReturnPayment     float64 `db:"return_payment"`
	DownPaymentNumber int64   `db:"down_payment_number"`
	DownPayment       float64 `db:"down_payment"`
	Status            int64   `db:"status"`
	CreatedTime       string  `db:"created_time"`
	LastUpdate        string  `db:"last_update"`
	CreatedBy         string  `db:"created_by"`
	UpdatedBy         string  `db:"updated_by"`
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

	invoiceDate, err := time.Parse(constant.TimeFormat, request.InvoiceDate)
	if err != nil {
		return arModel, err
	}

	arModel.CompanyCode = request.CompanyCode
	arModel.InvoiceDate = sql.NullTime{
		Time:  invoiceDate,
		Valid: true,
	}
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
	arModel.CollectorID = request.CollectorID
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

type ARFilterList struct {
	CompanyCode string `db:"company_code" query:"company_code"`
	DocDate     string `db:"doc_date" query:"company_code"`
	PostingDate string `db:"posting_date" query:"posting_date"`
	SalesID     int64  `db:"sales_id" query:"sales_id"`
	OutletID    int64  `db:"outlet_id" query:"outlet_id"`
	CollectorID int64  `db:"collector_id" query:"collector_id"`
	BankID      int64  `db:"bank_id" query:"bank_id"`
	Description string `db:"description" query:"description"`
	Page        int64  `query:"page"`
	Limit       int64  `query:"limit"`
	Offset      int64  //payload for repository
}

type GetAllARResponse struct {
	Data      []AR
	TotalPage int64
	TotalItem int64
	Page      int64
	Size      int64
	First     bool
	Last      bool
}

type ARUpdate struct {
	InvoiceType       string  `json:"invoice_type"`
	InvoiceValue      float64 `json:"invoice_value"`
	TotalPayment      float64 `json:"total_payment"`
	DiscountPayment   float64 `json:"discount_payment"`
	CashPayment       float64 `json:"cash_payment"`
	GiroNumber        int64   `json:"giro_number"`
	CNDNNumber        int64   `json:"cndn_number"`
	CNDNPayment       float64 `json:"cndn_payment"`
	ReturnNumber      int64   `json:"return_number"`
	ReturnPayment     float64 `json:"return_payment"`
	DownPaymentNumber int64   `db:"down_payment_number"`
	DownPayment       float64 `db:"down_payment"`
	Status            int64   `json:"status"`
	UpdatedBy         string  `json:"updated_by"`
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
