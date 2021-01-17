package v1

import (
	"bingcu.top/go-gin-users-permission-control/pkg/app"
	"bingcu.top/go-gin-users-permission-control/pkg/e"
	"bingcu.top/go-gin-users-permission-control/service/role_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddRoleForm struct {
	Name 		string				`json:"name" form:"name" valid:"Required;"`
	Description string				`json:"description" form:"description" valid:"Required;"`
	Remark 		string				`json:"remark" form:"remark" valid:"Required;"`
}

type AddRolePermissionForm struct {
	RoleId				int		`json:"role_id"`
	PermissionId		int		`json:"permission_id"`
}

type ListRoleForm struct {
	Page 	int `valid:"Required;"`
	Size 	int `valid:"Required;"`
}

type ListRolePermissionsForm struct {
	RoleId			int			`json:"role_id"`
}

type DeletePermissionForm struct {
	RolePermissionId	int		`json:"role_permission_id"`
}



// @Summary Add role
// @Produce  json
// @Param role body AddRoleForm true  "RoleAdd"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/role/add [post]
func AddRole(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form AddRoleForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	//获取role的service
	roleService := role_service.Role{
		Name: 		form.Name,
		Description:form.Description,
		Remark:   	form.Remark,
	}

	if err := roleService.Add(); err!=nil{
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER,nil)
		return
	}
}

// @Summary Add rolePermission
// @Produce  json
// @Param rolePermission body AddRolePermissionForm true  "RolePermissionAdd"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/role/permission/add [post]
func AddRolePermission(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form AddRolePermissionForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	//获取role的service
	rp := role_service.RolePermission{
		RoleId:  		form.RoleId,
		PermissionId:  	form.PermissionId,
	}

	if err := rp.Add();err!= nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER,nil)
		return
	}
	appG.Response(http.StatusOK,e.SUCCESS, nil)
}

// @Summary Delete rolePermission
// @Produce  json
// @Param rolePermission body DeletePermissionForm true  "RolePermissionDelete"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/role/permission/delete [post]
func DeletePermission(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form DeletePermissionForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	//获取role的service
	roleService := role_service.RolePermission{
		ID: 	form.RolePermissionId,
	}

	if err := roleService.Delete(); err!=nil{
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER,nil)
		return
	}
}


// @Summary List roles
// @Produce  json
// @Param queryCondition body ListRoleForm true  "RolesQuery"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/role/list [post]
func ListRoles(c *gin.Context) {
	var(
		appG = app.Gin{c}
		form ListRoleForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode!=e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	//获取role的service
	roleService := role_service.Role{
		Page:	  form.Page,
		Size: 	  form.Size,
	}

	roles,total, err := roleService.List()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER,nil)
		return
	}
	data := map[string]interface{}{
		"list": roles,
		"total": total,
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// @Summary List rolePermission
// @Produce  json
// @Param rolePermission body ListRolePermissionsForm true  "RolePermissionList"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/role/permission/list [post]
func ListPermissionsByRoleId(c *gin.Context) {
	var(
		appG = app.Gin{c}
		form ListRolePermissionsForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode!=e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	//获取role的service
	rp := role_service.RolePermission{
		RoleId:  		form.RoleId,
	}

	rps, err := rp.ListPermissions()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER,nil)
		return
	}

	data:=map[string]interface{}{
		"list": rps,
		"size": len(rps),
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)
}



