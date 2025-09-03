# Go GraphQL Book Library

This project is a simple GraphQL server for managing a collection of books. It's built with Go, gqlgen, Fiber, and GORM.

## Features

- **GraphQL API:** Provides a GraphQL API for CRUD operations on books.
- **SQLite Database:** Uses a SQLite database to store the book data.
- **GORM:** Uses GORM as the ORM for database interactions.
- **Fiber:** Uses Fiber as the HTTP server.
- **GraphQL Playground:** Includes a GraphQL playground for easy testing of the API.

## Getting Started

### Prerequisites

- Go 1.25.0 or higher
- A C compiler (for the SQLite driver)

### Installation

1.  Clone the repository:
    ```bash
    git clone https://github.com/Abhishek031bit/go-graphql.git
    ```
2.  Navigate to the project directory:
    ```bash
    cd go-graphql
    ```
3.  Install the dependencies:
    ```bash
    go mod tidy
    ```

### Running the Application

1.  Run the server:
    ```bash
    go run server.go
    ```
2.  The server will start on port 8080. You can access the GraphQL playground at `http://localhost:8080/`.

## API Endpoints

The GraphQL endpoint is available at `/query`. You can use the GraphQL playground to interact with the API.

### Queries

- `books`: Returns a list of all books.
- `book(id: ID!)`: Returns a single book by its ID.

### Mutations

- `createBook(input: NewBook!)`: Creates a new book.
- `updateBook(id: ID!, input: NewBook!)`: Updates an existing book.
- `deleteBook(id: ID!)`: Deletes a book.

## Technologies Used

- [Go](https://golang.org/)
- [GraphQL](https://graphql.org/)
- [gqlgen](https://gqlgen.com/)
- [Fiber](https://gofiber.io/)
- [GORM](https://gorm.io/)
- [SQLite](https://www.sqlite.org/)
