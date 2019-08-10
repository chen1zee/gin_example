package api

type auth struct {
	PubDesc  int    `valid:"Range(1,254)"`
	Username string `valid:"Required; MaxSize(50)"`
}

// TODO 在routers下的api目录新建auth.go文件，写入内容：
// TODO https://book.eddycjy.com/golang/gin/jwt.html
