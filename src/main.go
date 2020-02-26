package main

import (
	"fmt"
	"log"
	"net/http"
	"tagsForGit-api/src/config"
	"tagsForGit-api/src/config/dao"
	router "tagsForGit-api/src/router"

	"github.com/gorilla/handlers"
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
	headersOk := handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "content-type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "OPTIONS"})

	r.HandleFunc("/api/v1/tags", router.GetAllTags).Methods("GET")
	// r.HandleFunc("/api/v1/tags/{id}", router.GetTagByID).Methods("GET")
	r.HandleFunc("/api/v1/tags", router.CreateTag).Methods("PATCH")
	// Função apenas para testar dados do banco
	r.HandleFunc("/api/v1/test", router.GetTagAllTest).Methods("GET")
	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
