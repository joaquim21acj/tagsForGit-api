package router

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	. "../config/dao"
	. "../models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var daoTags = TagsDAO{}
var urlGraphQL = "https://api.github.com/graphql"
var token = "3f8053ee8083c9e7c6894fa1757b1f0c38898e60"

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

	urlQuery, err = url.Parse(r.URL.Path)
	if (err != nil) || (urlQuery.Query["userLogin"] == nil) {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlGraphQL, nil)
	req.Header.set("query", getRepositories(userLogin))
	req.Header.Set("Authorization", "Bearer"+token)
	res, err := client.Do(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	log.Println(string(body))

	tags, err := daoTags.GetAllTags()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, tags)
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
	var tags Tags
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
