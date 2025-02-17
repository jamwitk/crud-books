# CRUD Books

CRUD Books is a simple Go application that provides a RESTful API for managing a collection of books. It uses the Gin web framework and connects to a PostgreSQL database.

## Features

- List all books
- Create a new book
- Get details of a book by ID
- Update a book by ID
- Delete a book by ID

## Getting Started

### Prerequisites

- Go 1.22.3 or later
- PostgreSQL

### Installation

1. Clone the repository:

```sh
git clone https://github.com/jamwitk/crud-books.git
cd crud-books
```

2. Install the dependencies:

```sh
go mod download
```

3. Set up your PostgreSQL database and configure the environment variables accordingly.

### Running the Application

1. Load environment variables and connect to the database:

```go
initializers.LoadEnvVars()
initializers.Connect()
```

2. Run the application:

```sh
go run main.go
```

The application will start and listen on port 8080.

### API Endpoints

- `GET /books` - Retrieve a list of all books
- `POST /books` - Create a new book
- `GET /books/:id` - Retrieve a book by ID
- `PUT /books/:id` - Update a book by ID
- `DELETE /books/:id` - Delete a book by ID

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
