package api

import (
	"gin_example/src/gin-blog/models"
	"gin_example/src/gin-blog/pkg/e"
	"gin_example/src/gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required"`
}

type reqBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetAuth(c *gin.Context) {
	var body reqBody
	_ = c.ShouldBind(&body)

	valid := validation.Validation{}
	a := auth{Username: body.Username, Password: body.Password}

	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})

	var code int
	var authResult *models.Auth
	var token string
	var err error
	if !ok {
		code = e.InvalidParams
		util.PrintValidError(valid.Errors)
	} else if authResult, ok = models.CheckAuth(body.Username, body.Password); !ok { // 不存在用户
		code = e.ErrorAuth
	} else if token, err = util.GenerateToken(body.Username, authResult.PubDesc); err != nil {
		// token生成错误
		code = e.ErrorAuthToken
	} else {
		data["token"] = token
		code = e.Success
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
