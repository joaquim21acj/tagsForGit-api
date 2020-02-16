package dao

import (
	"log"
	"tagsForGit-api/src/models"

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

func (m *TagsDAO) GetAllTags() ([]models.GitRepositories, error) {
	var tags []models.GitRepositories
	err := db.C(COLLECTION).Find(bson.M{}).All(&tags)
	return tags, err
}

func (m *TagsDAO) GetTagsByID(id string) (models.Tags, error) {
	var tags models.Tags
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&tags)
	return tags, err
}

func (m *TagsDAO) CreateTags(tags interface{}) error {
	err := db.C(COLLECTION).Insert(&tags)
	return err
}
