# Goflix

Goflix is a technical API inspire by Netflix which permits you to manage films and users

## Contributing

- Thomas Chastaingt

### Requirements

Make sure the following dependencies are installed:
- [Go](https://golang.org/dl/)

### Running



Start the backend :
```bash
$ cd Goflix
$ go mod download
$ go run main.go
```

## Codebase

#### Technologies
Here is a list of all the big technologies we use:

    - Go (backend)

#### Folder structure

```
.
├── _utils
│   └── utils.go
├── _user
│   └── user.go
├── _movie
│   └── movie.go
├── _server
│   ├── server.go
│   ├── routes_auth.go
│   ├── routes_movie.go
│   ├── middleware.go
│   └── routes.go
└── main.go

```

