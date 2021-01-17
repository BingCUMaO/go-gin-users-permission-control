package permission_service

import "bingcu.top/go-gin-users-permission-control/models"

//用于CRUD所映射的对象，是一个CRUD的公有对象
type Permission struct {
	Id				int			`json:"id"`
	PermissionKey	string		`json:"permission_key"`
	Description 	string   	`json:"description"`
	Remark			string		`json:"remark"`

	Page 			int
	Size 			int
}

//添加一个权限类型
func (p *Permission)Add()error  {
	perms := map[string]interface{}{
		"permission_key": 		p.PermissionKey,
		"description": 			p.Description,
		"remark": 				p.Remark,
	}

	if err := models.AddOnePermission(perms); err != nil {
		return err
	}
	return nil
}


//分页查询权限类型
//返回分页查询到的权限类型和总共数量
func (p *Permission) List() ([]*models.Permission,int64, error) {
	ps,total, err := models.ListPermissions(p.Page, p.Size)
	if err!=nil  {
		return nil,0, err
	}
	return ps,total, nil
}





