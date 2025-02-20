# go-library-api

⏱️ **Time Limit**: Set by the admin  
🚀 **Difficulty**: Junior Level  
💻 **Stack**: Golang

## Implement Simple CRUD API for Library Management

Your task is to implement a simple CRUD API using an in-memory database underneath.

✅ **Core Technologies**

- **Golang** 1.24.0
- **In-memory storage** (no databases, all data stored in memory)

## Details:

1.  The task must be solved using only **net/http** and provided modules in **go.mod**. Any libraries and packages **are prohibited**.
2.  API path `/books`:

    - **GET** `/books` - returns a list of all books. there should be a possibility to filter books by `author` , `title` , and `genre ` query parameters and sort by `year ` query parameters. If no query parameters are provided, all books should be returned.
    - **GET** `/books/:id` - returns a single book by id.
    - **POST** `/books` - creates a new book.
    - **PUT** `/books/:id` - updates a book by id.

3.  The book object should have the following properties:

    - `id` - unique identifier
    - `title` - title of the book
    - `author` - author of the book
    - `genre` - genre of the book
    - `year` - year of publication

4.  Requests to non-existing endpoints (e.g., `/some-non/existing/resource`) should be handled.
5.  Internal server errors should be handled and processed correctly.

### Start Development

```bash
go run main.go
```

### Run Tests

```bash
go test ./..
```

## Evaluation Criteria

📊 **Scoring Breakdown**

| Category       | Weight |
| -------------- | ------ |
| Functionality  | 60%    |
| Code Quality   | 25%    |
| Error Handling | 15%    |

## Tips for Success

### Time Management

- **0-30 mins**: Implement POST and GET endpoints.
- **30-60 mins**: Add PATCH and DELETE functionality.
- **60-90 mins**: Add validation, error handling, and filtering.
- **90-120 mins**: Final testing, error checking, and refinement.

### Priority Order

1. **Basic CRUD operations**: GET, POST, PATCH, DELETE.
2. **Input validation**: Validate required fields and reject invalid data.
3. **Error handling**: Provide meaningful error messages

### Test Early

- Use **curl** or **Postman** for endpoint testing.
- **Test all error scenarios** (e.g., invalid inputs, missing items).

---

**Good luck!** 🍀


