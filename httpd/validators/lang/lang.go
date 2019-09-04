package lang

import (
	"github.com/go-playground/universal-translator"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

type Fields map[string]string

type Rules map[string]string

type Lang struct {
	Fields Fields
	Rules  Rules
}

func transFunc(ut ut.Translator, fe validator.FieldError) string {
	fld, err := ut.T(fe.Field())
	if err != nil {
		log.Error("警告: 字段翻译错误: %#v", fe)
		fld = fe.Field()
	}
	t, err := ut.T(fe.Tag(), fld)
	if err != nil {
		log.Error("警告: 验证规则翻译错误: %#v", fe)
		return fe.(error).Error()
	}
	return t
}

func registerFieldTranslation(trans ut.Translator, fields Fields) {
	for field, translation := range fields {
		_ = trans.Add(field, translation, false)
	}
}

func registerRuleTranslation(validator *validator.Validate, trans ut.Translator, rules Rules) {
	for rule, translation := range rules {
		_ = validator.RegisterTranslation(rule, trans, func(ut ut.Translator) error {
			return ut.Add(rule, translation, false)
		}, transFunc)
	}
}

func RegisterTranslation(validator *validator.Validate, trans ut.Translator, lang Lang) {
	registerFieldTranslation(trans, lang.Fields)
	registerRuleTranslation(validator, trans, lang.Rules)
}
