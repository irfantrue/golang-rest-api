package helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type SuccessProp struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type Page struct {
	Limit     int `json:"limit"`
	TotalData int `json:"totalData"`
	TotalPage int `json:"totalPage"`
	Page      int `json:"page"`
}

type PaginateData struct {
	Page  *Page       `json:"page"`
	Datas interface{} `json:"datas"`
}

func Pagination(limit int, totalData int, totalPage int, page int, data interface{}) *PaginateData {
	result := &PaginateData{
		Page: &Page{
			Limit:     10,
			TotalData: 100,
			TotalPage: 10,
			Page:      1,
		},
		Datas: &data,
	}

	return result
}

func SuccessBody(lang string, code int, data interface{}) (int, *SuccessProp) {
	var response *SuccessProp

	if code >= 400 || code <= 599 {
		code = 200
	}

	response = &SuccessProp{
		Status:  true,
		Message: HttpMessage(lang, code),
		Result:  &data,
	}

	return code, response
}

func ResponseSuccess(c *gin.Context, code int, data interface{}) {
	lang := c.GetHeader("Accept-Language")

	var statusCode int
	var response interface{}

	statusCode, response = SuccessBody(lang, code, data)

	c.JSON(statusCode, response)
}

func MiddlewareResponse(c *gin.Context, code int) {
	lang := c.GetHeader("Accept-Language")

	response := &ErrorProp{
		Status:  false,
		Message: HttpMessage(lang, code),
		Errors:  nil,
	}

	c.AbortWithStatusJSON(code, response)
}

func NotFound(c *gin.Context) {
	lang := c.GetHeader("Accept-Language")

	response := &ErrorProp{
		Status:  false,
		Message: HttpMessage(lang, 404),
		Errors:  nil,
	}

	c.AbortWithStatusJSON(404, response)
}

func MethodNotAllowed(c *gin.Context) {
	lang := c.GetHeader("Accept-Language")

	response := &ErrorProp{
		Status:  false,
		Message: HttpMessage(lang, 405),
		Errors:  nil,
	}

	c.AbortWithStatusJSON(405, response)
}

func TooManyRequest(c *gin.Context) {
	lang := c.GetHeader("Accept-Language")

	response := &ErrorProp{
		Status:  false,
		Message: HttpMessage(lang, 429),
		Errors:  nil,
	}

	c.AbortWithStatusJSON(429, response)
}

func Panic(c *gin.Context) {
	lang := c.GetHeader("Accept-Language")

	defer func() {
		if r := recover(); r != nil {
			// Tangani panic di sini
			fmt.Println("Panic occurred:", r)

			response := &ErrorProp{
				Status:  false,
				Message: HttpMessage(lang, 500),
				Errors:  nil,
			}

			c.AbortWithStatusJSON(500, response)
		}
	}()

	c.Next()
}
