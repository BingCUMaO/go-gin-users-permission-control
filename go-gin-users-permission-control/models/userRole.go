package models

//数据库user_role表所相映射的对象
type UserRole struct {
	ID         	int 	`gorm:"primary_key" json:"id"`
	UserId 		int 	`json:"user_id"`
	RoleId		int		`json:"role_id"`
}

//为指定用户添加一个角色
func AddUserRole(data map[string]interface{})error {
	ur := UserRole{
		UserId:     data["user_id"].(int),
		RoleId:     data["role_id"].(int),
	}

	if err := db.Create(&ur).Error;err != nil {
		return err
	}
	return nil
}

//查询指定用户的所有角色
func ListUserRoles(userId int) ([]*UserRole, error) {
	var urs []*UserRole

	err := db.Where("user_id=?", userId).Find(&urs).Error
	if err != nil {
		return nil, err
	}
	return urs,nil
}




