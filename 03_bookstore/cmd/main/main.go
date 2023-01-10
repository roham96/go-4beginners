package main


import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"github.com/roham96/go-bookstore/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}