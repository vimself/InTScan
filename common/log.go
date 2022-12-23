package common

import (
	"fmt"
	"strings"
	"time"
)

var LogSucTime int64
var LogErrTime int64
var WaitTime int64

func init() {
	LogSucTime = time.Now().Unix()
}

func LogError(errinfo interface{}) {
	if WaitTime == 0 {
		fmt.Println(fmt.Sprintf("%v", errinfo))
	} else if (time.Now().Unix()-LogSucTime) > WaitTime && (time.Now().Unix()-LogErrTime) > WaitTime {
		fmt.Println(fmt.Sprintf("%v", errinfo))
		LogErrTime = time.Now().Unix()
	}
}

func CheckErrs(err error) bool {
	if err == nil {
		return false
	}
	errs := []string{
		"closed by the remote host", "too many connections",
		"i/o timeout", "EOF", "A connection attempt failed",
		"established connection failed", "connection attempt failed",
		"Unable to read", "is not allowed to connect to this",
		"no pg_hba.conf entry",
		"No connection could be made",
		"invalid packet size",
		"bad connection",
	}
	for _, key := range errs {
		if strings.Contains(strings.ToLower(err.Error()), strings.ToLower(key)) {
			return true
		}
	}
	return false
}
