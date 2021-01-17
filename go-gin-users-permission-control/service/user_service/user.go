package user_service

import "bingcu.top/go-gin-users-permission-control/models"

//用于CRUD所映射的对象，是一个CRUD的公有对象
type User struct {
	Id				int			`json:"id"`
	Username		string		`json:"username"`
	Password		string		`json:"password"`
	Email			string		`json:"email"`
	Phone			string		`json:"phone"`
	Sex				int8		`json:"sex"`
	Remark			string		`json:"remark"`

	Page 			int
	Size 			int
}

//添加一个用户
func (u *User)Add() error  {
	user := map[string]interface{}{
		"username":		u.Username,
		"password":		u.Password,
		"email":		u.Email,
		"phone":		u.Phone,
		"sex":			u.Sex,
		"remark":		u.Remark,
	}

	if err := models.AddOneUser(user); err != nil {
		return err
	}
	return nil
}

//分页查询用户列表
//返回分页查询到的用户和总共数量
func (u *User) List() ([]*models.User,int64, error) {
	user := map[string]interface{}{}
	//if u.Id != 0 		{ user["id"] 		= u.Id			}
	//if u.Username != "" { user["username"] 	= u.Username	}
	//if u.Password != "" { user["password"] 	= u.Password	}
	//if u.Email != "" 	{ user["email"] 	= u.Email		}
	//if u.Phone != "" 	{ user["phone"] 	= u.Phone		}
	//if u.Sex != 0 		{ user["sex"] 		= u.Sex			}
	//if u.Remark != "" 	{ user["remark"] 	= u.Remark		}

	users,total, err := models.ListUsers(u.Page, u.Size, user)
	if err!=nil  {
		return nil,0, err
	}
	return users,total, nil
}






