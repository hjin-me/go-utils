package cerror

import (
	"strings"
	"testing"
)

type TStruct struct {
	Required int `validate:"required" json:"test_required"`
	NoJSON   int `validate:"required"`
}

func TestParseValidateError(t *testing.T) {
	ts := TStruct{}
	err := Validate(ts)
	r := ParseValidateError(err)
	if !strings.Contains(r, "test_required") {
		t.Error(r, "json tag not work")
	}
	if !strings.Contains(r, "NoJSON") {
		t.Error(r, "no json tag should output StructField")
	}
}
