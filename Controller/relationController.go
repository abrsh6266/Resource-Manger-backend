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

func GetRelation(ctx *gin.Context) {
	jsonData := map[string]string{
		"query": `{
			materialCustodianRel {
			  id
			  userName
			  materialsSerialNumber
			  amount
			  Custodian {
				id
				name
				phoneNumber
				email
				department
			  }
			  Material {
				Id
				diskType
				model
				name
				networkType
				owner
				processor
				scanType
				serialNumber
				type
			  }
			}
		  }		  
		  `}
	marshaling, _ := json.Marshal(jsonData)
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(marshaling))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something wrong"})
		return
	}
	var result struct {
		Data struct {
			MaterialUserRelations []model.MaterialUserRelation `json:"materialCustodianRel"`
		} `json:"data"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse category data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": result.Data.MaterialUserRelations})
}
func AddRelation(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(body))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	} else if string(respBody)[:8] == "{\"errors" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}
	var result struct {
		Data struct {
			Relations []model.MaterialUserRelation `json:"insert_materialCustodianRel_one"`
		} `json:"data"`
		UpdatedMaterial struct {
			AffectedRows int `json:"affected_rows"`
		} `json:"update_Material"`
	}
	if json.Unmarshal(respBody, &result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse data"})
		return
	}
	c.JSON(http.StatusOK, result)
}
func UpdateRelation(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(body))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	} else if string(respBody)[:8] == "{\"errors" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}
	var result struct {
		Data struct {
			UpdatedMaterial struct {
				AffectedRows int `json:"affected_rows"`
			} `json:"update_Material"`
			UpdatedRelation struct {
				AffectedRows int `json:"affected_rows"`
			} `json:"update_materialCustodianRel"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody, &result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse 	data"})
		return
	}
	fmt.Println(result.Data.UpdatedMaterial.AffectedRows)
	c.JSON(http.StatusOK, result)
}
func RemoveRelation(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(body))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	} else if string(respBody)[:8] == "{\"errors" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}
	var result struct {
		Data struct {
			DeleteRelation struct {
				AffectedRows int `json:"affected_rows"`
			} `json:"delete_materialCustodianRel"`
			UpdatedMaterial struct {
				AffectedRows int `json:"affected_rows"`
			} `json:"update_Material"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody, &result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse data"})
		return
	}
	c.JSON(http.StatusOK, result)
}
