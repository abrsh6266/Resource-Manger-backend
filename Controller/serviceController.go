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
func GetServices(ctx *gin.Context){
	body, err := ioutil.ReadAll(ctx.Request.Body)
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(body))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "some sort of error"})
		return
	}
	var result struct {
		Data struct {
			Materials  	 []model.Material `json:"Material"`
			Departments  []model.Department `json:"Department"`
			Relations    []model.MaterialUserRelation `json:"materialCustodianRel"`
			Custodians   []model.Custodian `json:"Custodian"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse category data"})
		return
		}
	ctx.JSON(http.StatusOK,result)
}
func AddCustodian(c *gin.Context) {
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
			Materials []model.Material `json:"insert_Custodian_one"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse Custodian data"})
		return
		}
	c.JSON(http.StatusOK,result)
}
func RemoveCustodian(c *gin.Context) {
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
			} `json:"delete_Custodian"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse data"})
		return
		}
	c.JSON(http.StatusOK,result)
}
func UpdateCustodian(c *gin.Context) {
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
			UpdatedMaterial struct{
				AffectedRows int  `json:"affected_rows"`
			} `json:"update_Custodian"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse Custodian data"})
		return
		}
		fmt.Println(result.Data.UpdatedMaterial.AffectedRows)
		c.JSON(http.StatusOK,result)
}