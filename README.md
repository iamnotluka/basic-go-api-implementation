# Book Inventory Management API

A simple API written in go. Goal of this project is just to get familiar with how go works, how to execute the APIs locally.

### API Endpoints

Below are the available endpoints in the project:

- GET /books
  Retrieves a list of all books in the inventory.

- GET /books/:id
  Retrieves a specific book by it's ID.

- POST /books
  Adds a new book to the inventory. Requires a JSON payload with `id`, `title`, `author` and `quantity` (use body.json file, look at instructions on how to test)

- PATCH /checkout
  Checks out a book from the inventory using the book's ID passed as a querty parameter. Reduces the quantity of the specified book by one.

- PATCH /return
  Returns a book to the inventory using the book's ID passed as a query parameter. Increases the quantity of the specified book by one.

#### Prereqs

Before running the project, ensure to have go installed on your machine. You can download and install Go from https://go.dev/dl/.

Additionally, project uses Gin web framework, which can be installed using:

```
go get -u github.com/gin-gonic/gin
```

#### Installation

Clone the repo to your local machine:

```
git clone https://github.com/iamnotluka/basic-go-api-implementation.git
cd basic-go-api-implementation
```

### Testing

#### Running the Application

To run the app, navigate to the project dir and execute:

```
go run .
```

You should see output that looks like this to validate:

```
➜  basic-go-api-implementation git:(main) ✗ go run .
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /books                    --> main.getBooks (3 handlers)
[GIN-debug] GET    /books/:id                --> main.getBookById (3 handlers)
[GIN-debug] POST   /books                    --> main.createBook (3 handlers)
[GIN-debug] PATCH  /checkout                 --> main.checkoutBook (3 handlers)
[GIN-debug] PATCH  /return                   --> main.returnBook (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
```

#### Calling the APIs

```
➜  basic-go-api-implementation git:(main) ✗ curl localhost:8080/books
[
    {
        "id": "1",
        "title": "Beyond Good and Evil",
        "author": "Fredrich Nietzhe",
        "quantity": 4
    },
    {
        "id": "2",
        "title": "Think and Grow Rich",
        "author": "Napoleon Hill",
        "quantity": 5
    },
    {
        "id": "3",
        "title": "The Changing World Order",
        "author": "Ray Dalio",
        "quantity": 6
    }
]%
```

To make a call locally with the custom body.json use these flags:

```
--include --header "Content-Type: application/json" -d @body.json --request "POST"
```
