package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gilab.com/pragmaticreviews/golang-gin-poc/initializer"
	"gilab.com/pragmaticreviews/golang-gin-poc/utils"
	"github.com/gin-gonic/gin"
)

// JWTMiddleware is a middleware that validates JWT tokens
func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
		fmt.Println(tokenString)
        if tokenString == "" {
			fmt.Println("not ok 2")
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            c.Abort()
            return
        }

        email, err := utils.ParseToken(tokenString)
		fmt.Println(email)
        if err != nil {
			fmt.Println("not ok 3")
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
		jsonData := map[string]string{
			"query": fmt.Sprintf(`{
				Manager_by_pk(email: "%s") {
				  email
				  name
				  password
				  role
				}
			  }`,email),
		}
		marshaling,_ := json.Marshal(jsonData)
		res, err := initializer.HasuraRequest(http.MethodPost, string(marshaling))
		if err != nil {
			fmt.Println("not ok 4")
			c.AbortWithStatus(http.StatusUnauthorized)
			c.Abort()
			return
		}else if string(res)=="{\"data\":{\"Manager_by_pk\":null}}"{
			fmt.Println("not ok 5")
			c.AbortWithStatus(http.StatusUnauthorized)
			c.Abort()
			return 
		}
		fmt.Println("middleware passed")
        c.Next()
    }
}