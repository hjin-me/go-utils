package cerror

import (
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"strings"
)

var validate = validator.New()

func Validate(v interface{}) error {
	return validate.Struct(v)
}

func init() {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

}
