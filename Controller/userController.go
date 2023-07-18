package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"gilab.com/pragmaticreviews/golang-gin-poc/initializer"
	"gilab.com/pragmaticreviews/golang-gin-poc/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	//"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) (){
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	manager2 := model.User{}
	if json.Unmarshal(body,&manager2); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	jsonData := map[string]string{
		"query": fmt.Sprintf(`{
			Manager_by_pk(email: "%s") {
			  email
			  name
			  password
			}
		  }`,manager2.Email),
	}
	marshaling,_ := json.Marshal(jsonData)
	respBody, err2 := initializer.HasuraRequest(http.MethodPost, string(marshaling))
	if err2 != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server Error"})
		return
	}else if string(respBody)=="{\"data\":{\"Manager_by_pk\":null}}"{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Incorrect Email"})
		return
	}
	var result struct {
		Data struct {
			Users model.User `json:"Manager_by_pk"`
		} `json:"data"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	fmt.Println(manager2.Password)
	fmt.Println(result.Data.Users.Password)
	errr := bcrypt.CompareHashAndPassword([]byte(result.Data.Users.Password),[]byte(manager2.Password))
	if errr != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"subs": manager2.Email,
		"exp" : time.Now().Add(time.Hour*24*30).Unix(),
	})
	tokenString,err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to create token"})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization",tokenString,3600*24*30,"","",false,true)
	c.JSON(http.StatusOK, gin.H{
		"message":result.Data.Users,
		"token":tokenString,
	})
}
func Signup(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	manager := model.User{}
	if json.Unmarshal(body,&manager); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	hashed,errr := bcrypt.GenerateFromPassword([]byte(manager.Password),10)
     if errr != nil {
		fmt.Println(err)
		return
	}
	jsonData := map[string]string{
		"query": fmt.Sprintf(`mutation {
			insert_Manager_one(object: {password: "%s", name: "%s", email: "%s"}) {
			  email
			  name
			  password
			}
		}`, hashed, manager.Name, manager.Email),
	}
	marshaling,_ := json.Marshal(jsonData)
	val, err := initializer.HasuraMutationRequest(marshaling)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": err})
		return
	}else if (string(val))[0:5]=="{\"err"{
		c.JSON(http.StatusInternalServerError, gin.H{"Error": string(val)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": manager})
}
func Validate(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"message":"I am logged in",
	})
}
func ShowAll(ctx *gin.Context){
	ctx.HTML(http.StatusOK,"index.html",gin.H{
		"title": "Login",
		"form": "xoxo",
	})
}