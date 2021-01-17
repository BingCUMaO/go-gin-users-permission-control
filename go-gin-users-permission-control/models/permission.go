package models

import "github.com/jinzhu/gorm"

//数据库permission表所相映射的对象
type Permission struct {
	Model

	PermissionKey	string 		`json:"permission_key"`
	Description 	string   	`json:"description"`
	Remark 			string 		`json:"remark"`
}

//添加一个权限类型
func AddOnePermission(data map[string]interface{}) error {
	permiss := Permission{
		PermissionKey:      data["permission_key"].(string),
		Description:      	data["description"].(string),
		Remark:      		data["remark"].(string),
	}

	if err := db.Create(&permiss).Error;err != nil {
		return err
	}
	return nil
}

//分页查询权限类型
//返回分页查询到的权限类型和总共数量
func ListPermissions(page int, size int) ([]*Permission,int64, error) {
	var ps []*Permission

	offset := (page-1)*size
	limit := size
	find := db.Offset(offset).Limit(limit).Find(&ps)
	rowsAffected := find.RowsAffected
	err:=find.Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil,0, err
	}

	return ps,rowsAffected, nil
}


