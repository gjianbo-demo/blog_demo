package v1

import (
	"gitee.com/gjianbo/web/pkg/app"
	"gitee.com/gjianbo/web/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {

	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}

// @Summary 获取多个标题
// @Description 登录
// @Produce json
// @Param body body controllers.LoginParams true "body参数"
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
// @Router /user/person/login [post]
func (a Article) List(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}

func (a Article) Create(c *gin.Context) {

}

func (a Article) Update(c *gin.Context) {

}

func (a Article) Delete(c *gin.Context) {

}
