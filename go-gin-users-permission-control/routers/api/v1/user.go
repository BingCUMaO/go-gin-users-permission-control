package v1

import (
	"bingcu.top/go-gin-users-permission-control/pkg/app"
	"bingcu.top/go-gin-users-permission-control/pkg/e"
	"bingcu.top/go-gin-users-permission-control/service/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddUserForm struct {
	Username		string		`json:"username" from:"username" valid:"Required;"`
	Password		string		`json:"password" from:"password" valid:"Required;"`
	Email			string		`json:"email" from:"email"`
	Phone			string		`json:"phone" from:"phone"`
	Sex				int8		`json:"sex" from:"sex"`
	Remark			string		`json:"remark" from:"remark"`
}

type AddUserRoleForm struct {
	UserId 			int			`json:"user_id"`
	RoleId			int			`json:"role_id"`
}

type ListUserForm struct {
	Page			int			`json:"page" from:"page" valid:"Required;"`
	Size			int			`json:"size" from:"size" valid:"Required;"`

	//Username		string		`json:"username" from:"username"`
	//Password		string		`json:"password" from:"password"`
	//Email			string		`json:"email" from:"email"`
	//Phone			string		`json:"phone" from:"phone"`
	//Sex				int8		`json:"sex" from:"sex"`
	//Remark			string		`json:"remark" from:"remark"`
}

type ListUserRoleForm struct {
	UserId 			int			`json:"user_id"`
}






// @Summary Add user
// @Produce  json
// @Param user body AddUserForm true  "UserAdd"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/add [post]
func AddUser(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form AddUserForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	//获取user的service
	userService := user_service.User{
		Username: form.Username,
		Password: form.Password,
		Email:    form.Email,
		Phone:    form.Phone,
		Sex:      form.Sex,
		Remark:   form.Remark,
	}

	if err := userService.Add(); err!=nil{
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER,nil)
		return
	}
	appG.Response(http.StatusOK,e.SUCCESS, nil)
}

// @Summary Add user's role
// @Produce  json
// @Param userRole body AddUserRoleForm true  "UserRoleAdd"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/role/add [post]
func AddUserRole(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form AddUserRoleForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	//获取user的service
	ur := user_service.UserRole{
		UserId: form.UserId,
		RoleId: form.RoleId,
	}

	if err := ur.Add();err!= nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER,nil)
		return
	}
	appG.Response(http.StatusOK,e.SUCCESS, nil)
}

// @Summary List users
// @Produce  json
// @Param queryCondition body ListUserForm true  "UserQuery"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/list [post]
func ListUsers(c *gin.Context) {
	var(
		appG = app.Gin{c}
		form ListUserForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode!=e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	//获取user的service
	userService := user_service.User{
		//Username: form.Username,
		//Password: form.Password,
		//Email:    form.Email,
		//Phone:    form.Phone,
		//Sex:      form.Sex,
		//Remark:   form.Remark,

		Page:	  form.Page,
		Size: 	  form.Size,
	}

	users,total, err := userService.List()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER,nil)
		return
	}
	data := map[string]interface{}{
		"list": users,
		"total": total,
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// @Summary List user's roles
// @Produce  json
// @Param queryCondition body ListUserRoleForm true  "UserRolesQuery"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/role/list [post]
func ListRolesByUserId(c *gin.Context) {
	var(
		appG = app.Gin{c}
		form ListUserRoleForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode!=e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	//获取user的service
	ur := user_service.UserRole{
		UserId:  	form.UserId,
	}
	urs, err := ur.ListRoles()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER,nil)
		return
	}

	data:=map[string]interface{}{
		"list": urs,
		"size": len(urs),
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)
}


