package models

import (
	"github.com/SAMGREENHEAD/Book_management/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB


type Book struct {
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
	
}

// Initializes the database connection and migrates the Book struct to the database
func init() {
	// Connects to the database using the DSN provided in the config file
	config.Connect()

	// Retrieves the global database connection variable
	db = config.GetDB()

	// Automatically migrates the Book struct to the database
	// This creates the tables in the database if they don't already exist
	db.AutoMigrate(&Book{})
}

// Creates a new book in the database and returns it as a pointer to the Book struct
//
// Parameters: None
//
// Returns: A pointer to the newly created Book struct
func (b *Book) CreateBook() *Book {
	// Creates a new record in the database for the Book struct
	// This sets the ID field to 0, as Gorm automatically assigns an ID
	db.NewRecord(b)
	
	// Creates the book in the database
	// This will insert the new Book struct into the database
	// and assign an ID to it
	db.Create(&b)
	
	// Returns a pointer to the newly created Book struct
	// so it can be easily accessed by the caller
	return b
}


	// Retrieves all books from the database and returns them as a slice of Book structs
	//
	// Parameters: None
	//
	// Returns: A slice of Book structs containing all books from the database
func GetAllBooks() []Book {
	var Books []Book
	// Finds all books in the database and stores them in the Books slice
	// This uses the Find method of the Gorm database to retrieve all records from the books table
	db.Find(&Books)
	// Returns the slice of Book structs containing all books from the database
	return Books
	}

	// Retrieves a single book from the database and returns it as a pointer to the Book struct
	// and the Gorm DB
	//
	// Parameters:
	//  Id: The ID of the book to retrieve
	//
	// Returns:
	//  A pointer to the Book struct containing the book with the given ID
	//  The Gorm DB used to retrieve the book
func GetBookById(Id int64) (*Book,*gorm.DB)  {
	// Finds the book in the database with the given ID
	// This uses the Where method of the Gorm database to find a record in the books table with the given ID
	var getBook Book
	db:=db.Where("ID=?",Id).Find(&getBook)
	// Returns a pointer to the Book struct containing the book with the given ID
	// and the Gorm DB used to retrieve the book
	return &getBook,db
}



	// Deletes a book from the database and returns the deleted book
	//
	// Parameters:
	//  Id: The ID of the book to delete
	//
	// Returns:
	//  The deleted Book struct
func DeleteBook(Id int64) Book {
	// Finds the book in the database with the given ID
	// and deletes it
	// This uses the Where method of the Gorm database to find a record in the books table with the given ID
	// and the Delete method to delete it
	var book Book
	db.Where("ID=?",Id).Delete(&book)
	// Returns the deleted Book struct
	return book
}




