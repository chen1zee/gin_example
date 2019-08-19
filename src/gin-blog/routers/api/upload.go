package api

import (
	"gin_example/src/gin-blog/pkg/e"
	"gin_example/src/gin-blog/pkg/logging"
	"gin_example/src/gin-blog/pkg/upload"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

func responseCodeData(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func UploadImage(c *gin.Context) {
	var err error
	var file multipart.File
	var image *multipart.FileHeader

	data := make(map[string]string)
	file, image, err = c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		responseCodeData(c, e.Error, data)
		return
	}
	if image == nil {
		responseCodeData(c, e.InvalidParams, data)
		return
	}
	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()

	src := fullPath + imageName
	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		responseCodeData(c, e.ErrorUploadCheckImageFormat, data)
		return
	}
	if err = upload.CheckImage(fullPath); err != nil {
		logging.Warn(err)
		responseCodeData(c, e.ErrorUploadCheckImageFail, data)
		return
	}
	if err = c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		responseCodeData(c, e.ErrorUploadSaveImageFail, data)
		return
	}
	data["image_url"] = upload.GetImageFullUrl(imageName)
	data["image_save_url"] = savePath + imageName
	responseCodeData(c, e.Success, data)
}
