package config

import (
	"errors"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/spf13/viper"
)

type ValidatorConfig struct {
	Validate   *validator.Validate
	Translator ut.Translator
}

func NewValidator(viper *viper.Viper) *ValidatorConfig {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
	validator := &ValidatorConfig{
		Validate:   validate,
		Translator: trans,
	}
	return validator
}

func (c *ValidatorConfig) ValidateStruct(s interface{}) error {
	err := c.Validate.Struct(s)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		firstErr := errors.New(errs[0].Translate(c.Translator))
		return firstErr
	}
	return nil
}
