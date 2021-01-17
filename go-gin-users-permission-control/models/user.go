package models

import "github.com/jinzhu/gorm"

//数据库user表所相映射的对象
type User struct {
	Model

	Username		string		`json:"username"`
	Password		string		`json:"password"`
	Email			string		`json:"email"`
	Phone			string		`json:"phone"`
	Sex				int8		`json:"sex"`
	Remark			string		`json:"remark"`
}

//添加一个用户
func AddOneUser(data map[string]interface{}) error {
	user := User{
		Username: 	data["username"].(string),
		Password: 	data["password"].(string),
		Email: 		data["email"].(string),
		Phone: 		data["phone"].(string),
		Sex: 		data["sex"].(int8),
		Remark: 	data["remark"].(string),
	}

	if err := db.Create(&user).Error;err != nil {
		return err
	}
	return nil
}

//分页查询用户列表
//返回分页查询到的用户和总共数量
func ListUsers(page int, size int, maps interface{}) ([]*User,int64, error) {
	var users []*User

	offset := (page-1)*size
	limit := size
	find := db.Offset(offset).Limit(limit).Find(&users)
	rowsAffected := find.RowsAffected
	err:=find.Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil,0, err
	}

	return users,rowsAffected, nil
}





