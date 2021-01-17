package user_service

import "bingcu.top/go-gin-users-permission-control/models"

//用于CRUD所映射的对象，是一个CRUD的公有对象
type UserRole struct {
	ID         	int 		`json:"id"`
	UserId 		int 		`json:"user_id"`
	RoleId		int			`json:"role_id"`
}

//为指定用户添加一个角色
func (ur *UserRole) Add()error  {
	userRole := map[string]interface{}{
		"user_id":		ur.UserId,
		"role_id":		ur.RoleId,
	}

	if err := models.AddUserRole(userRole); err != nil {
		return err
	}
	return nil
}

//查询指定用户的所有角色
func (ur *UserRole) ListRoles()([]*models.UserRole, error)  {
	urs, err := models.ListUserRoles(ur.UserId)
	if err != nil {
		return nil, err
	}
	return urs,nil
}




