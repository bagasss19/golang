package giro

import (
	"fico_ar/domain/giro/feature"
	"fico_ar/domain/giro/model"
	"fico_ar/domain/shared/context"
	"fico_ar/domain/shared/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type GiroHandler interface {
	GetGiroList(c *fiber.Ctx) error
	GetOneGiro(c *fiber.Ctx) error
	CreateData(c *fiber.Ctx) error
	DeleteGiro(c *fiber.Ctx) error
	UpdateGiro(c *fiber.Ctx) error
}

type giroHandler struct {
	giroFeature feature.GiroFeature
}

func NewGiroHandler(giroFeature feature.GiroFeature) GiroHandler {
	return &giroHandler{
		giroFeature: giroFeature,
	}
}

// Get Giro godoc
// @Summary      Get Giro List
// @Description  show list of Giro
// @Tags         Giro
// @Param        page   query      string  false  "Page. Default is 1"
// @Param        limit   query      string  false  "Limit. Default is 5"
// @Router       /giro/list [get]
func (g giroHandler) GetGiroList(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	filter := model.GetGiroListPayload{
		Page:  int64(page),
		Limit: int64(limit),
	}

	resp, err := g.giroFeature.GetAllData(ctx, &filter)
	if err != nil {
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOKWithPagination(c, "OK!", resp)
}

// Get Giro godoc
// @Summary      Get one Giro
// @Description  show Giro by ID
// @Tags         Giro
// @Param        giro_id   query      string  true  "Giro ID"
// @Router       /giro [get]
func (g giroHandler) GetOneGiro(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	arID, _ := strconv.Atoi(c.Query("giro_id"))

	resp, err := g.giroFeature.GetOneData(ctx, int64(arID))
	if err != nil {
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOK(c, "OK!", resp)
}

// Create Giro godoc
// @Summary      Create Giro
// @Description  Create Giro
// @Tags         Giro
// @Param        payload    body   model.GiroRequest  true  "body payload"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /giro [post]
func (g giroHandler) CreateData(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	var request model.GiroRequest
	if err := c.BodyParser(&request); err != nil {
		return response.ResponseError(c, "Bad Request", err)
	}

	err := g.giroFeature.CreateData(ctx, request)
	if err != nil {
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOK(c, "OK!", nil)
}

// Delete Giro godoc
// @Summary      Delete Giro
// @Description  Delete Giro by ID, only can delete Giro with status 0 (Draft)
// @Tags         Giro
// @Param        giro_id   query      string  false  "Giro ID"
// @Router       /giro [delete]
func (g giroHandler) DeleteGiro(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	giroID, err := strconv.Atoi(c.Query("giro_id"))
	if err != nil {
		return response.ResponseError(c, "Bad Request", err)
	}
	err = g.giroFeature.DeleteData(ctx, int64(giroID))
	if err != nil {
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOK(c, "OK!", nil)
}

type GiroUpdateRequestJson struct {
	Data map[string]interface{} `json:"data"`
}

// Update Giro godoc
// @Summary      Update Giro
// @Description  Update Giro with dynamic fields, only can update Giro with status 0 (Draft) except update status itself
// @Tags         Giro
// @Param        giro_id   query      string  true  "Giro ID"
// @Param        payload    body   GiroUpdateRequestJson  true  "enter desired field that want to update"
// @Router       /giro [patch]
func (g giroHandler) UpdateGiro(c *fiber.Ctx) error {
	ctx := context.CreateContext()
	var (
		requestJson GiroUpdateRequestJson
		arID, _     = strconv.Atoi(c.Query("giro_id"))
	)

	if err := c.BodyParser(&requestJson); err != nil {
		return response.ResponseError(c, "Bad Request", err)
	}

	updatedGiro := model.GiroUpdatePayload{
		Data: requestJson.Data,
	}

	_, err := g.giroFeature.UpdateData(ctx, updatedGiro, int64(arID))
	if err != nil {
		return response.ResponseError(c, "service error", err)
	}

	return response.ResponseOK(c, "OK!", nil)
}
