package helpers


import (
	"github.com/asaskevich/govalidator"
)

type ResponseFormatter struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ApiResponseFormatter(status string, code int, message string, data interface{}) ResponseFormatter {
	meta := Meta{
		Status:  status,
		Code:    code,
		Message: message,
	}
	responseFormatter := ResponseFormatter{
		Meta: meta,
		Data: data,
	}
	return responseFormatter
}

func FormatError(err error) []string {
	var errors []string
	for _, e := range err.(govalidator.Errors).Errors() {
		errors = append(errors, e.Error())
	}

	return errors
}