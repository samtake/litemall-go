package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

//https://pkg.go.dev/github.com/go-playground/validator/v10#section-documentation
//自定义三个步骤：
//validator:1

func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	//使用正则表达式判断是否合法
	ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}
