package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"gilab.com/pragmaticreviews/golang-gin-poc/initializer"
	"gilab.com/pragmaticreviews/golang-gin-poc/model"
	"github.com/gin-gonic/gin"
)

func GetMaterials(ctx *gin.Context){
	jsonData := map[string]string{
		"query": `{
			Material {
			  diskType
			  model
			  name
			  networkType
			  owner
			  processor
			  scanType
			  serialNumber
			  taken
			  total
			  type
			  Id
			  MaterialCustodianRels {
				id
				userName
				materialsSerialNumber
				amount
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
			Materials []model.Material `json:"Material"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse category data"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message":result.Data.Materials})
}
func AddMaterial(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	material := model.Material{}
	if json.Unmarshal(body, &material); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}else if material.SerialNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	var attributes []string
	val := reflect.ValueOf(material)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := val.Type().Field(i).Tag.Get("json")
		value := field.Interface()

		// Skip fields that are zero values or empty strings
		if reflect.DeepEqual(value, reflect.Zero(field.Type()).Interface()) ||
			reflect.TypeOf(value).Kind() == reflect.String && value == "" {
			continue
		}
		// Add the attribute to the list
		attributes = append(attributes, tag+": \""+fmt.Sprintf("%v", value)+"\"")
	}

	object := strings.Join(attributes, ",")
	jsonData := map[string]string{
		"query": fmt.Sprintf(`mutation {
			insert_Material_one(object: {%s}) {
			  diskType
			  model
			  name
			  networkType
			  owner
			  processor
			  scanType
			  serialNumber
			  taken
			  total
			  type
			  Id
			}
		  }`,object),
	}
	fmt.Println(object)
	marshaling,_ := json.Marshal(jsonData)
	val2, err := initializer.HasuraMutationRequest(marshaling)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}else if string(val2)[:8]=="{\"errors"{
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": string(val2)})
}
func UpdateMaterial(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	material := model.Material{}
	if json.Unmarshal(body, &material); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	jsonData := map[string]string{
		"query": fmt.Sprintf(`mutation {
			update_Material(_set: {%s: "%s"}, where: {serialNumber: {_eq: "%s"}}) {
			  returning {
				Id
				diskType
				model
				name
				networkType
				owner
				processor
				scanType
				serialNumber
				taken
				total
				type
			  }
			}
		  }`,material.UpdateColumn,material.UpdateValue,material.SerialNumber),
	}
	marshaling2,_ := json.Marshal(jsonData)
	val,err := initializer.HasuraMutationRequest(marshaling2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}else if (string(val)) == "{\"data\":{\"update_Material\":{\"returning\" : []}}}" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "invalid input"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully Updated"})
}
func RemoveMaterial(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	material := model.Material{}
	if json.Unmarshal(body, &material); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	jsonData := map[string]string{
		"query": fmt.Sprintf(`mutation {
			delete_Material_by_pk(Id: %d) {
			  Id
			  name
			  serialNumber
			}
		  }`,material.Id),
	}
	marshaling,_ := json.Marshal(jsonData)
	val, err := initializer.HasuraMutationRequest(marshaling)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}else if string(val) == "{\"data\":{\"delete_Material_by_pk\":null}}" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": string(val)})
}