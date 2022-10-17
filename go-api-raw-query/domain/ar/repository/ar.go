package repository

import (
	"context"
	"database/sql"
	"fico_ar/domain/ar/model"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/fatih/structs"
)

func (ar arRepository) GetAllData(ctx context.Context, payload *model.ARFilterList) (resp model.GetAllARResponse, err error) {
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

	if payload.CompanyID != "" {
		query += fmt.Sprintf(`AND company_id = '%s'`, payload.CompanyID)
		queryCount += fmt.Sprintf(`AND company_id = '%s'`, payload.CompanyID)
	}

	if payload.PostingDate != "" {
		query += fmt.Sprintf(`AND date(posting_date) = '%s'`, payload.PostingDate)
		queryCount += fmt.Sprintf(`AND date(posting_date) = '%s'`, payload.PostingDate)
	}

	if payload.DocDate != "" {
		query += fmt.Sprintf(`AND date(doc_date) = '%s'`, payload.DocDate)
		queryCount += fmt.Sprintf(`AND date(doc_date) = '%s'`, payload.DocDate)
	}

	if payload.SalesID != 0 {
		query += fmt.Sprintf(`AND sales_id = '%d'`, payload.SalesID)
		queryCount += fmt.Sprintf(`AND sales_id = '%d'`, payload.SalesID)
	}

	if payload.OutletID != 0 {
		query += fmt.Sprintf(`AND outlet_id = '%v'`, payload.OutletID)
		queryCount += fmt.Sprintf(`AND outlet_id = '%v'`, payload.OutletID)
	}

	if payload.CollectorID != 0 {
		query += fmt.Sprintf(`AND collector_id = '%v'`, payload.CollectorID)
		queryCount += fmt.Sprintf(`AND collector_id = '%v'`, payload.CollectorID)
	}

	if payload.Description != "" {
		payload.Description = "%" + payload.Description + "%"
		query += fmt.Sprintf(`AND description LIKE '%s'`, payload.Description)
		queryCount += fmt.Sprintf(`AND description LIKE '%s'`, payload.Description)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		var count int64

		rows, err := ar.database.Query(queryCount)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			if err := rows.Scan(&count); err != nil {
				log.Fatal(err)
				return
			}
		}

		resp.TotalItem = count
	}()

	go func() {
		defer wg.Done()
		query += fmt.Sprintf(`ORDER BY "ar_id" LIMIT %d OFFSET %d`, payload.Limit, payload.Offset)
		err = ar.database.SelectContext(ctx, &resp.Data, query)
		if err != nil {
			err = fmt.Errorf("[GetAllData] failed when executed query. Error: %+v", err)
			return
		}
	}()

	wg.Wait()

	return
}

func (ar arRepository) GetOneDataSales(ctx context.Context, salesID int64) (data model.ARSales, err error) {
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

	err = ar.database.GetContext(ctx, &data, query, salesID)
	if err != nil {
		return data, fmt.Errorf("[GetOneDataSales] failed when executed query. Error: %+v", err)
	}
	return data, nil
}

func (ar arRepository) GetAllCompanyCode(ctx context.Context) (data []model.ARSales, err error) {
	query := `
	SELECT
		company_code
	FROM
		ar_sales`

	err = ar.database.SelectContext(ctx, &data, query)
	if err != nil {
		return nil, fmt.Errorf("[GetAllCompanyCode] failed when executed query. Error: %+v", err)
	}
	return data, nil
}

func (ar arRepository) CreateData(ctx context.Context, request model.AR) (arID int64, err error) {
	query := `
	INSERT INTO ar_account_receivable(
		company_id,
		doc_number,
		doc_date,
		posting_date,
		sales_id,
		outlet_id,
		collector_id,
		bank_id,
		description,
		status,
		created_time,
		last_update,
		created_by,
		updated_by
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) RETURNING ar_id`

	err = ar.database.GetContext(ctx, &arID, query,
		request.CompanyID,
		request.DocNumber,
		request.DocDate,
		request.PostingDate,
		request.SalesID,
		request.OutletID,
		request.CollectiorID,
		request.BankID,
		request.Description,
		request.Status,
		request.CreatedTime,
		request.LastUpdate,
		request.CreatedBy,
		request.UpdatedBy)
	if err != nil {
		return 0, fmt.Errorf("[CreateData] failed when executed query. Error: %+v", err)
	}

	return arID, nil
}

