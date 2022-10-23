package ar

import (
	"fico_ar/domain/ar/feature"
	"fico_ar/domain/ar/model"
	"fico_ar/domain/shared/context"
	"fico_ar/domain/shared/response"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ArHandler interface {
	GetArList(c *fiber.Ctx) error
	GetOneAr(c *fiber.Ctx) error
	GetAllCompanyCode(c *fiber.Ctx) error
	DeleteAr(c *fiber.Ctx) error
	UpdateAr(c *fiber.Ctx) error
	UpdateArStatus(c *fiber.Ctx) error
}

type arHandler struct {
	arFeature feature.ArFeature
}

func NewArHandler(arFeature feature.ArFeature) ArHandler {
	return &arHandler{
		arFeature: arFeature,
	}
}

// Get AR godoc
// @Summary      Get AR List
// @Description  show list of AR
// @Tags         Account Receivable
// @Param        company_id   query      string  false  "Company ID"
// @Param        doc_date   query      string  false  "Doc Date"
// @Param        posting_date   query      string  false  "Posting Date"
// @Param        description   query      string  false  "Description"
// @Param        sales_id   query      string  false  "Sales ID"
// @Param        outlet_id   query      string  false  "Outlet ID"
// @Param        collector_id   query      string  false  "Collector ID"
// @Param        bank_id   query      string  false  "Bank ID"
// @Param        page   query      string  false  "Page. Default is 1"
// @Param        limit   query      string  false  "Limit. Default is 5"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /ar/list [get]
func (ah arHandler) GetArList(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	var filter model.ARFilterList
	// companyID := c.Query("company_id")
	// docDate := c.Query("doc_date")
	// postingDate := c.Query("posting_date")
	// description := c.Query("description")
	// salesID, _ := strconv.Atoi(c.Query("sales_id"))
	// outletID, _ := strconv.Atoi(c.Query("outlet_id"))
	// collectorID, _ := strconv.Atoi(c.Query("collector_id"))
	// bankID, _ := strconv.Atoi(c.Query("bank_id"))
	// page, _ := strconv.Atoi(c.Query("page"))
	// limit, _ := strconv.Atoi(c.Query("limit"))

	// filter := model.ARFilterList{
	// 	CompanyCode: companyID,
	// 	DocDate:     docDate,
	// 	PostingDate: postingDate,
	// 	Description: description,
	// 	SalesID:     int64(salesID),
	// 	OutletID:    int64(outletID),
	// 	CollectorID: int64(collectorID),
	// 	BankID:      int64(bankID),
	// 	Page:        int64(page),
	// 	Limit:       int64(limit),
	// }

	if err := c.QueryParser(filter); err != nil {
		log.Println(err)
		response.ResponseErrorBadRequest(c, "bad request, check your payload", err)
	}

	resp, err := ah.arFeature.GetAllData(ctx, &filter)
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOKWithPagination(c, "OK!", resp)
}

// Get AR godoc
// @Summary      Get one AR
// @Description  show AR by ID
// @Tags         Account Receivable
// @Param        ar_id   query      string  true  "AR ID"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /ar [get]
func (ah arHandler) GetOneAr(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	arID, err := strconv.Atoi(c.Query("ar_id"))

	if err != nil {
		log.Println(err)
		response.ResponseErrorBadRequest(c, "bad request, check your payload", err)
	}

	resp, err := ah.arFeature.GetOneData(ctx, int64(arID))
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOK(c, "OK!", resp)
}

// Get Company Code godoc
// @Summary      Get All Company Code
// @Description  Get All Company Code from sales table
// @Tags         Account Receivable
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /ar/company [get]
func (ah arHandler) GetAllCompanyCode(c *fiber.Ctx) error {
	ctx := context.CreateContext()

	resp, err := ah.arFeature.GetAllCompanyCode(ctx)
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOK(c, "OK!", resp)
}

// Delete AR godoc
// @Summary      Delete AR
// @Description  Delete AR by ID, only can delete AR with status 0 (Draft)
// @Tags         Account Receivable
// @Param        ar_id   query      string  false  "AR ID"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /ar [delete]
func (ah arHandler) DeleteAr(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	arID, err := strconv.Atoi(c.Query("ar_id"))
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "Bad Request", err)
	}
	err = ah.arFeature.DeleteData(ctx, int64(arID))
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOK(c, "OK!", nil)
}

type ARUpdateRequestJson struct {
	Data map[string]interface{} `json:"data"`
}

// Update AR godoc
// @Summary      Update AR
// @Description  Update AR with dynamic fields, only can update AR with status 0 (Draft)
// @Tags         Account Receivable
// @Param        ar_id   query      string  true  "AR ID"
// @Param        payload    body   ARUpdateRequestJson  true  "enter desired field that want to update"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /ar [patch]
func (ah arHandler) UpdateAr(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	var (
		requestJson ARUpdateRequestJson
		arID, _     = strconv.Atoi(c.Query("ar_id"))
	)

	if err := c.BodyParser(&requestJson); err != nil {
		log.Println(err)
		return response.ResponseError(c, "Bad Request", err)
	}

	updatedAR := model.ARUpdatePayload{
		Data: requestJson.Data,
	}

	_, err := ah.arFeature.UpdateData(ctx, updatedAR, int64(arID))
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOK(c, "OK!", nil)
}

type ARUpdateStatusRequestJson struct {
	Status int64 `json:"status"`
}

// Update AR Status godoc
// @Summary      Update AR Status
// @Description  Update AR status only, for other field use Update AR
// @Tags         Account Receivable
// @Param        ar_id   query      string  true  "AR ID"
// @Param        payload    body   ARUpdateStatusRequestJson  true  "status update"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /ar [put]
func (ah arHandler) UpdateArStatus(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	var (
		requestJson ARUpdateStatusRequestJson
		arID, _     = strconv.Atoi(c.Query("ar_id"))
	)

	if err := c.BodyParser(&requestJson); err != nil {
		log.Println(err)
		return response.ResponseError(c, "Bad Request", err)
	}

	status := requestJson.Status

	err := ah.arFeature.UpdateDataStatus(ctx, status, int64(arID))
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOK(c, "OK!", nil)
}
