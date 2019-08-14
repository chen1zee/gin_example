package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	PubDesc  int    `json:"pub_desc"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username string, password string) (authResult *Auth, ok bool) {
	var auth Auth
	db.Select("id, pub_desc").Where(Auth{Username: username, Password: password}).First(&auth)
	return &auth, auth.ID > 0
}
