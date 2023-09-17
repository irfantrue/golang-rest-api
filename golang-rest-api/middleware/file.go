package middleware

import (
	"errors"
	"fmt"
	"golang-rest-api/config"
	"golang-rest-api/helper"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AllowExtension(c *gin.Context, ext string) (interface{}, error) {
	if ext != "webp" {
		return nil, errors.New(config.Translate(c, "upload.invalid.image.ext"))
	}
	return nil, nil
}

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")

	if err != nil {
		helper.ResponseError(c, 400, err)
		return
	}

	ext := strings.Split(file.Filename, ".")[1]
	_, notAllowedExt := AllowExtension(c, ext)

	if notAllowedExt != nil {
		helper.ResponseError(c, 400, notAllowedExt)
		return
	}

	newUUID, errUUID := uuid.NewUUID()
	if errUUID != nil {
		helper.ResponseError(c, 500, errUUID)
		return
	}

	destination := fmt.Sprintf("web/%s.%s", newUUID, ext)
	errorMove := c.SaveUploadedFile(file, destination)
	if errorMove != nil {
		helper.ResponseError(c, 500, errorMove)
		return
	}

	helper.ResponseSuccess(c, 200, file.Filename)
}
