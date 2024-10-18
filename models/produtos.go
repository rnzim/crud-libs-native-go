package models

import (

	"store-app/db"
)
type Product struct{
	Id int
	Name string
	Descricao string
	Preco float64

}

func FindAll() []Product  {
	query:="select * from produtos"
	db := db.ConnDb()
	list,err := db.Query(query)

	if err != nil{
		panic(err.Error())
	}
	produtos := []Product{}
	newProduct := Product{}
	for list.Next(){
		var id int
		var nome,descricao string
		var preco float64

		if err != nil{
			panic(err.Error())
		}
		err = list.Scan(&id,&nome,&descricao,&preco)
		newProduct.Name = nome
		newProduct.Descricao = descricao
		newProduct.Preco = preco
		newProduct.Id = id
		produtos = append(produtos,newProduct)
	}
    defer db.Close()
	return produtos
}
func FindById(id string) Product{
	db := db.ConnDb()
	item,err:= db.Query("Select * From produtos Where id=$1;",id)
	if err != nil {
		panic(err.Error())
	}
	
	Product := Product{}
	for item.Next(){
		
		var id int
		var nome,descricao string
		var preco float64

		item.Scan(&id,&nome,&descricao,&preco)
		Product.Name = nome
		Product.Descricao = descricao
		Product.Preco = preco
		Product.Id = id
		
	}
	defer db.Close()
	return Product
}
func Create(nome,descricao string,preco float64)  {
	db := db.ConnDb()

	insert,err:= db.Prepare("insert into produtos (nome,descricao,preco) values ($1,$2,$3)")

	if err != nil{
		panic(err.Error())
	}
	insert.Exec(nome,descricao,preco)
	defer db.Close()
}
func Delete(id string)  {
	db := db.ConnDb()
	result,err:=db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	result.Exec(id)
	defer db.Close()
}
func Update(id int,name,descricao string,preco float64){

	db := db.ConnDb()
	update,err := db.Prepare("update produtos set nome=$1,descricao=$2,preco=$3 where id=$4")
	if err!=nil {
		panic(err.Error())
	}
	update.Exec(name,descricao,preco,id)
	
	defer db.Close()
}
