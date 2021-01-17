package v1

import (
	"bingcu.top/go-gin-users-permission-control/pkg/app"
	"bingcu.top/go-gin-users-permission-control/pkg/e"
	"bingcu.top/go-gin-users-permission-control/service/permission_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddPermissionForm struct {
	PermissionKey		string		`json:"permission_key" from:"permission_key" valid:"Required;"`
	Description 		string		`json:"description" form:"description" valid:"Required;"`
	Remark 				string		`json:"remark" form:"remark" valid:"Required;"`
}

type ListPermissionForm struct {
	Page			int			`json:"page" from:"page" valid:"Required;"`
	Size			int			`json:"size" from:"size" valid:"Required;"`
}






// @Summary Add permission
// @Produce  json
// @Param permission body AddPermissionForm true  "PermissionAdd"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/permission/add [post]
func AddPermission(c *gin.Context)  {
	var (
		appG = app.Gin{c}
		form AddPermissionForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	//获取user的service
	permissionService := permission_service.Permission{
		PermissionKey: form.PermissionKey,
		Description:   form.Description,
		Remark:   	   form.Remark,
	}

	if err := permissionService.Add(); err!=nil{
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER,nil)
		return
	}
}




// @Summary List permissions
// @Produce  json
// @Param queryCondition body ListPermissionForm true  "PermissionQuery"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/permission/list [post]
func ListPermissions(c *gin.Context) {
	var(
		appG = app.Gin{c}
		form ListUserForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode!=e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	//获取permission的service
	permissionService := permission_service.Permission{
		Page:	  form.Page,
		Size: 	  form.Size,
	}

	ps,total, err := permissionService.List()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER,nil)
		return
	}
	data := map[string]interface{}{
		"list": ps,
		"total": total,
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)
}


