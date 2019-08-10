package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	PubDesc  int    `json:"pub_desc"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username string, pubDesc int) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, PubDesc: pubDesc}).First(&auth)
	return auth.ID > 0
}
