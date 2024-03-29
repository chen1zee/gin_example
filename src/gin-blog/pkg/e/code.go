package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	ErrorExistTag        = 10001
	ErrorNotExistTag     = 10002
	ErrorNotExistArticle = 10003

	ErrorAuthCheckTokenFail    = 20001
	ErrorAuthCheckTokenTimeout = 20002
	ErrorAuthToken             = 20003
	ErrorAuth                  = 20004

	ErrorUploadSaveImageFail    = 30001
	ErrorUploadCheckImageFail   = 30002
	ErrorUploadCheckImageFormat = 30003
)
