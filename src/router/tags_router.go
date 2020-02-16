package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"tagsForGit-api/src/config/dao"
	"tagsForGit-api/src/models"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

var daoTags = dao.TagsDAO{}

type graphResponse struct {
	Data interface{}
}

var urlGraphQL = "https://api.github.com/graphql"
var token = "14556d98644a4a02fe68d9325cc914020f7a4602"

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAllTags(w http.ResponseWriter, r *http.Request) {
	//Função que recebe a requisição para buscara todas as tags que o usuário possui
	//Validação se recebeu o userlogin para buscar na api do git
	userLogin, ok := r.URL.Query()["userLogin"]
	if !ok || len(userLogin[0]) < 1 {
		log.Println("Url Param 'userLogin' is missing")
		respondWithError(w, http.StatusBadRequest, "deu ruim")
		return
	}

	//Criação da query de login usando estrutura de query
	//A estrutura pode evoluir para separar a query das variáveis
	var queryString = models.GetRepositories(userLogin[0])
	requestBodyObj := struct {
		Query string `json:"query"`
	}{
		Query: queryString}
	//Encoder da estrutura de query para bytes
	var requestBody bytes.Buffer
	if err := json.NewEncoder(&requestBody).Encode(requestBodyObj); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	//Início das configurações da requisição http
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlGraphQL, &requestBody)
	req.Header.Set("Authorization", "Bearer "+token)
	log.Println(req.Body)
	res, err := client.Do(req)
	//fim da requisição
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	defer res.Body.Close()
	//Pega o resultado que vem no body e copia para o buffer
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, res.Body); err != nil {
		log.Println(errors.Wrap(err, "reading body"))
		return
	}
	gr := &graphResponse{
		Data: models.Tags{},
	}

	if err := json.NewDecoder(&buf).Decode(&gr); err != nil {
		if res.StatusCode != http.StatusOK {
			log.Println(fmt.Errorf("graphql: server returned a non-200 status code: %v", res.StatusCode))
		}
		log.Println(errors.Wrap(err, "decoding response"))
	}
	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// log.Println(res.StatusCode)
	// if res.StatusCode == 200 {
	// 	log.Println(res.Body)
	// 	var tags = getJson(res, models.Tags{})
	// 	// var tags []models.Tags
	// 	// tags = daoTags.GetAllTags()
	// 	respondWithJson(w, 200, tags)
	// 	log.Println("Url Param 'userLogin' is here")
	// 	return
	// }

	// b, _ := ioutil.ReadAll(res.Body)
	// log.Fatal(string(b))
	respondWithJson(w, 200, gr)
	// body, err := ioutil.ReadAll(res.Body)gr
	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, err.Error())
	// }
	// log.Println(string(body))

	// tags, err := daoTags.GetAllTags()
	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// respondWithJson(w, http.StatusOK, tags)
}

func GetTagByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag, err := daoTags.GetTagsByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, tag)
}

func CreateTag(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var tags models.Tags
	if err := json.NewDecoder(r.Body).Decode(&tags); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := daoTags.CreateTags(tags); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, tags)
}

func getJson(r *http.Response, target interface{}) error {
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
func getJsonMensage(r *http.Response, target string) error {
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
