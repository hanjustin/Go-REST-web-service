
## About

REST API endpoint:

`http://localhost:8080/api/notes`

supporting `GET`, `POST`, `PUT`, `DELETE` methods to create Note object and store in memory using dictionary `map[string]Note`.

Use `/notes/id` to access a specific Note.

## Prerequisites

Go 1.22

## Getting Started

1. Clone the repo
```sh
git clone https://github.com/hanjustin/Go-REST-web-service.git
```

2. Run main.go
```sh
go run .
```

3. Use `POST` to create Note

```
curl -X POST -H "Content-type: application/json" -d '{"title": "Title 1", "text": "Text 1"}' 'http://localhost:8080/api/notes'
```

4. Use `GET` to get all data

```
curl -X GET 'http://localhost:8080/api/notes'
```

## Output Example

### Create Note
```
{
    "id": "018ea579-e117-7ec1-9e1d-0f5b219076ca",
    "title": "Title 1",
    "text": "Text 1"
}
```

### Get ALL

```
[
    {
        "id": "018ea579-e117-7ec1-9e1d-0f5b219076ca",
        "title": "Title 1",
        "text": "Text 1"
    },
    {
        "id": "018ea579-fe87-7956-9e7b-f98e81cbb114",
        "title": "Title 2",
        "text": "Text 2"
    },
    {
        "id": "018ea57a-11b9-741e-9d91-78c6d96483e2",
        "title": "Title 3",
        "text": "Text 3"
    }
]
```