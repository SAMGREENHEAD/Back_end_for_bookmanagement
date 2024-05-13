// contain all the routes to the database

package routes

import (
	"github.com/SAMGREENHEAD/Book_management/pkg/controllers"
	"github.com/gorilla/mux"
)

//will have routes to get to controllers
var RegisterBookStoreRoutes = func(router *mux.Router) {
router.HandleFunc("/books/",controllers.CreateBook).Methods("POST")
router.HandleFunc("/books/",controllers.GetBook).Methods("GET")
router.HandleFunc("/books/{bookId}",controllers.GetBookById).Methods("GET")
router.HandleFunc("/books/{bookId}",controllers.UpdateBook).Methods("PUT")
router.HandleFunc("/books/{bookId}",controllers.DeleteBook).Methods("DELETE")
	
}
