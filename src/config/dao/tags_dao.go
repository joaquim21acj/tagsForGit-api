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

func (m *TagsDAO) GetAllTags() ([]models.User, error) {
	var tags []models.User
	err := db.C(COLLECTION).Find(bson.M{}).All(&tags)
	// err := db.C(COLLECTION).Find(bson.M{}).All(&tags)
	return tags, err
}

func (m *TagsDAO) GetTagsByUser(login string) (models.User, error) {
	var Tags models.User
	// err := db.C(COLLECTION).Find(bson.D{}).One(Tags)
	err := db.C(COLLECTION).Find(bson.D{{"user.login", login}}).One(&Tags)
	return Tags, err
}

func (m *TagsDAO) GetTagsByID(id string) (models.Tag, error) {
	var tags models.Tag
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&tags)
	return tags, err
}

func (m *TagsDAO) CreateTags(tags interface{}) error {
	err := db.C(COLLECTION).Insert(tags)
	return err
}
func (m *TagsDAO) UpdateTags(login string, user models.User) error {
	err := db.C(COLLECTION).Update(bson.D{{"user.login", login}}, &user)
	// Id(bson.ObjectIdHex(id), &user)
	return err
}
