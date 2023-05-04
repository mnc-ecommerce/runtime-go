package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
)

var translator *ut.UniversalTranslator

var validate *validator.Validate

func init() {
	validate = validator.New()

	english := en.New()
	translator = ut.New(english, english)

	transEng, ok := translator.GetTranslator("en")
	if ok {
		if err := enTranslation.RegisterDefaultTranslations(validate, transEng); nil != err {
			panic(err)
		}
	}

	validate.RegisterTagNameFunc(
		func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		},
	)
}

func Get() *validator.Validate {
	return validate
}

func GetTranslator() *ut.UniversalTranslator {
	return translator
}
