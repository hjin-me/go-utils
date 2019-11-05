package logex

import "encoding/json"

func Wrap(reqId string, moduleName string) Logger {
	logger := defaultLogger.Clone()
	logger.SetFormatter(func(v ...interface{}) (bytes []byte, e error) {
		return Format(reqId, moduleName, v...)
	})
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
	switch len(v) {
	case 0:
		// 啥都不做
	case 1:
		data.Info = v[0]
	default:
		data.Info = v
	}
	b, err := json.Marshal(data)
	if err != nil {
		Warningf("format log failed. req_id is %s, err is %w", reqId, err)
	}
	return b, err
}
