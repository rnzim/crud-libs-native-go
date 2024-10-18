package routes

import(
	"net/http"
	"store-app/controllers"
)


func LoadRoutes(){
	//GET
	http.HandleFunc("/",controllers.Index)
	//GET
	http.HandleFunc("/new",controllers.NewProduct)
	//POST
	http.HandleFunc("/insert",controllers.Insert)
	//GET
	http.HandleFunc("/delete",controllers.Delete)
	//GET
	http.HandleFunc("/edit",controllers.Edit)
	//POST
	http.HandleFunc("/update",controllers.UpdateProduto)
}