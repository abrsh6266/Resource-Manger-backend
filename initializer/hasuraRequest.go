package initializer

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)
var client = &http.Client{}
func HasuraRequest(method, query string) ([]byte, error) {
	req, err := http.NewRequest(method, os.Getenv("HASURAURL"), strings.NewReader(query))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Hasura-Admin-Secret", os.Getenv("HASURAADMINSECRET"))
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	return body, nil
}
func HasuraMutationRequest(query []byte) ([]byte, error) {
	return HasuraRequest(http.MethodPost,string(query))
}