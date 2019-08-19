package util

import (
	"gin_example/src/gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage(c *gin.Context) int {
	result := 0
	pageInt64, _ := strconv.ParseInt(c.Query("page"), 10, 0)
	page := int(pageInt64)
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}
	return result
}
