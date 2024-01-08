package web

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/mnc-ecommerce/runtime-go"
	v "github.com/mnc-ecommerce/runtime-go/validator"
)

type Response struct {
	Data       any        `json:"data,omitempty"`
	Errors     []Error    `json:"errors,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

func Err(err error) Response {
	var response Response
	if vErrs, ok := errors.Cause(err).(validator.ValidationErrors); ok {
		translator, _ := v.GetTranslator().GetTranslator("en")
		for _, vErr := range vErrs {
			response.Errors = append(
				response.Errors, Error{
					Code:    0,
					Message: vErr.Translate(translator),
				},
			)
		}
	}
	if vErr, ok := errors.Cause(err).(runtime.Error); ok {
		response.Errors = append(
			response.Errors, Error{
				Code:    vErr.Code,
				Message: vErr.Message,
			},
		)
	}

	response.Errors = append(
		response.Errors, Error{
			Message: err.Error(),
		},
	)

	return response
}

func Success(data any) Response {
	return Response{
		Data: data,
	}
}

func Collection(data any, pagination Pagination) Response {
	return Response{
		Data:       data,
		Pagination: pagination,
	}
}
