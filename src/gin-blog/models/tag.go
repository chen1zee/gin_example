package models

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return tags
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return count
}

func ExistTagByName(name string) bool {
	// TODO func ExistTagByName(name st
	//https://book.eddycjy.com/golang/gin/api-02.html#33-gin%E6%90%AD%E5%BB%BAblog-apis-%EF%BC%88%E4%BA%8C%EF%BC%89
	db.Select("id").Where("name=?", name).First(&Tag{})
}
