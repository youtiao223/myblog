package utils

import (
	"github.com/go-playground/locales/zh_Hans_CN"
	uniTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"myBlog/utils/errorUtils"
	"reflect"
)

// ValidateStruct 验证结构体表单
func ValidateStruct(data interface{}) (string, int) {
	// 实例化
	validate := validator.New()
	uni := uniTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	// 注册翻译方法
	_ = zhTrans.RegisterDefaultTranslations(validate, trans)

	// 通过反射获取struct中属性名
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err := validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errorUtils.ERROR
		}
	}
	return "", errorUtils.SUCCESS
}
