package model

import (
	"log"
	"net/http"
)

// Response data structure
type Response struct {
	Meta struct {
		Message string `json:"message"`
		Status  string `json:"status"`
		Code    int    `json:"code"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

// SuccessData write response
func (r *Response) SuccessData(data interface{}) (int, *Response) {
	r.Meta.Message = "data berhasil diproses"
	r.Meta.Status = "success"
	r.Meta.Code = http.StatusOK

	r.Data = data
	return http.StatusOK, r
}

// SuccessCreated write response
func (r *Response) SuccessCreated(data interface{}) (int, *Response) {
	r.Meta.Message = "data berhasil dibuat"
	r.Meta.Status = "success"
	r.Meta.Code = http.StatusCreated

	r.Data = data
	return http.StatusCreated, r
}

// SuccessUpdated write response
func (r *Response) SuccessUpdated(data interface{}) (int, *Response) {
	r.Meta.Message = "data berhasil diubah"
	r.Meta.Status = "success"
	r.Meta.Code = http.StatusOK

	r.Data = data
	return http.StatusOK, r
}

// ErrorBadRequest write response
func (r *Response) ErrorBadRequest(message string) (int, *Response) {
	r.Meta.Message = message
	r.Meta.Status = "Bad Request"
	r.Meta.Code = http.StatusBadRequest

	r.Data = nil
	return http.StatusBadRequest, r
}

// ErrorNotAuthorized write response
func (r *Response) ErrorNotAuthorized() (int, *Response) {
	r.Meta.Message = "tidak terautentikasi, silahkan login kembali"
	r.Meta.Status = "Not Authorized"
	r.Meta.Code = http.StatusUnauthorized

	r.Data = nil
	return http.StatusUnauthorized, r
}

// ErrorForbidden write response
func (r *Response) ErrorForbidden(message string) (int, *Response) {
	r.Meta.Message = message
	r.Meta.Status = "Forbidden"
	r.Meta.Code = http.StatusForbidden

	r.Data = nil
	return http.StatusForbidden, r
}

// ErrorDataNotFound write response
func (r *Response) ErrorDataNotFound() (int, *Response) {
	r.Meta.Message = "data tidak ditemukan"
	r.Meta.Status = "not found"
	r.Meta.Code = http.StatusNotFound

	r.Data = nil
	return http.StatusNotFound, r
}

// ErrorInternalServer write response
func (r *Response) ErrorInternalServer(err error) (int, *Response) {
	log.Printf("Error Response: %v \n", err)

	r.Meta.Message = err.Error()
	r.Meta.Status = "internal server error"
	r.Meta.Code = http.StatusInternalServerError

	r.Data = nil
	return http.StatusInternalServerError, r
}
