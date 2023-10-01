package main

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/initializer"
	"gilab.com/pragmaticreviews/golang-gin-poc/routes"
)
func init(){
	initializer.LoadEnvVariables()
}
func main() {
    r := routes.SetupRouter()
	r.Run(":8082")
}