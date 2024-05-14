# hse-24-parallele-verteilte-systeme
Git repository for lab by Fabian Etzler and Luis Urbitsch

## DockerHub image
[Link to DockerHub](https://hub.docker.com/r/faetzler/hse-24-parallele-verteilte-systeme)
To use our image use ```image: faetzler/hse-24-parallele-verteilte-systeme:v0.1``` in docker-compose.yaml

## API Endpoints

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
