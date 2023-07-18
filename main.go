package main

import (
	controller "gilab.com/pragmaticreviews/golang-gin-poc/Controller"
	"gilab.com/pragmaticreviews/golang-gin-poc/initializer"
	"gilab.com/pragmaticreviews/golang-gin-poc/middleware"
	"github.com/gin-gonic/gin"
)
func init(){
	initializer.LoadEnvVariables()
}
func main(){
	r := gin.Default()
	r.Static("/css","tempates/css")
	r.LoadHTMLGlob("templates/*.html")
	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/users",controller.Login)
		apiRoutes.POST("/users",controller.Signup)
		apiRoutes.GET("/department/get",middleware.RequireAuth,controller.GetDepartments)
		apiRoutes.POST("/department/add",middleware.RequireAuth,controller.AddDepatment)
		apiRoutes.POST("/department/remove",middleware.RequireAuth,controller.RemoveDepatment)
		apiRoutes.POST("/department/update",middleware.RequireAuth,controller.UpdateDepatment)
		apiRoutes.POST("/custodian/add",middleware.RequireAuth,controller.AddCustodian)
		apiRoutes.POST("/custodian/remove",middleware.RequireAuth,controller.RemoveCustodian)
		apiRoutes.POST("/custodian/update",middleware.RequireAuth,controller.UpdateCustodian)
		apiRoutes.GET("/category/get",middleware.RequireAuth,controller.GetCategories)
		apiRoutes.POST("/category/add",middleware.RequireAuth,controller.AddCategory)
		apiRoutes.POST("/category/update",middleware.RequireAuth,controller.UpdateCategory)
		apiRoutes.GET("/material/get",middleware.RequireAuth,controller.GetMaterials)
		apiRoutes.POST("/material/add",middleware.RequireAuth,controller.AddMaterial)
		apiRoutes.POST("/material/update",middleware.RequireAuth,controller.UpdateMaterial)
		apiRoutes.POST("/material/remove",middleware.RequireAuth,controller.RemoveMaterial)
		apiRoutes.GET("/usermat/get",middleware.RequireAuth,controller.GetRelation)
		apiRoutes.POST("/usermat/add",middleware.RequireAuth,controller.AddRelation)
		apiRoutes.POST("/usermat/update",middleware.RequireAuth,controller.UpdateRelation)
		apiRoutes.POST("/usermat/remove",middleware.RequireAuth,controller.RemoveRelation)
	}
	viewRoutes := r.Group("/view")
	{
		viewRoutes.GET("/userlogin",controller.ShowAll)
	}
	r.Run()
}