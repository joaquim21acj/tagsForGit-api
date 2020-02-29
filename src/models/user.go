package models

import "gopkg.in/mgo.v2/bson"

// User é o conjunto de dados que é recebida pela api do GitHubV4
// o bson.ObjectId representa o id que é gerado no bando pelo Mongo
type User struct {
	ID   bson.ObjectId `bson:"_id" json:"id"`
	User struct {
		ID                  string `bson:"id" json:"id_user"`
		Login               string `bson:"login" json:"login"`
		StarredRepositories struct {
			Edges []struct {
				NodeRepositories Node `bson:"node" json:"node"`
			} `bson:"edges" json:"edges"`
		} `bson:"starredRepositories" json:"starredRepositories"`
	} `bson:"user" json:"user"`
}
