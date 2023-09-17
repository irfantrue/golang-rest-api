package config

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("locales/en.json")
	bundle.MustLoadMessageFile("locales/id.json")
}

func Translate(c *gin.Context, messageID string) string {
	lang := c.GetHeader("Accept-Language")
	localizer := i18n.NewLocalizer(bundle, lang)
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})
}
