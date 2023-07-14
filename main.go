package main

import (
	"os"

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
	r.GET("/users",controller.Login)
	r.POST("/users",controller.Signup)
	r.GET("/department/get",middleware.RequireAuth,controller.GetDepartments)
	r.POST("/department/add",middleware.RequireAuth,controller.AddDepatment)
	r.POST("/department/remove",middleware.RequireAuth,controller.RemoveDepatment)
	r.POST("/department/update",middleware.RequireAuth,controller.UpdateDepatment)
	r.POST("/custodian/add",middleware.RequireAuth,controller.AddCustodian)
	r.POST("/custodian/remove",middleware.RequireAuth,controller.RemoveCustodian)
	r.POST("/custodian/update",middleware.RequireAuth,controller.UpdateCustodian)
	r.GET("/category/get",middleware.RequireAuth,controller.GetCategories)
	r.POST("/category/add",middleware.RequireAuth,controller.AddCategory)
	r.POST("/category/update",middleware.RequireAuth,controller.UpdateCategory)
	r.GET("/material/get",middleware.RequireAuth,controller.GetMaterials)
	r.POST("/material/add",middleware.RequireAuth,controller.AddMaterial)
	r.POST("/material/update",middleware.RequireAuth,controller.UpdateMaterial)
	r.POST("/material/remove",middleware.RequireAuth,controller.RemoveMaterial)
	r.GET("/usermat/get",middleware.RequireAuth,controller.GetRelation)
	r.POST("/usermat/add",middleware.RequireAuth,controller.AddRelation)
	r.POST("/usermat/update",middleware.RequireAuth,controller.UpdateRelation)
	r.POST("/usermat/remove",middleware.RequireAuth,controller.RemoveRelation)
	r.Run(os.Getenv("PORT"))
}