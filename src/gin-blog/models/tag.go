package models

type Tag struct {
	Model
	Name       string `json:"name" example:"asd"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state" example:"0"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return tags
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return count
}

func ExistTagByName(name string) bool { // 根据 name 查询 tag 是否存在
	var tag Tag
	db.Select("id").Where("name=?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createdBy string) bool { // 添加标签
	db.Create(&Tag{
		Name:      name,
		CreatedBy: createdBy,
		State:     state,
	})
	return true
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id=?", id).First(&tag)
	return tag.ID > 0
}

func DeleteTag(id int) bool {
	db.Where("id=?", id).Delete(&Tag{})
	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id=?", id).Update(data)
	return true
}

func CleanAllTag() bool { // 清除所有 软删除的tags
	db.Unscoped().Where("deleted_on != ?", 0).Delete(&Tag{})
	return true
}
