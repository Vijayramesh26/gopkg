package common

const (
	SuccessCode = "S"
	ErrorCode   = "E"

	// Log levels
	INFO  = "INFO"
	ERROR = "ERROR"
	DEBUG = "DEBUG"

	// Database
	SELECT = "SELECT"
	INSERT = "INSERT"
	UPDATE = "UPDATE"
)

type ClientDetailsResp struct {
	DetailsArr any    `json:"respData"`
	Status     string `json:"status"`
	ErrMsg     string `json:"errMsg"`
}
