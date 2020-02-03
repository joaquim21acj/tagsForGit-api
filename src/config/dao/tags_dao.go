package dao

import (
	"log"

	. "../../models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TagsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "tags"
)

func (m *TagsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *TagsDAO) GetAllTags() ([]Tags, error) {
	var tags []Tags
	err := db.C(COLLECTION).Find(bson.M{}).All(&tags)
	return tags, err
}

func (m *TagsDAO) GetTagsByID(id string) (Tags, error) {
	var tags Tags
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&tags)
	return tags, err
}

func (m *TagsDAO) CreateTags(tags Tags) error {
	err := db.C(COLLECTION).Insert(&tags)
	return err
}
