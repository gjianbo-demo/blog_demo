package app

import (
	"gitee.com/gjianbo/web/global"
	"gitee.com/gjianbo/web/pkg/conver"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	page := conver.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		page = 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := conver.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSettingS.DefaultPageSize
	}

	if pageSize > global.AppSettingS.MaxPageSize {
		return global.AppSettingS.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pagesize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pagesize
	}

	return result
}
