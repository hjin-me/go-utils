package config

import (
	"strings"
)

type Err struct {
	missingKeys []string
}

func (e *Err) Error() string {
	return "config keys not set: " + strings.Join(e.missingKeys, ", ")
}

func NewErr(missingKeys []string) *Err {
	return &Err{missingKeys: missingKeys}
}
