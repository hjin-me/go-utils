package config

import (
	"errors"
	"os"
	"reflect"
)

var configMap = make(map[string]Cfg)
var ErrConfigRequired = errors.New("missing some config")

type Cfg struct {
	Required bool
	Value    interface{}
}

func Require(key string) {
	cfg, ok := configMap[key]
	if !ok {
		cfg = Cfg{
			Required: true,
			Value:    nil,
		}
	}
	cfg.Required = true
	configMap[key] = cfg
}
func MustGet(key string, value interface{}) {
	err := Get(key, value)
	if err != nil {
		panic(err)
	}
}

func Get(key string, value interface{}) error {
	val := reflect.ValueOf(value)
	if val.Kind() != reflect.Ptr {
		return errors.New("params value should be pointer")
	}
	val = val.Elem()
	if !val.CanAddr() {
		return errors.New("result must be addressable (a pointer)")
	}
	cfg, ok := configMap[key]
	if !ok {
		val.Set(reflect.Zero(val.Type()))
		return errors.New("key[" + key + "] is not exits")
	}

	cp := cfg.Value
	if reflect.TypeOf(cp).AssignableTo(val.Type()) {
		val.Set(reflect.ValueOf(cp))
	} else {
		return errors.New("config's type [" + reflect.TypeOf(cp).Kind().String() + "] is not assignable to value's type [" +
			val.Type().Kind().String() + "]")
	}
	return nil
}
func Set(key string, value interface{}) {
	configMap[key] = Cfg{
		Value:    value,
		Required: true,
	}
}

func Parse() {
	missingKeys := make([]string, 0)
	for key, value := range configMap {
		if value.Required && value.Value == nil {
			missingKeys = append(missingKeys, key)
		}
	}
	if len(missingKeys) > 0 {
		panic(NewErr(missingKeys))
	}
}

func FromEnv(keys ...string) {
	for _, key := range keys {
		v := os.Getenv(key)
		if v != "" {
			Set(key, v)
		}
	}
}
