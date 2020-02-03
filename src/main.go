package main

import (
	"fmt"
	"log"
	"net/http"

	router "./router"

	. "./config"
	. "./config/dao"
	"github.com/gorilla/mux"
)

var dao = MoviesDAO{}
var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/movies", router.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/movies/{id}", router.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/movies", router.Create).Methods("POST")
	r.HandleFunc("/api/v1/movies/{id}", router.Update).Methods("PUT")
	r.HandleFunc("/api/v1/movies/{id}", router.Delete).Methods("DELETE")

	r.HandleFunc("/api/v1/tags", router.GetAllTags).Methods("GET")
	r.HandleFunc("/api/v1/tags/{id}", router.GetTagByID).Methods("GET")
	r.HandleFunc("/api/v1/tags", router.CreateTag).Methods("POST")

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
