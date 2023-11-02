package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"gilab.com/pragmaticreviews/golang-gin-poc/initializer"
	"gilab.com/pragmaticreviews/golang-gin-poc/model"
	"gilab.com/pragmaticreviews/golang-gin-poc/utils"
	"github.com/gin-gonic/gin"
)
func GetMangers(ctx *gin.Context){
	body, err := io.ReadAll(ctx.Request.Body)
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(body))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "some sort of error"})
		return
	}
	var result struct {
		Data struct {
			Manager []model.User `json:"Manager"`
			Delete_Manager_by_pk model.User `json:"delete_Manager_by_pk"`
			Update_Manager_by_pk model.User `json:"update_Manager_by_pk"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse Manage's data"})
		return
		}
	ctx.JSON(http.StatusOK,result)
}
func Login(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	respBody, err2 := initializer.HasuraRequest(http.MethodPost, string(body))
	if err2 != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server Error"})
		return
	} else if string(respBody) == "{\"data\":{\"Manager\": []}}" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Email"})
		return
	}
	var result struct {
		Data struct {
			Managers []model.User `json:"Manager"`
		} `json:"data"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		fmt.Println("failed to Parse the data")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	if result.Data.Managers[0].Email == "" {
		fmt.Println("Error!!")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error"})
		return
	}
	token, err := utils.GenerateToken(result.Data.Managers[0].Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	result.Data.Managers[0].Token = token
	c.JSON(http.StatusOK, result)
}
func Signup(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	val, err := initializer.HasuraMutationRequest(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else if (string(val))[0:5] == "{\"err" {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": string(val)})
		return
	}
	var result struct {
		Data struct {
			Managers model.User `json:"insert_Manager_one"`
		} `json:"data"`
	}
	if err := json.Unmarshal(val, &result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	c.JSON(http.StatusOK, result)
}
func Validate(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "I am logged in",
	})
}
