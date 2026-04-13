# Go Task API + Swagger (v1)

This repository is upgraded into a clean Go REST API starter with a task CRUD service and
static OpenAPI documentation.

## Features

- Task CRUD endpoints under `/api/v1/tasks`
- Health check endpoint `/health`
- In-memory store with thread-safe access
- Static OpenAPI spec at `/swagger/openapi.json`
- Legacy Beego scaffold preserved under `legacy/`

## Endpoints

- `GET /health`
- `GET /api/v1/tasks`
- `POST /api/v1/tasks`
- `GET /api/v1/tasks/{taskID}`
- `PUT /api/v1/tasks/{taskID}`
- `DELETE /api/v1/tasks/{taskID}`

## Run

```bash
cp .env.example .env
go run ./cmd/server
```

## Docs

- Swagger info page: `http://localhost:8080/swagger/swagger.html`
- OpenAPI JSON: `http://localhost:8080/swagger/openapi.json`

## Project Structure

```text
golang-beego-restapi-swagger-v1/
├── cmd/server/main.go
├── internal/
│   ├── api/
│   ├── model/
│   └── store/
├── docs/
└── legacy/
```

## Notes

`go` toolchain was unavailable in the current environment, so runtime verification should be run on a Go-enabled machine.