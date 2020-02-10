package router

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"tagsForGit-api/src/config/dao"
	"tagsForGit-api/src/models"

	"github.com/friendsofgo/graphiql"
	"github.com/gorilla/mux"
)

var daoTags = dao.TagsDAO{}
var urlGraphQL = "https://api.github.com/graphql"
var token = "d8193fb19afe1daf8fd0d6443614bdc2d2d2a027"

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
	userLogin, ok := r.URL.Query()["userLogin"]
	if !ok || len(userLogin[0]) < 1 {
		log.Println("Url Param 'userLogin' is missing")
		respondWithError(w, http.StatusBadRequest, "deu ruim")
		return
	}

	// Realiza o processo de converter a string para o formato do graphql
	var grql = graphiql.Handler{}
	log.Println(grql.ServeHTTPMustParseSchema(s, &models.GetRepositories(userLogin)))
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlGraphQL, bytes.NewBufferString(models.GetRepositories(userLogin[0])))
	req.Header.Set("Authorization", "Bearer "+token)
	log.Println(req.Body)
	res, err := client.Do(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Println(res.StatusCode)
	if res.StatusCode == 200 {
		var tags = models.Tags{}
		getJson(res, tags)
		respondWithError(w, 200, "ok")
		log.Println("Url Param 'userLogin' is here")
		return
	}

	b, _ := ioutil.ReadAll(res.Body)
	log.Fatal(string(b))
	respondWithError(w, 200, "ok")
	// body, err := ioutil.ReadAll(res.Body)
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
