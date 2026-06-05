# Bulletin API

A Scalable API created for simply store and list Bulletins.

## Database Schema

The application uses PostgreSQL with the following schema for bulletins:

```sql
CREATE TABLE IF NOT EXISTS bulletins (
    id serial PRIMARY KEY,
    author text NOT NULL,
    content text NOT NULL,
    created_at timestamp with time zone DEFAULT current_timestamp
);
```

## API Documentation

### 1. Get All Bulletins
Retrieve a list of all bulletins.

- **URL:** `/board`
- **Method:** `GET`
- **Success Response:**
  - **Code:** 200 OK
  - **Content:** Array of bulletin objects.

### 2. Add Bulletin
Create a new bulletin.

- **URL:** `/board`
- **Method:** `POST`
- **Data Params:**
  ```json
  {
    "author": "String",
    "content": "String"
  }
  ```
- **Success Response:**
  - **Code:** 200 OK
  - **Content:** `{"status": "ok"}`

## Docker Support

The repository includes Docker support for easy database and API setup.
A `docker-compose.yml` file is provided to spin up both the PostgreSQL database and the API service.

## Running Locally

Build first

```makefile
make build
```
Then 

```makefile
make run
```