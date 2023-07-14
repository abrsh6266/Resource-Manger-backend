package middleware

import (
	"encoding/json"
	"fmt"
	"go/token"
	"net/http"
	"os"
	"time"

	"gilab.com/pragmaticreviews/golang-gin-poc/initializer"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(ctx *gin.Context) {
	fmt.Println("in middleware")
	tokenString,err := ctx.Cookie("Authorization")
	if err != nil{
		fmt.Println("notok1")
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	token , err := jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {
		if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
			fmt.Println("notok2")
			return nil,fmt.Errorf("unexpected signing method : %v",token.HighestPrec)
		}
		return []byte(os.Getenv("SECRET")),nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok&&token.Valid{
		if float64(time.Now().Unix())>claims["exp"].(float64){
			fmt.Println("notok3")
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		jsonData := map[string]string{
			"query": fmt.Sprintf(`{
				Manager_by_pk(email: "%s") {
				  email
				  name
				  password
				}
			  }`,claims["sub"]),
		}
		marshaling,_ := json.Marshal(jsonData)
		_, err := initializer.HasuraRequest(http.MethodPost, string(marshaling))
		if err != nil {
			fmt.Println("notok4")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}else{
		fmt.Println("ok5")
		ctx.AbortWithStatus(http.StatusUnauthorized)
	} 
}