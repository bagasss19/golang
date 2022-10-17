package repository

import (
	"context"
	"database/sql"
	"fico_ar/domain/giro/model"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/fatih/structs"
)

func (g giroRepository) GetAllData(ctx context.Context, payload *model.GetGiroListPayload) (resp model.GetAllGiroResponse, err error) {
	query := `
	SELECT
		girocek_id,
		company_id,
		profit_center,
		bank_name,
		type,
		giro_date,
		giro_num,
		account_id,
		account_name,
		giro_amount,
		due_date,
		status,
		created_time,
		last_update,
		created_by,
		updated_by
	FROM
		ar_giro
	WHERE
		1=1`

	queryCount := `
	SELECT COUNT(*) 
	FROM
		ar_giro
	WHERE
		1=1`

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		var count int64

		rows, err := g.database.Query(queryCount)
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
		query += fmt.Sprintf(`ORDER BY "girocek_id" LIMIT %d OFFSET %d`, payload.Limit, payload.Offset)
		err = g.database.SelectContext(ctx, &resp.Data, query)
		if err != nil {
			err = fmt.Errorf("[GetAllData] failed when executed query. Error: %+v", err)
			return
		}
	}()

	wg.Wait()

	return
}

func (g giroRepository) CreateData(ctx context.Context, request model.Giro) (err error) {
	query := `
	INSERT INTO ar_giro(
		company_id,
		giro_date,
		profit_center,
		bank_name,
		type,
		giro_num,
		account_id,
		account_name,
		giro_amount,
		due_date,
		status,
		created_time,
		last_update,
		created_by,
		updated_by
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) RETURNING girocek_id`

	_, err = g.database.ExecContext(ctx, query,
		request.CompanyID,
		request.GiroDate,
		request.ProfitCenter,
		request.BankName,
		request.Type,
		request.GiroNum,
		request.AccountID,
		request.AccountName,
		request.GiroAmount,
		request.DueDate,
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

func (g giroRepository) GetOneData(ctx context.Context, giroID int64) (data model.Giro, err error) {
	query := `
	SELECT
		girocek_id,
		company_id,
		profit_center,
		bank_name,
		type,
		giro_date,
		giro_num,
		account_id,
		account_name,
		giro_amount,
		due_date,
		status,
		created_time,
		last_update,
		created_by,
		updated_by
	FROM
		ar_giro
	WHERE
		girocek_id = $1`

	err = g.database.GetContext(ctx, &data, query, giroID)
	if err != nil {
		if err == sql.ErrNoRows {
			return data, fmt.Errorf("[GetOneData] data not found!. Error: %+v", err)
		}
		return data, fmt.Errorf("[GetOneData] failed when executed query. Error: %+v", err)
	}
	return data, nil
}

func (g giroRepository) DeleteData(ctx context.Context, giroID int64) (err error) {
	query1 := `
	DELETE FROM
		ar_giro
	WHERE 
		girocek_id = $1`

	_, err = g.database.ExecContext(ctx, query1, giroID)
	if err != nil {
		return fmt.Errorf("[DeleteData] failed when executed query. Error: %+v", err)
	}

	return nil
}

func (g giroRepository) UpdateData(ctx context.Context, request model.GiroRequest, columns map[string]interface{}, giroID int64) (resp bool, err error) {
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
	query := "UPDATE ar_giro SET " + strings.Join(values, ",") + fmt.Sprintf(" where giro_id = %v", giroID)
	fmt.Println(query)

	_, err = g.database.ExecContext(ctx, query)
	if err != nil {
		return false, fmt.Errorf("[UpdateData] failed when executed query. Error: %+v", err)
	}

	return true, nil
}
