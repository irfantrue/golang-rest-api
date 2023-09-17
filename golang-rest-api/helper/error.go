package helper

import (
	"fmt"
	"golang-rest-api/model"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorProp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Errors  any    `json:"errors"`
}

type ErrorValidationPayload struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidatePayload(payload *model.PayloadLogin) []*ErrorValidationPayload {
	validate := validator.New()
	err := validate.Struct(payload)
	errorsArray := []*ErrorValidationPayload{}
	if err != nil {
		// Inisialisasi array objek untuk menyimpan pesan kesalahan

		// Loop melalui pesan kesalahan yang ditemukan
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			tag := err.Tag()

			var params string
			if err.Param() != "" {
				params = strings.Join([]string{tag, err.Param()}, " ")
			} else {
				params = tag
			}

			// Membuat pesan kesalahan dalam format yang diinginkan
			errorMessage := fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", fieldName, params)

			// Menambahkan pesan kesalahan ke dalam array objek
			errorsArray = append(errorsArray, &ErrorValidationPayload{
				Field:   fieldName,
				Message: errorMessage,
			})
		}
	}

	if len(errorsArray) == 0 {
		return nil
	}

	return errorsArray
}

func ResponseValidation(c *gin.Context, err []*ErrorValidationPayload) {
	lang := c.GetHeader("Accept-Language")

	var statusCode int
	var response interface{}

	statusCode, response = ErrorBody(lang, 400, &err)

	c.AbortWithStatusJSON(statusCode, response)
}

func ErrorMapping(err error) []*ErrorValidationPayload {
	errorsArray := []*ErrorValidationPayload{}

	if len(err.Error()) != 0 {
		errorsArray = append(errorsArray, &ErrorValidationPayload{
			Field:   "",
			Message: err.Error(),
		})
		return errorsArray
	}

	return nil
}

func ErrorBody(lang string, code int, data any) (int, *ErrorProp) {
	// Save error in database

	var response *ErrorProp

	if code < 400 || code > 599 {
		code = 500
	}

	response = &ErrorProp{
		Status:  false,
		Message: HttpMessage(lang, code),
		Errors:  &data,
	}

	return code, response
}

func ResponseError(c *gin.Context, code int, err error) {
	lang := c.GetHeader("Accept-Language")

	var statusCode int
	var response interface{}

	errs := ErrorMapping(err)
	statusCode, response = ErrorBody(lang, code, errs)

	c.AbortWithStatusJSON(statusCode, response)
}

func HandleError(c *gin.Context, code int, err error) {
	if err != nil {
		ResponseError(c, code, err)
	}
}
