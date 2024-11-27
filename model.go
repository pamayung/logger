package logger

// LogTdrModel or Transaction Data Record
type LogTdrModel struct {
	AppName    string `json:"app"`
	AppVersion string `json:"version"`
	ThreadID   string `json:"tid"`

	Path         string `json:"path"`
	Method       string `json:"method"`
	IP           string `json:"ip"`
	Port         int    `json:"port"`
	SrcIP        string `json:"srcIP"`
	RespTime     int64  `json:"latency"`
	ResponseCode string `json:"response_code"`

	Header   interface{} `json:"header"`
	Request  interface{} `json:"req"`
	Response interface{} `json:"resp"`
	Error    string      `json:"error"`

	AdditionalData interface{} `json:"addData"`
}
