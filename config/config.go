package config

import (
	"errors"
	"flag"
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

func Get(key string, value interface{}) error {
	val := reflect.ValueOf(value)
	if val.Kind() != reflect.Ptr {
		return errors.New("params value should be pointer")
	}
	val = val.Elem()
	if !val.CanAddr() {
		return errors.New("result must be addressable (a pointer)")
	}
	//rValue.Convert(rValue.Type())
	//rValue.Pointer()
	//ptr := reflect.ValueOf(value).Pointer()
	cfg, ok := configMap[key]
	if !ok {
		value = nil
		return errors.New("key[" + key + "] is not exits")
	}

	cp := cfg.Value
	val.Set(reflect.ValueOf(cp))
	return nil
}
func Set(key string, value interface{}) {
	configMap[key] = Cfg{
		Value:    value,
		Required: true,
	}
}

func Parse() {
	pwd, _ := os.Getwd()
	distPath := flag.String("d", pwd, "dist 目录的完整路径")
	flag.Parse()

	Set("distPath", *distPath)

	fromEnv("DB_DSN", "JIRA_BASIC_AUTH")

	for _, value := range configMap {
		if value.Required && value.Value == nil {
			panic(ErrConfigRequired)
		}
	}
}

func fromEnv(keys ...string) {
	for _, key := range keys {
		v := os.Getenv(key)
		if v != "" {
			Set(key, v)
		}
	}
}
