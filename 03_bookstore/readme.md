#
#
#
##  Make movies crud using go

### Before starting the web server run the following command:
```sh
go mod init github.com/username/bookstore
go mod tidy
```

### To start web server, run :
```sh
go build
go run cmd/main/main.go 
```

### You can call the following routes in postman and see results 

```go
    	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
```

Enjoy it :)
