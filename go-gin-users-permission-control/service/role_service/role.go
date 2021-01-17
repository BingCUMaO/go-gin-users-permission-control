package role_service

import "bingcu.top/go-gin-users-permission-control/models"

//用于CRUD所映射的对象，是一个CRUD的公有对象
type Role struct {
	Id				int			`json:"id"`
	Name			string		`json:"name"`
	Description		string		`json:"description"`
	Remark			string		`json:"remark"`

	Page 			int
	Size 			int
}

//添加一个角色
func (r *Role)Add()error  {
	role := map[string]interface{}{
		"id":			r.Id,
		"name":			r.Name,
		"description":	r.Description,
		"remark":		r.Remark,
	}

	if err := models.AddOneRole(role);err!=nil {
		return err
	}
	return nil
}

//分页查询角色列表
//返回分页查询到的角色和总共数量
func (r *Role)List()([]*models.Role, int64, error)  {
	roles, total, err := models.ListRoles(r.Page, r.Size)
	if err!=nil  {
		return nil,0, err
	}

	return roles, total, nil
}







