package models

import "gopkg.in/mgo.v2/bson"

type Tags struct {
	ID   bson.ObjectId `bson:"_id" json:"id"`
	Tags []string      `bson:"name" json:"name"`
}
