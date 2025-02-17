package api

import (
	"net/http"
	"sync"
)

// Book structure
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Year   int    `json:"year"`
}

// In-memory storage
var (
	Books   =  map[int]Book{
					1: {Author: "John Doe", Title: "Go Programming", Genre: "Programming", Year: 2020},
					2: {Author: "Jane Smith", Title: "Fictional World", Genre: "Fiction", Year: 2019},
					3: {Author: "John Doe", Title: "Mastering Go", Genre: "Programming", Year: 2021},
					4: {Author: "Alice Green", Title: "Fictional Adventure", Genre: "Fiction", Year: 2022},
				}
	mu      sync.RWMutex
	bookIDCounter = 5
)

// SetupRouter initializes HTTP routes
func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// TODO: define routes here


	return mux
}


// TODO implement CRUD of book