func (ar arRepository) CreateDataDetail(ctx context.Context, request model.ARDetail) (status bool, err error) {
	query := `
	INSERT INTO ar_account_receivable_det(
		ar_id,
		transaction_id,
		invoice,
		total_payment,
		disc_payment,
		cash_payment,
		giro_num,
		giro_amount,
		transfer_num,
		transfer_amount,
		cndn_num,
		cndn_amount,	
		return_num,
		return_amount,
		status
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)`

	_, err = ar.database.ExecContext(ctx, query,
		request.ARID,
		request.TransactionID,
		request.Invoice,
		request.TotalPayment,
		request.DiscPayment,
		request.CashPayment,
		request.GiroNumber,
		request.GiroAmount,
		request.TransferNumber,
		request.TransferAmount,
		request.CNDNNumber,
		request.CNDNAmount,
		request.ReturnNumber,
		request.ReturnAmount,
		request.Status)
	if err != nil {
		return false, fmt.Errorf("[CreateDataDetail] failed when executed query. Error: %+v", err)
	}

	return true, nil
}

func (ar arRepository) GetOneData(ctx context.Context, arID int64) (data model.ARResponse, err error) {
	query := `
	SELECT
		ar_id,
		status
	FROM
		ar_account_receivable
	WHERE
		ar_id = $1`

	err = ar.database.GetContext(ctx, &data, query, arID)
	if err != nil {
		if err == sql.ErrNoRows {
			return data, fmt.Errorf("[GetOneData] data not found!. Error: %+v", err)
		}
		return data, fmt.Errorf("[GetOneData] failed when executed query. Error: %+v", err)
	}
	return data, nil
}

func (ar arRepository) DeleteData(ctx context.Context, arID int64) (err error) {
	tx := ar.database.MustBegin()
	query1 := `
	DELETE FROM
		ar_account_receivable_det
	WHERE 
		ar_id = $1`

	query2 := `
		DELETE FROM
			ar_account_receivable
		WHERE 
			ar_id = $1`

	_, err = ar.database.ExecContext(ctx, query1, arID)
	if err != nil {
		return fmt.Errorf("[DeleteData] failed when executed query. Error: %+v", err)
	}

	tx.Commit()
	_, err = ar.database.ExecContext(ctx, query2, arID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[DeleteData] failed when executed query. Error: %+v", err)
	}

	return nil
}

func (ar arRepository) UpdateData(ctx context.Context, request model.ARUpdate, columns map[string]interface{}, arID int64) (resp bool, err error) {
	var (
		values []string
	)
	tx := ar.database.MustBegin()

	requestMap := structs.Map(request)

	for json, column := range columns {
		for key, value := range requestMap {
			if column == key && value != "" {
				values = append(values, fmt.Sprintf("%s = %v", json, value))
				continue
			}
		}
	}

	//joining string
	query := "UPDATE ar_account_receivable_det SET " + strings.Join(values, ",") + fmt.Sprintf(" where ar_id = %v", arID)
	fmt.Println(query)

	_, err = ar.database.ExecContext(ctx, query)
	if err != nil {
		return false, fmt.Errorf("[UpdateData] failed when executed query. Error: %+v", err)
	}

	tx.Commit()
	query2 := "UPDATE ar_account_receivable SET last_update=now()"
	_, err = ar.database.ExecContext(ctx, query2)
	if err != nil {
		return false, fmt.Errorf("[UpdateData] failed when executed query. Error: %+v", err)
	}

	return true, nil
}

func (ar arRepository) UpdateStatusData(ctx context.Context, status int64, arID int64) (err error) {
	tx := ar.database.MustBegin()
	query1 := `
	UPDATE
		ar_account_receivable_det
	SET
		status = $1
	WHERE 
		ar_id = $2`

	query2 := `
	UPDATE
		ar_account_receivable
	SET
		status = $1
	WHERE 
		ar_id = $2`

	_, err = ar.database.ExecContext(ctx, query1, status, arID)
	if err != nil {
		return fmt.Errorf("[UpdateStatusData] failed when executed query. Error: %+v", err)
	}

	tx.Commit()
	_, err = ar.database.ExecContext(ctx, query2, status, arID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[UpdateStatusData] failed when executed query. Error: %+v", err)
	}

	return nil
}
