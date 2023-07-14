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

func AddCustodian(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	custodian := model.Custodian{}
	if json.Unmarshal(body, &custodian); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}else if custodian.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	jsonData := map[string]string{
		"query": fmt.Sprintf(`mutation {
			insert_Custodian_one(object: {department: "%s", email: "%s", name: "%s", phoneNumber: "%s"}) {
			  department
			  email
			  id
			  name
			  phoneNumber
			}
		  }`,custodian.Department,custodian.Email,custodian.Name,custodian.PhoneNumber),
	}
	marshaling,_ := json.Marshal(jsonData)
	val, err := initializer.HasuraMutationRequest(marshaling)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}else if string(val)[0:8]=="{\"errors"{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	fmt.Println(string(val))
	c.JSON(http.StatusOK, gin.H{"message": "Successfully added"})
}
func RemoveCustodian(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	custodian := model.Custodian{}
	if json.Unmarshal(body, &custodian); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	jsonData := map[string]string{
		"query": fmt.Sprintf(`mutation {
			delete_Custodian(where: {name: {_eq: "%s"}}) {
			  returning {
				department
				name
				email
				phoneNumber
			  }
			}
		  }`,custodian.Name),
	}
	marshaling,_ := json.Marshal(jsonData)
	val, err := initializer.HasuraMutationRequest(marshaling)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}else if string(val) == "{\"data\":{\"delete_Custodian\":{\"returning\" : []}}}" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully removed"})
}
func UpdateCustodian(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	custodian := model.Custodian{}
	if json.Unmarshal(body, &custodian); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	jsonData := map[string]string{
			"query": fmt.Sprintf(`mutation {
				update_Custodian(where: {name: {_eq: "%s"}}, _set: {%s: "%s"}) {
				  returning {
					department
					email
					name
					phoneNumber
				  }
				}
			  }`,custodian.Name,custodian.UpdateColumn,custodian.UpdateValue),
		}
	marshaling2,_ := json.Marshal(jsonData)
	val,err := initializer.HasuraMutationRequest(marshaling2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}else if string(val) == "{\"data\":{\"update_Custodian\":{\"returning\" : []}}}" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully Updated"})
}