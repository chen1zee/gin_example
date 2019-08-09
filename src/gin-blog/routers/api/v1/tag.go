package v1

import (
	"fmt"
	"gin_example/src/gin-blog/models"
	"gin_example/src/gin-blog/pkg/e"
	"gin_example/src/gin-blog/pkg/setting"
	"gin_example/src/gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	state := -1
	if arg := c.Query("state"); arg != "" {
		state, _ = strconv.Atoi(arg)
		maps["state"] = state
	}

	code := e.Success

	fmt.Println(maps)

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 新增文章标签
func AddTag(c *gin.Context) {
	var (
		err       error
		state     int
		createdBy string
		name      string
	)
	name = c.Query("name")
	state, err = strconv.Atoi(c.DefaultQuery("state", "0"))
	if err != nil {
		state = 0
	}
	createdBy = c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为 100字")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为 100字")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.InvalidParams
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.Success
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ErrorExistTag
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 修改文章标签
func EditTag(c *gin.Context) {
	var (
		id         int
		name       string
		modifiedBy string
		state      = -1
	)
	id, _ = strconv.Atoi(c.Param("id"))
	name = c.Query("name")
	modifiedBy = c.Query("modified_by")

	valid := validation.Validation{}
	if arg := c.Query("state"); arg != "" {
		state, _ = strconv.Atoi(arg)
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为 100字")
	valid.MaxSize(name, 100, "name").Message("名称最长为 100字")

	code := e.InvalidParams
	if !valid.HasErrors() {
		code = e.Success
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		} else {
			code = e.ErrorNotExistTag
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 删除文章标签
func DeleteTag(c *gin.Context) {
	var (
		id   = -1
		code int
	)
	id, _ = strconv.Atoi(c.Param("id"))

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		code = e.InvalidParams
	} else if !models.ExistTagByID(id) {
		code = e.ErrorNotExistTag
	} else {
		models.DeleteTag(id)
		code = e.Success
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
