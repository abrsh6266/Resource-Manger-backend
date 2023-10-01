package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gilab.com/pragmaticreviews/golang-gin-poc/initializer"
	"gilab.com/pragmaticreviews/golang-gin-poc/model"
	"github.com/gin-gonic/gin"
)

func GetDepartments(ctx *gin.Context){
	body, err := ioutil.ReadAll(ctx.Request.Body)
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(body))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "some sort of error"})
		return
	}
	var result struct {
		Data struct {
			Departments []model.Department `json:"Department"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	ctx.JSON(http.StatusOK,result)
}
func AddDepatment(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(body))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}else if string(respBody)[:8]=="{\"errors"{
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}
	var result struct {
		Data struct {
			Materials []model.Material `json:"insert_Material_one"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse category data"})
		return
		}
	c.JSON(http.StatusOK,result)
}
func RemoveDepartment(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(body))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}else if string(respBody)[:8]=="{\"errors"{
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}
	var result struct {
		Data struct {
			DeletedMaterial struct{
				AffectedRows int  `json:"affected_rows"`
			} `json:"delete_Department"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse data"})
		return
		}
	c.JSON(http.StatusOK,result)
}
func UpdateDepatment(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(body))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}else if string(respBody)[:8]=="{\"errors"{
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}
	var result struct {
		Data struct {
			UpdatedDepartment struct{
				AffectedRows int  `json:"affected_rows"`
			} `json:"update_Department"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse category data"})
		return
		}
		c.JSON(http.StatusOK,result)
}