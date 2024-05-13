package main

import (
	"log"
	"net/http"

	"github.com/SAMGREENHEAD/Book_management/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main(){
// main initializes the router, registers book store routes, and starts the server.
//
// No parameters.
// No return values.
	r:= mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
