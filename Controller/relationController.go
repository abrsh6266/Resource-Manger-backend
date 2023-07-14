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

func GetRelation(ctx *gin.Context){
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
	marshaling,_ := json.Marshal(jsonData)
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
	if err := json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse category data"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message":result.Data.MaterialUserRelations})
}
func AddRelation(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	relation := model.MaterialUserRelation{}
	if err := json.Unmarshal(body, &relation); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	jsonData1 := map[string]string{
		"query": fmt.Sprintf(`{
			Material(where: {serialNumber: {_eq: "%s"}}) {
			  Id
			  serialNumber
			  taken
			  total
			  model
			  name
			}
		  }					
		  `,relation.SerialNumber),}
	marshaling1,_ := json.Marshal(jsonData1)
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(marshaling1))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something wrong"})
		return
	}
	var result struct {
		Data struct {
			Material []model.Material `json:"Material"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse category data"})
		return
	}
	fmt.Println(relation.SerialNumber)
	fmt.Println(relation.Amount)
	fmt.Println(result.Data.Material[0].Total)
	fmt.Println(result.Data.Material[0].Taken)
	if relation.Amount<=0 || relation.Amount > (result.Data.Material[0].Total - result.Data.Material[0].Taken){
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	
	jsonData3 := map[string]string{
		"query": fmt.Sprintf(`mutation {
			insert_materialCustodianRel_one(object: {userName: "%s", materialsSerialNumber: "%s", amount: "%d"}) {
			  userName
			  materialsSerialNumber
			  amount
			}
		  }`,relation.Name,relation.SerialNumber,relation.Amount),
	}
	marshaling,_ := json.Marshal(jsonData3)
	val, err := initializer.HasuraMutationRequest(marshaling)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else if string(val)[:8]=="{\"errors"{
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": relation})
	jsonData2 := map[string]string{
		"query": fmt.Sprintf(`mutation {
			update_Material(_set: {taken: "%d"}, where: {serialNumber: {_eq: "%s"}}) {
			  returning {
				Id
			  }
			}
		  }`,(result.Data.Material[0].Taken+relation.Amount),result.Data.Material[0].SerialNumber),
	}
	marshaling2,_ := json.Marshal(jsonData2)
	val,err = initializer.HasuraMutationRequest(marshaling2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}else if string(val)[0:8] == "{\"errors" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
}
func UpdateRelation(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	relation := model.MaterialUserRelation{}
	if json.Unmarshal(body, &relation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	jsonData1 := map[string]string{
		"query": fmt.Sprintf(`{
			Material(where: {serialNumber: {_eq: "%s"}}) {
			  Id
			  serialNumber
			  taken
			  total
			  model
			  name
			}
		  }					
		  `,relation.SerialNumber),}
	marshaling1,_ := json.Marshal(jsonData1)
	fmt.Println("check 1")
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(marshaling1))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something wrong"})
		return
	}else if (string(respBody)) == "{\"data\":{\"Material\":[]}}" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "invalid input"})
		return
	}
	var result struct {
		Data struct {
			Material []model.Material `json:"Material"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse category data"})
		return
	}
	fmt.Println("check 2")
	fmt.Println(relation.SerialNumber)
	fmt.Println(relation.Amount)
	fmt.Println(result.Data.Material[0].Total)
	fmt.Println(result.Data.Material[0].Taken)
	if relation.Amount+result.Data.Material[0].Taken>result.Data.Material[0].Total || relation.Amount+result.Data.Material[0].Taken<0{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	fmt.Println("check 4")
	jsonData := map[string]string{
		"query": fmt.Sprintf(`mutation {
			update_materialCustodianRel(where: {id: {_eq: "%d"}}, _set: {amount: "%d"}) {
			  returning {
				id
				userName
				materialsSerialNumber
				amount
			  }
			}
		  }`,relation.Id,relation.Amount),
	}
	marshaling2,_ := json.Marshal(jsonData)
	val2,err := initializer.HasuraMutationRequest(marshaling2)
	fmt.Println(string(val2))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}else if (string(val2)) == "{\"data\":{\"update_materialCustodianRel\":{\"returning\" : []}}}" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "invalid input"})
		return
	}
	fmt.Println("check 3")
	jsonData2 := map[string]string{
		"query": fmt.Sprintf(`mutation {
			update_Material(_set: {taken: "%d"}, where: {serialNumber: {_eq: "%s"}}) {
			  returning {
				Id
			  }
			}
		  }`,(result.Data.Material[0].Taken+relation.Amount),result.Data.Material[0].SerialNumber),
	}
	marshaling,_ := json.Marshal(jsonData2)
	val,err := initializer.HasuraMutationRequest(marshaling)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}else if string(val)[0:8] == "{\"errors" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully Updated"})
}
func RemoveRelation(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	relation := model.MaterialUserRelation{}
	if json.Unmarshal(body, &relation); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	jsonData := map[string]string{
		"query": fmt.Sprintf(`{
			materialCustodianRel_by_pk(id: %d) {
			  amount
			  materialsSerialNumber
			}
		  }	  
		  `,relation.Id),
		}
	marshaling,_ := json.Marshal(jsonData)
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(marshaling))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something wrong"})
		return
	}else if (string(respBody)) == "{\"data\":{\"materialCustodianRel_by_pk\": null}}" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "invalid input"})
		return
	}
	var resultamount struct {
		Data struct {
			MaterialUserRelations model.MaterialUserRelation `json:"materialCustodianRel_by_pk"`
		} `json:"data"`
	}
	if err := json.Unmarshal(respBody,&resultamount); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse category data"})
		return
	}
	jsonData5 := map[string]string{
		"query": fmt.Sprintf(`{
			Material(where: {serialNumber: {_eq: "%s"}}) {
			  taken
			}
		  }
		  `,resultamount.Data.MaterialUserRelations.SerialNumber),
		}
	marshaling5,_ := json.Marshal(jsonData5)
	respBody, err = initializer.HasuraRequest(http.MethodPost, string(marshaling5))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something wrong"})
		return
	}
	var result struct {
		Data struct {
			Materials []model.Material `json:"Material"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse category data"})
		return
	}
	fmt.Println(resultamount.Data.MaterialUserRelations.Amount)
	jsonData2 := map[string]string{
		"query": fmt.Sprintf(`mutation {
			update_Material(_set: {taken: "%d"}, where: {serialNumber: {_eq: "%s"}}) {
			  returning {
				Id
			  }
			}
		  }`,(result.Data.Materials[0].Taken-resultamount.Data.MaterialUserRelations.Amount),resultamount.Data.MaterialUserRelations.SerialNumber),
	}
	marshaling3,_ := json.Marshal(jsonData2)
	val,err := initializer.HasuraMutationRequest(marshaling3)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	jsonData3 := map[string]string{
		"query": fmt.Sprintf(`mutation {
			delete_materialCustodianRel_by_pk(id: "%d") {
			  amount
			  id
			  materialsSerialNumber
			  userName
			}
		  }`,relation.Id),
	}
	marshaling1,_ := json.Marshal(jsonData3)
	val, err = initializer.HasuraMutationRequest(marshaling1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}else if string(val) == "{\"data\":{\"delete_materialCustodianRel_by_pk\":null}}" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": string(val)})
}