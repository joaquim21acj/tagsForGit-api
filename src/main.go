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

// Variáveis para realizar a configuracao do banco
var configServer = config.Config{}
var daoTags = dao.TagsDAO{}

func init() {
	// Leitura das configuracoes de banco
	configServer.Read()

	daoTags.Server = configServer.Server
	daoTags.Database = configServer.Database
	daoTags.Connect()
}

func main() {
	// Mapeamento de rotas e definicao do header
	r := mux.NewRouter()
	headersOk := handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "content-type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "OPTIONS"})

	r.HandleFunc("/api/v1/tags", router.GetAllTags).Methods("GET")
	r.HandleFunc("/api/v1/tags", router.CreateTag).Methods("PATCH")
	// Função apenas para testar dados do banco
	r.HandleFunc("/api/v1/test", router.GetTagAllTest).Methods("GET")
	var port = ":3000"
	fmt.Println("Server running in port:", port)
	// Sobe o servidor com as configuracoes feitas acima
	log.Fatal(http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
