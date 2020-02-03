package router

import (
	"encoding/json"
	"net/http"

	. "../config/dao"
	. "../models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var daoTags = TagsDAO{}

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
	movies, err := daoTags.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

func GetTagByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := daoTags.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, movie)
}

func CreateTag(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var tags Tags
	if err := json.NewDecoder(r.Body).Decode(&tags); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	tags.ID = bson.NewObjectId()
	if err := daoTags.Create(tags); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, tags)
}
