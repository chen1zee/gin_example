package v1

import (
	"gin_example/src/gin-blog/models"
	"gin_example/src/gin-blog/pkg/e"
	"gin_example/src/gin-blog/pkg/setting"
	"gin_example/src/gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 获取单个文章
func GetArticle(c *gin.Context) {
	var (
		id   int
		code int
		data interface{}
	)
	id, _ = strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if valid.HasErrors() {
		code = e.InvalidParams
		util.PrintValidError(valid.Errors)
	} else if !models.ExistArticleByID(id) {
		code = e.ErrorNotExistArticle
	} else {
		data = models.GetArticle(id)
		code = e.Success
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 获取多个文章
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}
	var (
		state = -1
		tagId = -1
		code  int
	)
	if arg := c.Query("state"); arg != "" {
		state, _ = strconv.Atoi(arg)
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	if arg := c.Query("tag_id"); arg != "" {
		tagId, _ = strconv.Atoi(arg)
		maps["tag_id"] = tagId
		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}
	if valid.HasErrors() {
		code = e.InvalidParams
		util.PrintValidError(valid.Errors)
	} else {
		code = e.Success
		data["lists"] = models.GetArticles(util.GetPage(c), setting.AppSetting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// 新增文章
func AddArticle(c *gin.Context) {
	tagId, _ := strconv.Atoi(c.Query("tag_id"))
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state, _ := strconv.Atoi(c.DefaultQuery("state", "0"))

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	var code int
	if valid.HasErrors() {
		code = e.InvalidParams
		util.PrintValidError(valid.Errors)
	} else if !models.ExistTagByID(tagId) {
		code = e.ErrorNotExistTag
	} else {
		data := make(map[string]interface{})
		data["tag_id"] = tagId
		data["title"] = title
		data["desc"] = desc
		data["content"] = content
		data["created_by"] = createdBy
		data["state"] = state
		models.AddArticle(data)
		code = e.Success
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

// 修改文章
func EditArticle(c *gin.Context) {
	valid := validation.Validation{}

	id, _ := strconv.Atoi(c.Param("id"))
	tagId, _ := strconv.Atoi(c.Query("tag_id"))
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	state := -1
	if arg := c.Query("state"); arg != "" {
		state, _ = strconv.Atoi(arg)
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	var code int
	if valid.HasErrors() {
		code = e.InvalidParams
		util.PrintValidError(valid.Errors)
	} else if !models.ExistArticleByID(id) {
		code = e.ErrorNotExistArticle
	} else if !models.ExistTagByID(tagId) {
		code = e.ErrorNotExistTag
	} else {
		data := make(map[string]interface{})
		if tagId > 0 {
			data["tag_id"] = tagId
		}
		if title != "" {
			data["title"] = title
		}
		if desc != "" {
			data["desc"] = desc
		}
		if content != "" {
			data["content"] = content
		}
		data["modified_by"] = modifiedBy
		models.EditArticle(id, data)
		code = e.Success
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	var code int
	if valid.HasErrors() {
		code = e.InvalidParams
		util.PrintValidError(valid.Errors)
	} else if !models.ExistArticleByID(id) {
		code = e.ErrorNotExistArticle
	} else {
		models.DeleteArticle(id)
		code = e.Success
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
