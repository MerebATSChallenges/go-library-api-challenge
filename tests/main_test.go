package tests

import (
	"go-library-api/api"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var server *httptest.Server

func setup() {
	server = httptest.NewServer(api.SetupRouter())
}

func teardown() {
	server.Close()
}

func TestBookCRUD(t *testing.T) {
	setup()
	defer teardown()

	fmt.Println("This is the server ", server.URL)
	t.Run("Test Get", func(t *testing.T) {

		var expectedBooks []api.Book
		for _, book := range api.Books {
			expectedBooks = append(expectedBooks, book)
		}
	
		res, err := http.Get(server.URL + "/books")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
	
		var actualBooks []api.Book
		err = json.NewDecoder(res.Body).Decode(&actualBooks)
		assert.NoError(t, err)
	
		assert.ElementsMatch(t, expectedBooks, actualBooks)
	})
	

	t.Run("Test Filter by Author", func(t *testing.T) {

        res, err := http.Get(server.URL + "/books?author=John")
        assert.NoError(t, err)

        assert.Equal(t, http.StatusOK, res.StatusCode)

        var actualBooks []api.Book
        err = json.NewDecoder(res.Body).Decode(&actualBooks)
        assert.NoError(t, err)

        assert.Len(t, actualBooks, 2)
        assert.Contains(t, actualBooks[0].Author, "John")
        assert.Contains(t, actualBooks[1].Author, "John")
    })

	// Test filtering by genre
    t.Run("Test Filter by Genre", func(t *testing.T) {
        res, err := http.Get(server.URL + "/books?genre=Fiction")
        assert.NoError(t, err)

        assert.Equal(t, http.StatusOK, res.StatusCode)

        var actualBooks []api.Book
        err = json.NewDecoder(res.Body).Decode(&actualBooks)
        assert.NoError(t, err)

        // Check that only books with "Fiction" genre are returned
        assert.Len(t, actualBooks, 2)
        assert.Contains(t, actualBooks[0].Genre, "Fiction")
        assert.Contains(t, actualBooks[1].Genre, "Fiction")
    })

    // Test filtering by title
    t.Run("Test Filter by Title", func(t *testing.T) {
        res, err := http.Get(server.URL + "/books?title=Go")
        assert.NoError(t, err)

        assert.Equal(t, http.StatusOK, res.StatusCode)

        var actualBooks []api.Book
        err = json.NewDecoder(res.Body).Decode(&actualBooks)
        assert.NoError(t, err)

        // Check that only books with "Go" in the title are returned
        assert.Len(t, actualBooks, 2)
        assert.Contains(t, actualBooks[0].Title, "Go")
        assert.Contains(t, actualBooks[1].Title, "Go")
    })

    // Test sorting by year
    t.Run("Test Sort by Year", func(t *testing.T) {
        res, err := http.Get(server.URL + "/books?year=2020")
        assert.NoError(t, err)


        assert.Equal(t, http.StatusOK, res.StatusCode)

		var actualBooks []api.Book
        err = json.NewDecoder(res.Body).Decode(&actualBooks)
        assert.NoError(t, err)

        // Check that the books are sorted by year (ascending)
        assert.Len(t, actualBooks, 4)
        assert.True(t, actualBooks[0].Year <= actualBooks[1].Year)
        assert.True(t, actualBooks[1].Year <= actualBooks[2].Year)
        assert.True(t, actualBooks[2].Year <= actualBooks[3].Year)
    })

	t.Run("Test Post", func(t *testing.T) {
		newBook := api.Book{
			Title:   "The Da Vinci Code",
			Genre:   "Mistery",
			Author:  "Dan Brown",
			Year:    2003,
		}

		payload, _ := json.Marshal(newBook)
		res, err := http.Post(server.URL+"/books", "application/json", bytes.NewBuffer(payload))
		books := api.Books
		fmt.Println("This is books db ", books)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
	t.Run("Test Get By ID", func(t *testing.T) {
		res, err := http.Get(server.URL + "/books/1")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
	t.Run("Test Post Validation All Empty", func(t *testing.T) {
		res, err := http.Post(server.URL+"/books", "application/json", bytes.NewBuffer([]byte(`{}`)))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Test Post Validation Author Empty", func(t *testing.T) {

		data := api.Book{
			Title: "The Da Vinci Code",
			Genre: "Mistery",
			Year:  2003,	
		}
		payload, _ := json.Marshal(data)
		res, err := http.Post(server.URL+"/books", "application/json", bytes.NewBuffer(payload))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Test Put", func(t *testing.T) {
		updatedBook := api.Book{
			Title:   "The Da Vinci Code",
			Genre:   "Mistery",
			Author:  "Dan Brown",
			Year:    2003,
		}

		payload, _ := json.Marshal(updatedBook)
		req, _ := http.NewRequest(http.MethodPut, server.URL+"/books/1", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)

	})

	t.Run("Test Delete", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, server.URL+"/books/1", nil)
		client := &http.Client{}
		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)

	})

	t.Run("Test Non-Existing Book", func(t *testing.T) {
		res, err := http.Get(server.URL + "/books/9999")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})

	t.Run("Test Non-Existing Endpoint", func(t *testing.T) {
		res, err := http.Get(server.URL + "/test/non-existing/endpoint")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
}