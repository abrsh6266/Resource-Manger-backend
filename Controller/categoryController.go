package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"gilab.com/pragmaticreviews/golang-gin-poc/initializer"
	"gilab.com/pragmaticreviews/golang-gin-poc/model"
	"github.com/gin-gonic/gin"
)
func GetCategory(ctx *gin.Context){
	body, err := io.ReadAll(ctx.Request.Body)
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(body))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "some sort of error"})
		return
	}
	var result struct {
		Data struct {
			Category []model.Category `json:"Category"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	ctx.JSON(http.StatusOK,result)
}
func GetCategories(ctx *gin.Context){
	jsonData := map[string]string{
		"query": `{
			Category {
			  id
			  name
			  total
			  Materials {
				Id
				name
				model
				networkType
				owner
				processor
				scanType
				serialNumber
				ram
				storage
				taken
				total
				type
				diskType
			  }
			}
		  }
		  `}
	marshaling,_ := json.Marshal(jsonData)
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(marshaling))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something wrong"})
		return
	}
	var result struct {
		Data struct {
			Categories []model.Category `json:"Category"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse category data"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message":result.Data.Categories})
}
func AddCategory(c *gin.Context) {
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
			Catigory model.Category `json:"insert_Category_one"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse category data"})
		return
		}
	c.JSON(http.StatusOK,result)
}
func UpdateCategory(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	category := model.Category{}
	if json.Unmarshal(body, &category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	jsonData := map[string]string{
			"query": fmt.Sprintf(`mutation {
				update_Category(where: {name: {_eq: "%s"}}, _set: {name: "%s"}) {
				  returning {
					id
					name
					total
				  }
				}
			  }`,category.Name,category.UpdateValue),
		}
	marshaling2,_ := json.Marshal(jsonData)
	val,err := initializer.HasuraMutationRequest(marshaling2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}else if string(val) == "{\"data\":{\"update_Category\":{\"returning\" : []}}}" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully Updated"})
}