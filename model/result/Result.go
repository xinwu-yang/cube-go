package result

import (
	"time"
)

type Result struct {
	Success   bool
	Message   string
	Code      int
	Result    interface{}
	Timestamp int64
}

func Ok(message string) Result {
	return Result{Success: true, Code: 200, Message: message, Timestamp: time.Now().Unix()}
}

func Err(message string) Result {
	return Result{Success: false, Code: 500, Message: message, Timestamp: time.Now().Unix()}
}
