# hse-24-parallele-verteilte-systeme
Git repository for lab by Fabian Etzler and Luis Urbitsch

## API Endpoints

### Create Todo Item

Create a new Todo item with the provided title.

- **URL:** `/todos/:title`
- **Method:** `POST`
- **URL Params:**
  - `title`: Title of the Todo item (string)
- **Success Response:**
  - **Code:** `201 Created`
- **Error Responses:**
  - **Code:** `500 Internal Server Error`
    - **Content:** `{ error: "Internal Server Error" }`

### Update Todo Item

Update an existing Todo item's title by ID.

- **URL:** `/todos/:id/:title`
- **Method:** `PUT`
- **URL Params:**
  - `id`: ID of the Todo item to update (integer)
  - `title`: New title for the Todo item (string)
- **Success Response:**
  - **Code:** `200 OK`
- **Error Responses:**
  - **Code:** `500 Internal Server Error`
    - **Content:** `{ error: "Internal Server Error" }`
  - **Code:** `404 Not Found`
    - **Content:** `{ error: "Todo item not found" }`

### Get Todo Item by ID

Retrieve a Todo item by its ID.

- **URL:** `/todos/:id`
- **Method:** `GET`
- **URL Params:**
  - `id`: ID of the Todo item to retrieve (integer)
- **Success Response:**
  - **Code:** `200 OK`
    - **Content:** `{ id: 1, title: "Buy Groceries" }`
- **Error Responses:**
  - **Code:** `404 Not Found`
    - **Content:** `{ error: "Todo item not found" }`

### Get All Todo Items

Retrieve all Todo items.

- **URL:** `/todos`
- **Method:** `GET`
- **Success Response:**
  - **Code:** `200 OK`
    - **Content:** `[{ id: 1, title: "Buy Groceries" }, { id: 2, title: "Walk the Dog" }, ...]`
- **Error Responses:**
  - **Code:** `500 Internal Server Error`
    - **Content:** `{ error: "Internal Server Error" }`

### Delete Todo Item

Delete a Todo item by its ID.

- **URL:** `/todos/:id`
- **Method:** `DELETE`
- **URL Params:**
  - `id`: ID of the Todo item to delete (integer)
- **Success Response:**
  - **Code:** `200 OK`
- **Error Responses:**
  - **Code:** `500 Internal Server Error`
    - **Content:** `{ error: "Internal Server Error" }`
  - **Code:** `404 Not Found`
    - **Content:** `{ error: "Todo item not found" }`
