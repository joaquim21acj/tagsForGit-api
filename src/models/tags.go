package models

import "gopkg.in/mgo.v2/bson"

type Tags struct {
	ID  bson.ObjectId `bson:"_id_tags" json:"id"`
	Tag string        `bson:"tag" json:"tag"`
}
