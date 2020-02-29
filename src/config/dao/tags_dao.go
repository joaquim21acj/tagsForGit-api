package dao

import (
	"log"
	"tagsForGit-api/src/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TagsDAO é uma struct para armazenar as configuracoes do servidor
// e do database
type TagsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// COLLECTION usada na conexao
	COLLECTION = "tags"
)

// Connect faz a conexão com o banco e não possui retorno
func (m *TagsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// GetAllTags busca todos as tags que cada usuário possui
// retorna as tags dos usuários e error se houver
func (m *TagsDAO) GetAllTags() ([]models.User, error) {
	var tags []models.User
	err := db.C(COLLECTION).Find(bson.M{}).All(&tags)
	return tags, err
}

// GetTagsByUser realiza uma busca de todas as tags por um usuário
// retorna as tags para o usuário e error se houver
func (m *TagsDAO) GetTagsByUser(login string) (models.User, error) {
	var Tags models.User
	err := db.C(COLLECTION).Find(bson.D{{"user.login", login}}).One(&Tags)
	return Tags, err
}

// CreateTags recebe uma interface correspondente ao usuário com seus repositorios
// retorna erro caso haja
func (m *TagsDAO) CreateTags(tags interface{}) error {
	err := db.C(COLLECTION).Insert(tags)
	return err
}

// UpdateTags recebe o login do usuário e o objeto do usuário e o atualiza
// retorna erro caso haja
func (m *TagsDAO) UpdateTags(login string, user models.User) error {
	err := db.C(COLLECTION).Update(bson.D{{"user.login", login}}, &user)
	return err
}
