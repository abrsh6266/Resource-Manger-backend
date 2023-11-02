package routes

import (
	controller "gilab.com/pragmaticreviews/golang-gin-poc/Controller"
	"gilab.com/pragmaticreviews/golang-gin-poc/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
	
// SetupRouter configures the application's routes
func SetupRouter() *gin.Engine {
    r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Set AllowAllOrigins to true
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
    
	// Public routes
    r.POST("/register", controller.Signup)
    r.POST("/login", controller.Login)

    // Protected routes
    protected := r.Group("/api")
    protected.Use(middleware.JWTMiddleware())
    {
		protected.POST("/managers",controller.GetMangers)
		protected.POST("/report",controller.GetCategory)
		protected.POST("/department/get",controller.GetDepartments)
		protected.POST("/department/add",controller.AddDepatment)
		protected.POST("/department/remove",controller.RemoveDepartment)
		protected.POST("/department/update",controller.UpdateDepatment)
		protected.POST("/custodian/add",controller.AddCustodian)
		protected.POST("/custodian/remove",controller.RemoveCustodian)
		protected.POST("/custodian/update",controller.UpdateCustodian)
		protected.GET("/category/get",controller.GetCategories)
		protected.POST("/category/add",controller.AddCategory)
		protected.POST("/category/update",controller.UpdateCategory)
		protected.POST("/getservices",controller.GetServices)
		protected.POST("/material/get",controller.GetMaterials)
		protected.POST("/material/add",controller.AddMaterial)
		protected.POST("/material/update",controller.UpdateMaterial)
		protected.POST("/material/remove",controller.RemoveMaterial)
		protected.GET("/usermat/get",controller.GetRelation)
		protected.POST("/usermat/add",controller.AddRelation)
		protected.POST("/usermat/update",controller.UpdateRelation)
		protected.POST("/usermat/remove",controller.RemoveRelation)
    }
    return r
}