# Go-API

Simple API for Go using Mux with mock data.

Run:

```sh
go run main.go
```

### GET method:

Get all languages:

```sh
localhost:8080/language
```

Get one language, using an `ID`:

```sh
localhost:8080/language/<id>
```

### POST method:

Add a new language:

Sample data:

```
{"id": "1", "creator": "Guido Van Rossum", "name": "Python", "paradigm": "Multi-paradigm: functional, imperative, object-oriented, reflective"}
```

### DELETE method:

Delete a language, using an `ID`:

```sh
localhost:8080/language/<id>
```