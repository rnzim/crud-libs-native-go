package main

import (
	"net/http"
	"store-app/routes"
)
 



func main(){
	routes.LoadRoutes()
	http.ListenAndServe(":8000",nil)
	
}
