package router

import (
	"encoding/json"
	"log"
	"net/http"
	"tagsForGit-api/src/config/dao"
	"tagsForGit-api/src/models"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var daoTags = dao.TagsDAO{}
var urlGraphQL = "https://api.github.com/graphql"
var token = "93937194fe160657e94588bb4dc1ee0c7ab62550"

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

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlGraphQL, nil)
	req.PostFormValue(models.GetRepositories(userLogin[0]))
	req.Header.Set("Authorization", "Bearer "+token)
	log.Println(req)
	res, err := client.Do(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Println(res)
	respondWithError(w, 200, "ok")
	log.Println("Url Param 'userLogin' is here")
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
	tags.ID = bson.NewObjectId()
	if err := daoTags.CreateTags(tags); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, tags)
}
