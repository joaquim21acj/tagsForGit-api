package models

import "gopkg.in/mgo.v2/bson"

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
