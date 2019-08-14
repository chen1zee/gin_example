package util

import (
	"gin_example/src/gin-blog/pkg/logging"
	"github.com/astaxie/beego/validation"
)

/** 输出 valid Error */

func PrintValidError(errs []*validation.Error) {
	for _, err := range errs {
		logging.Info(err.Key, err.Message)
	}
}
