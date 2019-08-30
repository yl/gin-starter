package validators

import (
	"reflect"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/yangliulnn/gin-starter/httpd/validators/lang"
	enLang "github.com/yangliulnn/gin-starter/httpd/validators/lang/en"
	zhLang "github.com/yangliulnn/gin-starter/httpd/validators/lang/zh"
	"gopkg.in/go-playground/validator.v9"
)

type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var UT *ut.UniversalTranslator
var Validator *validator.Validate
var _ binding.StructValidator = &DefaultValidator{}

func (v *DefaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyInit()

		if err := v.validate.Struct(obj); err != nil {
			return error(err)
		}
	}

	return nil
}

func (v *DefaultValidator) Engine() interface{} {
	v.lazyInit()
	return v.validate
}

func (v *DefaultValidator) lazyInit() {
	v.once.Do(func() {
		Validator = validator.New()

		v.validate = Validator
		v.validate.SetTagName("binding")

		// add any custom validators etc. here
		//_ = v.Validate.RegisterValidation("unique", Unique)

		UT = ut.New(zh.New(), en.New())
		zhTrans, _ := UT.GetTranslator("zh")
		enTrans, _ := UT.GetTranslator("en")
		lang.RegisterTranslation(Validator, zhTrans, zhLang.Lang)
		lang.RegisterTranslation(Validator, enTrans, enLang.Lang)
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
