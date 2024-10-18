package controllers

import (
	"log"
	"net/http"
	_ "os"
	"store-app/models"
	"strconv"
	"text/template"
)
var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter,r *http.Request){
	produtos := models.FindAll()
	
    temp.ExecuteTemplate(w,"Index",produtos)
}
func NewProduct(w http.ResponseWriter,r *http.Request)  {

	temp.ExecuteTemplate(w,"New",nil)
}
func Insert(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		nome := r.FormValue("nome")
		descricao:= r.FormValue("descricao")
		preco,_:= strconv.ParseFloat(r.FormValue("preco"),64)
		models.Create(nome,descricao,preco)
		http.Redirect(w, r, "/", 301)
        return
		
	}
	http.Redirect(w,r,"/",http.StatusSeeOther)

}
func Delete(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET"{
		id:= r.URL.Query().Get("id")
		models.Delete(id)
		http.Redirect(w,r,"/",http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter,r *http.Request){
	id:= r.URL.Query().Get("id")
	produto:=models.FindById(id)
	temp.ExecuteTemplate(w,"Edit",produto)
}

func UpdateProduto(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		id,err:= strconv.Atoi(r.FormValue("id"))
		if err != nil{
			log.Println("Não Foi Possivel Converter")
		}

		name:= r.FormValue("nome")
		descricao:= r.FormValue("descricao")
		preco,_:= strconv.ParseFloat(r.FormValue("preco"),64)
		log.Println(id,name,descricao,preco)

		models.Update(id,name,descricao,preco)
		http.Redirect(w, r, "/", 301)
        return
		
	}
	
}