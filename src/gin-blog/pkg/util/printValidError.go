package util

import (
	"github.com/astaxie/beego/validation"
	"log"
)

/** 输出 valid Error */

func PrintValidError(errs []*validation.Error) {
	for _, err := range errs {
		log.Printf("err.key:%s, err.message:%s,", err.Key, err.Message)
	}
}
