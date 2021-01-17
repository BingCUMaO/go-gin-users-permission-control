package models

//数据库role_permission表所相映射的对象
type RolePermission struct {
	ID         			int 	`gorm:"primary_key" json:"id"`
	RoleId				int		`json:"role_id"`
	PermissionId		int		`json:"permission_id"`
}

//为指定角色添加一个权限
func AddRolePermission(data map[string]interface{})error {
	rp := RolePermission{
		RoleId: 		data["role_id"].(int),
		PermissionId: 	data["permission_id"].(int),
	}

	if err := db.Create(&rp).Error;err != nil {
		return err
	}
	return nil
}

//为指定角色删除对应的某个权限
func DeleteRolePermission(rolePermissionId int) error{
	err := db.Where("id=?", rolePermissionId).Delete(RolePermission{}).Error
	if err != nil {
		return err
	}
	return nil
}

//查询指定角色的所有权限
func ListRolePermissions(roleId int) ([]*RolePermission, error) {
	var rps []*RolePermission

	err := db.Where("role_id=?", roleId).Find(&rps).Error
	if err != nil {
		return nil, err
	}
	return rps,nil
}



