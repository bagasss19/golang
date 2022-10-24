package downpayment

import (
	"fico_ar/domain/down_payment/feature"
	"fico_ar/domain/down_payment/model"
	"fico_ar/domain/shared/context"
	"fico_ar/domain/shared/response"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type DPHandler interface {
	GetDPList(c *fiber.Ctx) error
	GetDPDetailList(c *fiber.Ctx) error
	GetOneDP(c *fiber.Ctx) error
	GetOneDPDetail(c *fiber.Ctx) error
	CreateData(c *fiber.Ctx) error
	CreateDataDetail(c *fiber.Ctx) error
	DeleteDP(c *fiber.Ctx) error
	DeleteDPDetail(c *fiber.Ctx) error
	UpdateDP(c *fiber.Ctx) error
}

type dpHandler struct {
	dpFeature feature.DPFeature
}

func NewGiroHandler(dpFeature feature.DPFeature) DPHandler {
	return &dpHandler{
		dpFeature: dpFeature,
	}
}

// Get DP godoc
// @Summary      Get DP List
// @Description  show list of DP
// @Tags         DP
// @Param        page   query      string  false  "Page. Default is 1"
// @Param        limit   query      string  false  "Limit. Default is 5"
// @Router       /dp/list [get]
func (d dpHandler) GetDPList(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	filter := model.GetDPListPayload{
		Page:  int64(page),
		Limit: int64(limit),
	}

	resp, err := d.dpFeature.GetAllData(ctx, &filter)
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOKWithPagination(c, "OK!", resp)
}

// TODO: create DPDetailList godoc
func (d dpHandler) GetDPDetailList(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	filter := model.GetDPDetailListPayload{
		Page: int64(page),
		Limit: int64(limit),
	}

	resp, err := d.dpFeature.GetAllDataDetail(ctx, &filter)
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOKWithPagination(c, "OK!", resp)
}

// Get DP godoc
// @Summary      Get one DP
// @Description  show DP by ID
// @Tags         DP
// @Param        dp_id   query      string  true  "DP ID"
// @Router       /dp [get]
func (d dpHandler) GetOneDP(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	arID, _ := strconv.Atoi(c.Query("dp_id"))

	resp, err := d.dpFeature.GetOneData(ctx, int64(arID))
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOK(c, "OK!", resp)
}

// TODO: create GetOneDPDetail godoc
func (d dpHandler) GetOneDPDetail(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	dpDetailID, _ := strconv.Atoi(c.Params("dp_detail_id"))

	resp, err := d.dpFeature.GetOneDataDetail(ctx, int64(dpDetailID))
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOK(c, "OK!", resp)
}

// Create DP godoc
// @Summary      Create DP
// @Description  Create DP with dynamic fields
// @Tags         DP
// @Param        payload    body   model.DownPaymentRequest  true  "body payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /dp [post]
func (d dpHandler) CreateData(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	var request model.DownPaymentRequest
	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		return response.ResponseError(c, "Bad Request", err)
	}

	resp, err := d.dpFeature.CreateData(ctx, request)
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOK(c, "OK!", resp)
}

// TODO: create CreateDataDetail godoc
func (d dpHandler) CreateDataDetail(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	var request model.DownPaymentDetailRequest
	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		return response.ResponseError(c, "Bad Request", err)
	}

	resp, err := d.dpFeature.CreateDataDetail(ctx, request)
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOK(c, "OK!", resp)
}

// Delete DP godoc
// @Summary      Delete DP
// @Description  Delete DP by ID, only can delete DP with status 0 (Draft)
// @Tags         DP
// @Param        dp_id   query      string  false  "DP ID"
// @Router       /dp [delete]
func (d dpHandler) DeleteDP(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	dpID, err := strconv.Atoi(c.Query("dp_id"))
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "Bad Request", err)
	}
	err = d.dpFeature.DeleteData(ctx, int64(dpID))
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOK(c, "OK!", nil)
}

// TODO: create DeleteDPDetail godoc
func (d dpHandler) DeleteDPDetail(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	dpDetailID, err := strconv.Atoi(c.Params("dp_detail_id"))
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "Bad Request", err)
	}
	err = d.dpFeature.DeleteDataDetail(ctx, int64(dpDetailID))
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOK(c, "OK!", nil)
}

type DPUpdateRequestJson struct {
	Data map[string]interface{} `json:"data"`
}

// Update DP godoc
// @Summary      Update DP
// @Description  Update DP with dynamic fields, only can update DP with status 0 (Draft) except update status itself
// @Tags         DP
// @Param        dp_id   query      string  true  "Giro ID"
// @Param        payload    body   DPUpdateRequestJson  true  "enter desired field that want to update"
// @Router       /dp [patch]
func (d dpHandler) UpdateDP(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	var (
		requestJson DPUpdateRequestJson
		arID, _     = strconv.Atoi(c.Query("dp_id"))
	)

	if err := c.BodyParser(&requestJson); err != nil {
		log.Println(err)
		return response.ResponseError(c, "Bad Request", err)
	}

	updatedGiro := model.DPUpdatePayload{
		Data: requestJson.Data,
	}

	_, err := d.dpFeature.UpdateData(ctx, updatedGiro, int64(arID))
	if err != nil {
		log.Println(err)
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOK(c, "OK!", nil)
}
