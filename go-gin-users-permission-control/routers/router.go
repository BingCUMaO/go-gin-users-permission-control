package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "bingcu.top/go-gin-users-permission-control/docs"
	"bingcu.top/go-gin-users-permission-control/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//设置跨域
	r.Use(cors.Default())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")

	//user表的CURD
	//增加用户
	apiv1.POST("/user/add", v1.AddUser)
	//查询用户列表
	apiv1.POST("/user/list", v1.ListUsers)

	//role表的CURD
	//增加角色
	apiv1.POST("/role/add", v1.AddRole)
	//查询角色列表
	apiv1.POST("/role/list", v1.ListRoles)

	//permission表的CURD
	//增加权限类型
	apiv1.POST("/permission/add", v1.AddPermission)
	//查询权限类型列表
	apiv1.POST("/permission/list", v1.ListPermissions)

	//user_role表的CURD
	//添加指定用户的角色
	apiv1.POST("/user/role/add", v1.AddUserRole)
	//查询指定用户的所有角色
	apiv1.POST("/user/role/list", v1.ListRolesByUserId)

	//role_permission表的CURD
	//添加指定角色的权限
	apiv1.POST("/role/permission/add", v1.AddRolePermission)
	//删除指定角色对应的权限
	apiv1.POST("/role/permission/delete", v1.DeletePermission)
	//查询指定角色的所有权限
	apiv1.POST("/role/permission/list", v1.ListPermissionsByRoleId)



	return r
}
