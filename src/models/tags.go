package models

type Tag struct {
	// ID  bson.ObjectId `bson:"_id_tags" json:"id"`
	Tag string `bson:"tag" json:"tag"`
}
