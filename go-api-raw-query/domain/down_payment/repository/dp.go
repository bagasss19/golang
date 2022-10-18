package repository

import (
	"context"
	"database/sql"
	"fico_ar/domain/down_payment/model"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/fatih/structs"
)

func (dp dpRepository) GetAllData(ctx context.Context, payload *model.GetDPListPayload) (resp model.GetAllDPResponse, err error) {
	query := `
	SELECT
		dp_id,
		doc_number,
		doc_date,
		doc_type,
		doc,
		period,
		posting_date,
		company_id,
		currency_id,
		amount,
		reference,
		header_text,
		translation_date,
		taxreporting_date,
		trading_part,
		outlet_id,
		gl_id,
		trans_type_id,
		reason,
		status,
		created_time,
		last_update,
		created_by,
		updated_by
	FROM
		ar_down_payment
	WHERE
		1=1`

	queryCount := `
	SELECT COUNT(*) 
	FROM
		ar_down_payment
	WHERE
		1=1`

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		var count int64

		rows, err := dp.database.Query(queryCount)
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
		query += fmt.Sprintf(`ORDER BY "dp_id" LIMIT %d OFFSET %d`, payload.Limit, payload.Offset)
		err = dp.database.SelectContext(ctx, &resp.Data, query)
		if err != nil {
			err = fmt.Errorf("[GetAllData] failed when executed query. Error: %+v", err)
			return
		}
	}()

	wg.Wait()

	return
}

func (dp dpRepository) CreateData(ctx context.Context, request model.DownPayment) (err error) {
	query := `
	INSERT INTO ar_down_payment(
		doc_number,
		doc_date,
		doc_type,
		doc,
		period,
		posting_date,
		company_id,
		currency_id,
		amount,
		reference,
		header_text,
		translation_date,
		taxreporting_date,
		trading_part,
		outlet_id,
		gl_id,
		trans_type_id,
		reason,
		status,
		created_time,
		last_update,
		created_by,
		updated_by
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23) RETURNING dp_id`

	_, err = dp.database.ExecContext(ctx, query,
		request.DocNumber,
		request.DocDate,
		request.DocType,
		request.Doc,
		request.Period,
		request.PostingDate,
		request.CompanyID,
		request.CurrencyID,
		request.Amount,
		request.Reference,
		request.HeaderText,
		request.TranslationDate,
		request.TaxreportingDate,
		request.TradingPart,
		request.OutletID,
		request.GLID,
		request.TransTypeID,
		request.Reason,
		request.Status,
		request.CreatedTime,
		request.LastUpdate,
		request.CreatedBy,
		request.UpdatedBy)
	if err != nil {
		return fmt.Errorf("[CreateData] failed when executed query. Error: %+v", err)
	}

	return nil
}

func (dp dpRepository) CreateDataDetail(ctx context.Context, request model.DownPaymentDetail) (err error) {
	query := `
	INSERT INTO ar_dp_detail(
		dp_id
	)	VALUES ($1) RETURNING dp_detail_id
	`
	_, err = dp.database.ExecContext(ctx, query, request.DPID)
	if err != nil {
		return fmt.Errorf("[CreateDataDetail] failed when executed query. Error: %+v", err)
	}
	
	return nil
}

func (dp dpRepository) GetOneData(ctx context.Context, dpID int64) (data model.DownPayment, err error) {
	query := `
	SELECT
		dp_id,
		doc_number,
		doc_date,
		doc_type,
		doc,
		period,
		posting_date,
		company_id,
		currency_id,
		amount,
		reference,
		header_text,
		translation_date,
		taxreporting_date,
		trading_part,
		outlet_id,
		gl_id,
		trans_type_id,
		reason,
		status,
		created_time,
		last_update,
		created_by,
		updated_by
	FROM
		ar_down_payment
	WHERE
		dp_id = $1`

	err = dp.database.GetContext(ctx, &data, query, dpID)
	if err != nil {
		if err == sql.ErrNoRows {
			return data, fmt.Errorf("[GetOneData] data not found!. Error: %+v", err)
		}
		return data, fmt.Errorf("[GetOneData] failed when executed query. Error: %+v", err)
	}
	return data, nil
}

func (dp dpRepository) DeleteData(ctx context.Context, dpID int64) (err error) {
	query1 := `
	DELETE FROM
		ar_down_payment
	WHERE 
		dp_id = $1`

	_, err = dp.database.ExecContext(ctx, query1, dpID)
	if err != nil {
		return fmt.Errorf("[DeleteData] failed when executed query. Error: %+v", err)
	}

	return nil
}

func (dp dpRepository) UpdateData(ctx context.Context, request model.DownPaymentRequest, columns map[string]interface{}, arID int64) (resp bool, err error) {
	var (
		values []string
	)

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
	query := "UPDATE ar_down_payment SET " + strings.Join(values, ",") + fmt.Sprintf(" where ar_id = %v", arID)
	fmt.Println(query)

	_, err = dp.database.ExecContext(ctx, query)
	if err != nil {
		return false, fmt.Errorf("[UpdateData] failed when executed query. Error: %+v", err)
	}

	return true, nil
}
