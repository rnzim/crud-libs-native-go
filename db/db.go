package db

import (
	 "database/sql"
	_"github.com/lib/pq"
)
func ConnDb() *sql.DB{
	connDb := "user=postgres dbname=produtos password=1234 host=localhost port=5432 sslmode=disable"
	db,err := sql.Open("postgres",connDb)
	if err != nil{
	   panic("Erro")
	}
	return db
   }