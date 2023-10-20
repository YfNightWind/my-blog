package validator

import (
	"fmt"
	"reflect"

	"github.com/YfNightWind/my-blog/utils/errormsg"
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	validator2 "github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
)

func Validate(data interface{}) (string, int) {
	validate := validator2.New()
	// 翻译为中文
	uni := ut.New(zh_Hans_CN.New())
	translator, _ := uni.GetTranslator("zh_Hans_CN")
	err := zh.RegisterDefaultTranslations(validate, translator)
	if err != nil {
		fmt.Println("err: ", err)
	}

	// 注册label
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	// 确定data传的是结构体
	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator2.ValidationErrors) {
			return v.Translate(translator), errormsg.ERROR
		}
	}

	return "", errormsg.SUCCESS
}
