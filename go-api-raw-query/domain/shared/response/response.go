package response

import (
	"context"
	"fico_ar/domain/shared/constant"
	Shared "fico_ar/domain/shared/context"
	Error "fico_ar/domain/shared/error"
	"fico_ar/infrastructure/logger"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Data struct {
	Pagination Pagination  `json:"pagination"`
	Items      interface{} `json:"items"`
}
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    *Data  `json:"data"`
}

type Pagination struct {
	LimitPerPage int64 `json:"limit_per_page"`
	CurrentPage  int64 `json:"current_page"`
	TotalPage    int64 `json:"total_page"`
	TotalRows    int64 `json:"total_rows"`
	TotalItems   int64 `json:"total_items"`
	First        bool  `json:"first"`
	Last         bool  `json:"last"`
}

func ResponseOK(c *fiber.Ctx, msg string, data interface{}) error {
	logger.LogInfoWithData(data, constant.RESPONSE, msg)
	response := Response{
		Status:  constant.SUCCESS,
		Message: msg,
		Data: &Data{
			Items: data,
		},
	}

	return c.Status(http.StatusOK).JSON(response)
}

func ResponseOKWithPagination(c *fiber.Ctx, msg string, data Data) error {
	logger.LogInfoWithData(data, constant.RESPONSE, msg)
	response := Response{
		Status:  constant.SUCCESS,
		Message: msg,
		Data:    &data,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func ResponseError(c *fiber.Ctx, msg string, err error) error {
	response := Response{
		Status:  constant.ERROR,
		Message: fmt.Sprintf("%s: %s", msg, err.Error()),
		Data:    nil,
	}

	return c.Status(http.StatusBadGateway).JSON(response)
}

func ResponseErrorWithContext(ctx context.Context, err error) error {

	var (
		errType    string
		statusCode = http.StatusBadGateway
	)

	errType, err = Error.TrimMesssage(err)
	// Set Status Code
	if errType == constant.ErrDatabase || errType == constant.ErrTimeout {
		statusCode = http.StatusInternalServerError
	} else if errType == constant.ErrAuth {
		statusCode = http.StatusUnauthorized
	}

	logger.LogError(constant.RESPONSE, errType, err.Error())

	response := Response{
		Status:  constant.ERROR,
		Message: errType,
		Data:    nil,
	}

	c := Shared.GetValueFiberFromContext(ctx)

	return c.Status(statusCode).JSON(response)
}

// ErrorBadRequestFromError write response with error argument
func ResponseErrorBadRequest(c *fiber.Ctx, msg string, err error) error {
	response := Response{
		Status:  constant.ERROR,
		Message: fmt.Sprintf("%s: %s", msg, err.Error()),
		Data:    nil,
	}

	return c.Status(http.StatusBadRequest).JSON(response)
}
