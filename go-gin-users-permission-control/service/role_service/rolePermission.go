package role_service

import "bingcu.top/go-gin-users-permission-control/models"

//用于CRUD所映射的对象，是一个CRUD的公有对象
type RolePermission struct {
	ID         			int 	`json:"id"`
	RoleId				int		`json:"role_id"`
	PermissionId		int		`json:"permission_id"`
}

//为指定角色添加一个权限
func (rp *RolePermission)Add() error {
	rolePermission := map[string]interface{}{
		"role_id":			rp.RoleId,
		"permission_id":	rp.PermissionId,
	}
	if err := models.AddRolePermission(rolePermission); err != nil {
		return err
	}
	return nil
}


//为指定角色删除对应的某个权限
func (p *RolePermission)Delete()error  {
	err := models.DeleteRolePermission(p.ID)
	if err!=nil {
		return err
	}
	return nil
}

//查询指定角色的所有权限
func (rp *RolePermission) ListPermissions()([]*models.RolePermission, error) {
	rps, err := models.ListRolePermissions(rp.RoleId)
	if err != nil {
		return nil, err
	}
	return rps,nil
}


