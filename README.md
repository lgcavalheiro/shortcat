# Shortcat

## A minimalistic url shortener made in Go

## Requirements

- Go
- Docker
- Docker-compose (optional, only used for development)

## Running

1. Run `make dbu`, this will initiate MariaDB's and Adminer's docker images (for development only)
2. Run `make dev` on another terminal, this will initiate shortcat, which can be accessed @ http://localhost:8034 by default

## Api Routes

| Route      | Description                                   | Method | Params |
| ---------- | --------------------------------------------- | ------ | ------ |
| /api/      | Gets all shortcats registered in the database | GET    | None   |
| /api/      | Persists a new shortcat to the database       | POST   | None   |
| /api/auth/ | Route used to attempt admin login             | POST   | None   |

JSON body for `/api/auth/`:

```json
{
  "user": "string",
  "pwd": "string"
}
```

JSON body for `/api/`(POST):

```json
{
	"url": "string",
	"user_defined_id": "string|optional",
	"id_size": integer|optional
}
```

## Web Routes

| Route         | Description                                                      | Method | Params            |
| ------------- | ---------------------------------------------------------------- | ------ | ----------------- |
| /go/:shorturl | Redirects to the url that was shortened                          | GET    | shorturl - string |
| /             | Redirects to `/login`                                            | GET    | None              |
| /login        | Gets the login page, if already logged in, redirects to `/admin` | GET    | None              |
| /login        | Attempts to login admin                                          | POST   | None              |
| /admin        | Gets the admin dashboard page                                    | GET    | None              |
| /admin        | Gets admin page (temporary)                                      | POST   | None              |
| /logout       | Logs user out, redirecting to `/login`                           | GET    | None              |
