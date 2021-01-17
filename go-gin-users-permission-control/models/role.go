package models

import "github.com/jinzhu/gorm"

//数据库role表所相映射的对象
type Role struct {
	Model

	Name 			string   	`json:"name"`
	Description 	string   	`json:"description"`
	Remark 			string 		`json:"remark"`
}

//添加一个角色
func AddOneRole(data map[string]interface{}) error {
	role := Role{
		Name: 			data["name"].(string),
		Description:	data["description"].(string),
		Remark: 		data["remark"].(string),
	}

	if err := db.Create(&role).Error;err!=nil {
		return err
	}
	return nil
}

//分页查询角色列表
//返回分页查询到的角色和总共数量
func ListRoles(page int, size int)([]*Role, int64, error) {
	var roles []*Role

	offset := (page-1)*size
	limit := size
	find := db.Offset(offset).Limit(limit).Find(&roles)
	RowsAffected := find.RowsAffected
	err := find.Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil,0, err
	}

	return roles, RowsAffected, nil
}


