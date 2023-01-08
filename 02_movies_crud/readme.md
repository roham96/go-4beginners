#
#
#
##  Make movies crud using go

### Before starting the web server run the following command:
```sh
go install github.com/gorilla/mux@latest
```
and run 
```sh
go mod init 02_movies_crud
go mod tidy
```
***  (02_movies_crud not working ! use the current working directory)

### To start web server, run :
```sh
go run main.go 
```

### You can call the following routes in postman and see results 

```go
    r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
```

Enjoy it :)