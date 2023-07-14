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
	jsonData := map[string]string{
		"query": `{
			Department {
			  id
			  location
			  name
			  Custodians {
				name
				id
				phoneNumber
				email
			  }
			}
		  }`}
	marshaling,_ := json.Marshal(jsonData)
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(marshaling))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Incorrect Email"})
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
	ctx.JSON(http.StatusOK,gin.H{"message":result.Data.Departments})
}
func AddDepatment(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	department := model.Department{}
	if json.Unmarshal(body, &department); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}else if department.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	jsonData := map[string]string{
		"query": fmt.Sprintf(`mutation {
			insert_Department_one(object: {location: "%s", name: "%s"}) {
			  id
			  location
			  name
			}
		  }`,department.Location,department.Name),
	}
	marshaling,_ := json.Marshal(jsonData)
	val, err := initializer.HasuraMutationRequest(marshaling)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else if (string(val))[0:5] == "{\"err" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "invalid input"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully added"})
}
func RemoveDepatment(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	department := model.Department{}
	if json.Unmarshal(body, &department); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	jsonData := map[string]string{
		"query": fmt.Sprintf(`mutation {
			delete_Department(where: {name: {_eq: "%s"}}) {
			  returning {
				id
				location
				name
			  }
			}
		  }`,department.Name),
	}
	marshaling,_ := json.Marshal(jsonData)
	val, err := initializer.HasuraMutationRequest(marshaling)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else if (string(val)) == "{\"data\":{\"delete_Department\":{\"returning\" : []}}}" {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": string(val)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": string(val)})
}
func UpdateDepatment(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	department := model.Department{}
	if json.Unmarshal(body, &department); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	jsonData := map[string]string{
		"query": fmt.Sprintf(`mutation {
			update_Department(where: {name: {_eq: "%s"}}, _set: {%s: "%s"}) {
			  returning {
				id
				location
				name
			  }
			}
		  }`,department.Name,department.UpdateColumn,department.UpdateValue),
	}	
	marshaling2,_ := json.Marshal(jsonData)
	val,err := initializer.HasuraMutationRequest(marshaling2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}else if string(val)=="{\"data\":{\"update_Department\":{\"returning\" : []}}}"{
		c.JSON(http.StatusBadRequest,gin.H{"error":"department not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated"})
}