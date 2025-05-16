package i18n

import (
	_ "embed"
	"strings"
	"reflect"
	"encoding/json"
	"github.com/bd878/lesnotes_bot/internal/logger"
)

type translations struct {
	En map[string]string  `json:"en,omitempty"`
	Ru map[string]string  `json:"ru,omitempty"`
}

type declinations struct {
	En map[string][]string `json:"en,omitempty"`
	Ru map[string][]string `json:"ru,omitempty"`
}

var emptyDecl []string = make([]string, 3)

func (t translations) Get(code LangCode, key string) string {
	value := reflect.ValueOf(t)
	fieldValue := value.FieldByName(code.String())
	if !fieldValue.IsValid() {
		logger.Log.Errorw("field value is invalid", "value", fieldValue.String())
		return ""
	}

	dict, ok := fieldValue.Interface().(map[string]string)
	if !ok {
		logger.Log.Error("not ok")
		return ""
	}

	text, ok := dict[key]
	if !ok {
		logger.Log.Errorln("cannot convert field value to map[string]string")
		return ""
	}

	return text
}

func (d declinations) Get(code LangCode, key string) []string {
	value := reflect.ValueOf(d)
	fieldValue := value.FieldByName(code.String())
	if !fieldValue.IsValid() {
		logger.Log.Errorw("field value is invalid", "value", fieldValue.String())
		return emptyDecl
	}

	translations, ok := fieldValue.Interface().(map[string][]string)
	if !ok {
		logger.Log.Error("not ok")
		return emptyDecl
	}

	texts, ok := translations[key]
	if !ok {
		logger.Log.Errorln("cannot convert field value to map[string]string")
		return emptyDecl
	}

	return texts
}

//go:embed texts.json
var textsFile []byte
var texts translations

//go:embed decls.json
var declsFile []byte
var decls declinations

func init() {
	if err := json.Unmarshal(textsFile, &texts); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(declsFile, &decls); err != nil {
		panic(err)
	}
}

type Translator interface {
	Text(key string) string
	Decl(key string) string
}

type LangCode string

// reflected on translations/declinations struct keys
const (
	LangRu LangCode = "Ru"
	LangEn LangCode = "En"
	LangUnknown LangCode = ""
)

func LangFromString(code string) LangCode {
	switch code {
	case LangRu.String(), strings.ToLower(LangRu.String()):
		return LangRu
	case LangEn.String(), strings.ToLower(LangEn.String()):
		return LangEn
	default:
		return LangRu
	}
}

func (code LangCode) String() string { return string(code) }

func (code LangCode) Text(key string) string { return texts.Get(code, key) }
func (code LangCode) Decl(key string) []string { return decls.Get(code, key) }