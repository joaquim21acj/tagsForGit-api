package main

import (
	"fmt"
	"log"
	"net/http"
	"tagsForGit-api/src/config"
	"tagsForGit-api/src/config/dao"
	router "tagsForGit-api/src/router"

	"github.com/gorilla/mux"
)

var configServer = config.Config{}
var daoTags = dao.TagsDAO{}

func init() {
	configServer.Read()

	daoTags.Server = configServer.Server
	daoTags.Database = configServer.Database
	daoTags.Connect()
}

func main() {
	r := mux.NewRouter()
	// r.HandleFunc("/api/v1/movies", router.GetAll).Methods("GET")
	// r.HandleFunc("/api/v1/movies/{id}", router.GetByID).Methods("GET")
	// r.HandleFunc("/api/v1/movies", router.Create).Methods("POST")
	// r.HandleFunc("/api/v1/movies/{id}", router.Update).Methods("PUT")
	// r.HandleFunc("/api/v1/movies/{id}", router.Delete).Methods("DELETE")

	r.HandleFunc("/api/v1/tags", router.GetAllTags).Methods("GET")
	r.HandleFunc("/api/v1/tags/{id}", router.GetTagByID).Methods("GET")
	r.HandleFunc("/api/v1/tags", router.CreateTag).Methods("POST")

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
