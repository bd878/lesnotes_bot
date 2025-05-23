package models

import (
	"github.com/bd878/lesnotes_bot/internal/i18n"
)

type Chat struct {
	ID int64
	Token string
	Name string
	Lang i18n.LangCode
}