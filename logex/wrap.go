package logex

import (
	"encoding/json"
	"reflect"
)

func Wrap(reqId string, moduleName string) Logger {
	logger := defaultLogger.Clone()
	logger.SetFormatter(func(v ...interface{}) (bytes []byte, e error) {
		return Format(reqId, moduleName, v...)
	})
	logger.DepthIncrease(-1)
	return logger
}

type UniversalLog struct {
	ReqId      string      `json:"req_id"`
	ModuleName string      `json:"module"`
	Info       interface{} `json:"info,omitempty"`
}

// 按照项目规范格式化代码
func Format(reqId string, moduleName string, v ...interface{}) ([]byte, error) {
	data := UniversalLog{ReqId: reqId, ModuleName: moduleName}
	for i, _ := range v {
		if errV, ok := v[i].(error); ok {
			v[i] = errV.Error()
		}
	}
	switch len(v) {
	case 0:
		// 啥都不做
	case 1:
		data.Info = v[0]
	default:
		data.Info = v
	}
	rv := reflect.ValueOf(data.Info)
	if rv.Kind() != reflect.String && data.Info != nil {
		b, err := json.Marshal(data.Info)
		if err != nil {
			Warningf("format log failed. req_id is %s, err is %w", reqId, err)
		}
		data.Info = string(b)
	}
	b, err := json.Marshal(data)
	if err != nil {
		Warningf("format log failed. req_id is %s, err is %w", reqId, err)
	}
	return b, err
}
