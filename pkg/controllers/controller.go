package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SAMGREENHEAD/Book_management/pkg/models"
	"github.com/SAMGREENHEAD/Book_management/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

// Retrieves all books from the database and returns them as JSON
//
// Parameters: None
//
// Returns: A JSON array of Book structs
func GetBook(w http.ResponseWriter, r *http.Request) {
	// Retrieves all books from the database
	books := models.GetAllBooks()

	// Marshal the books into a JSON array
	res, _:= json.Marshal(books)

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the response to the client
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Retrieves a single book from the database and returns it as JSON
//
// Parameters:
//  bookId: The ID of the book to retrieve
//
// Returns:
//  A JSON representation of the Book struct with the given ID
func GetBookById(w http.ResponseWriter, r *http.Request) {
	// Get the book ID from the URL parameters
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	// Parse the book ID into an integer
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		// If there is an error, print it to the console
		fmt.Println("error while parsing")
	}

	// Retrieve the book from the database
	bookDetails, _ := models.GetBookById(ID)

	// Marshal the book into a JSON array
	res, _ := json.Marshal(bookDetails)

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the response to the client
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Creates a new book in the database and returns it as JSON
//
// Parameters: None
//
// Returns: A JSON representation of the newly created Book struct
func CreateBook(w http.ResponseWriter, r *http.Request) {
	// Create a new Book struct and parse the request body into it
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)

	// Create the book in the database and store it in the variable b
	b := CreateBook.CreateBook()

	// Marshal the Book struct into a JSON array and write it to the response
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Deletes a book from the database and returns it as JSON
//
// Parameters:
//  bookId: The ID of the book to delete
//
// Returns:
//  A JSON representation of the deleted Book struct
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Get the book ID from the URL parameters
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	// Parse the book ID into an integer
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		// If there is an error, print it to the console
		fmt.Println("error while parsing")
	}

	// Delete the book from the database
	book := models.DeleteBook(ID)

	// Marshal the Book struct into a JSON array and write it to the response
	res, _ := json.Marshal(book)
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Set the HTTP status code to 200 OK
	w.WriteHeader(http.StatusOK)
	// Write the response to the client
	w.Write(res)
}
 
// Updates a book in the database and returns it as JSON
//
// Parameters:
//  bookId: The ID of the book to update
//
// Returns:
//  A JSON representation of the updated Book struct
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a Book struct
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)

	// Get the book ID from the URL parameters
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	// Parse the book ID into an integer
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		// If there is an error, print it to the console
		fmt.Println("Error parsing")
	}

	// Retrieve the book from the database
	bookDetails, db := models.GetBookById(ID)

	// Update the book fields if the new values are not empty
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	// Save the updated book to the database
	db.Save(&bookDetails)

	// Marshal the updated book into a JSON array and write it to the response
	res, _ := json.Marshal(bookDetails)
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Set the HTTP status code to 200 OK
	w.WriteHeader(http.StatusOK)
	// Write the response to the client
	w.Write(res)
}
