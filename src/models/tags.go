package models

import "gopkg.in/mgo.v2/bson"

type Tags struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	Tags          []string      `bson:"tags" json:"tags"`
	ID_repository string        `bson: "id_repository" json:"id_repository"`
}
