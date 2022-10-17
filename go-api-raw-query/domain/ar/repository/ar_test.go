package repository

import (
	"context"
	"fico_ar/domain/ar/model"
	"fico_ar/infrastructure/database"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func Test_arRepository_GetAllData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	assert.NoError(t, err)
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	query := `
	SELECT
		ar_account_receivable.ar_id,
		ar_account_receivable.company_id,
		ar_account_receivable.doc_number,
		ar_account_receivable.doc_date,
		ar_account_receivable.posting_date,
		ar_account_receivable.sales_id,
		ar_account_receivable.outlet_id,
		ar_account_receivable.collector_id,
		ar_account_receivable.bank_id,
		ar_account_receivable.description,
		ar_account_receivable.status,
		ar_account_receivable.created_time,
		ar_account_receivable.last_update,
		ar_account_receivable.created_by,
		ar_account_receivable.updated_by,
		ar_account_receivable_det.transaction_id,
		ar_account_receivable_det.invoice,
		ar_account_receivable_det.total_payment,
		ar_account_receivable_det.disc_payment,
		ar_account_receivable_det.cash_payment,
		ar_account_receivable_det.giro_num,
		ar_account_receivable_det.giro_amount,
		ar_account_receivable_det.transfer_num,
		ar_account_receivable_det.transfer_amount,
		ar_account_receivable_det.cndn_num,
		ar_account_receivable_det.cndn_amount,
		ar_account_receivable_det.return_num,
		ar_account_receivable_det.return_amount
	FROM
		ar_account_receivable
	FULL JOIN
		ar_account_receivable_det
	ON ar_account_receivable.ar_id = ar_account_receivable_det.ar_id
	WHERE
		1=1`

	queryCount := `
	SELECT COUNT(*) 
	FROM
		ar_account_receivable
	FULL JOIN
		ar_account_receivable_det
	ON ar_account_receivable.ar_id = ar_account_receivable_det.ar_id
		WHERE
			1=1`

	type fields struct {
		database *database.Database
	}
	type args struct {
		ctx     context.Context
		payload *model.ARFilterList
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData []model.ARResponse
		wantErr  bool
		mockFunc func(args)
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				database: &database.Database{sqlxDB},
			},
			args: args{
				ctx:     context.Background(),
				payload: &model.ARFilterList{},
			},
			wantErr: false,
			mockFunc: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(args.payload).WillReturnRows(sqlmock.NewRows([]string{"max"}).AddRow(5))
				mock.ExpectQuery(regexp.QuoteMeta(queryCount)).WithArgs(args.payload).WillReturnRows(sqlmock.NewRows([]string{"max"}).AddRow(5))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(tt.args)
			ar := &arRepository{
				database: tt.fields.database,
			}
			_, err := ar.GetAllData(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("arRepository.GetAllData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_arRepository_GetOneDataSales(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	assert.NoError(t, err)
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	query := `
	SELECT
		sales_id,
		company_code,
		gl_account,
		clearing_date,
		clearing_document,
		assignment,
		fiscal_year,
		document_number,
		line_item,
		posting_date,
		document_date,
		currency,
		reference,
		document_type,
		posting_period,
		posting_key,
		debit_credit,
		amount_in_local,
		amount,
		text,
		cost_center,
		baseline_payment_date,
		open_item,
		value_date,
		document_status,
		gl_currency,
		profit_center,
		gl_amount,
		clearing_year
	FROM
		ar_sales
	WHERE
		sales_id = $1`

	type fields struct {
		database *database.Database
	}
	type args struct {
		ctx     context.Context
		salesID int64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData model.ARSales
		wantErr  bool
		mockFunc func(args)
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				database: &database.Database{sqlxDB},
			},
			args: args{
				ctx:     context.Background(),
				salesID: 1,
			},
			wantErr: false,
			mockFunc: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(args.salesID).WillReturnRows(sqlmock.NewRows([]string{"sales_id"}).AddRow(5))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := arRepository{
				database: tt.fields.database,
			}
			tt.mockFunc(tt.args)
			_, err := ar.GetOneDataSales(tt.args.ctx, tt.args.salesID)
			if (err != nil) != tt.wantErr {
				t.Errorf("arRepository.GetOneDataSales() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
