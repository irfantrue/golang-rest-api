package controller

import (
	"golang-rest-api/config"
	"golang-rest-api/helper"
	"golang-rest-api/model"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var payload *model.PayloadLogin

	if err := c.ShouldBindJSON(&payload); err != nil {
		helper.ResponseError(c, 400, err)
		return
	}

	// Validasi
	if err := helper.ValidatePayload(payload); err != nil {
		helper.ResponseValidation(c, err)
		return
	}

	user := []*model.User{
		{
			Email:    payload.Email,
			Password: payload.Username,
			Name:     config.Translate(c, "user.invalid.notexists"),
		},
		{
			Email:    payload.Email,
			Password: payload.Username,
			Name:     config.Translate(c, "user.invalid.notexists"),
		},
		{
			Email:    payload.Email,
			Password: payload.Username,
			Name:     config.Translate(c, "user.invalid.notexists"),
		},
		{
			Email:    payload.Email,
			Password: payload.Username,
			Name:     config.Translate(c, "user.invalid.notexists"),
		},
	}

	result := helper.Pagination(10, 100, 10, 1, &user)

	helper.ResponseSuccess(c, 201, result)
}
