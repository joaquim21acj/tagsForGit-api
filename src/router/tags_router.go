package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"tagsForGit-api/src/config/dao"
	"tagsForGit-api/src/models"

	"github.com/pkg/errors"
)

var daoTags = dao.TagsDAO{}

type graphResponse struct {
	Data interface{}
}

var urlGraphQL = "https://api.github.com/graphql"
var token = "daba0d333d063bbb02869d055daa2583024caf05"

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetTagAllTest(w http.ResponseWriter, r *http.Request) {
	tags, err := daoTags.GetAllTags()
	log.Println(tags)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, tags)
}

func GetAllTags(w http.ResponseWriter, r *http.Request) {
	//Função que recebe a requisição para buscara todas as tags que o usuário possui
	//Validação se recebeu o userlogin para buscar na api do git
	userLogin, ok := r.URL.Query()["userLogin"]
	if !ok || len(userLogin[0]) < 1 {
		log.Println("Url Param 'userLogin' is missing")
		respondWithError(w, http.StatusBadRequest, "deu ruim")
		return
	}

	//Criação da query de login usando estrutura de query
	//A estrutura pode evoluir para separar a query das variáveis
	var queryString = models.GetRepositories(userLogin[0])
	requestBodyObj := struct {
		Query string `json:"query"`
	}{
		Query: queryString}
	//Encoder da estrutura de query para bytes
	var requestBody bytes.Buffer
	if err := json.NewEncoder(&requestBody).Encode(requestBodyObj); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	//Início das configurações da requisição http
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlGraphQL, &requestBody)
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := client.Do(req)
	//fim da requisição
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	defer res.Body.Close()
	//Pega o resultado que vem no body e copia para o buffer
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, res.Body); err != nil {
		log.Println(errors.Wrap(err, "reading body"))
		return
	}
	//definição da variável com o formato dos dados recebidos
	gr := &graphResponse{
		Data: models.GitRepositories{},
	}
	//Transforma os dados de bytes, em buffer, para o formato da struct criada
	if err := json.NewDecoder(&buf).Decode(&gr); err != nil {
		if res.StatusCode != http.StatusOK {
			log.Println(fmt.Errorf("graphql: server returned a non-200 status code: %v", res.StatusCode))
		}
		log.Println(errors.Wrap(err, "decoding response"))
	}

	user, err := daoTags.GetTagsByUser(userLogin[0])

	if err != nil {
		if strings.Compare(err.Error(), "not found") == 0 {
			if err := daoTags.CreateTags(gr.Data); err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			respondWithJson(w, http.StatusCreated, gr.Data)
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJson(w, http.StatusOK, user)
	return
}

// func GetTagByID(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	tag, err := daoTags.GetTagsByID(params["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
// 		return
// 	}
// 	respondWithJson(w, http.StatusOK, tag)
// }

func CreateTag(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// Decodifica o objeto que vem no corpo do request e o transforma em tags
	var repository models.Node

	//Pega o resultado que vem no body e copia para o buffer
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r.Body); err != nil {
		log.Println(errors.Wrap(err, "reading body"))
		return
	}

	//Transforma os dados de bytes, em buffer, para o formato da struct criada
	if err := json.NewDecoder(&buf).Decode(&repository); err != nil {
		log.Println(errors.Wrap(err, "decoding response"))
	}

	// if err := json.NewDecoder(r.Body).Decode(&repository); err != nil {
	// 	respondWithError(w, http.StatusBadRequest, "Objeto inválido")
	// 	return
	// }
	// Pega parametro userLogin da url
	userLogin, ok := r.URL.Query()["userLogin"]
	if !ok || len(userLogin[0]) < 1 {
		log.Println("Url Param 'userLogin' is missing")
		respondWithError(w, http.StatusBadRequest, "deu ruim")
		return
	}

	user, err := daoTags.GetTagsByUser(userLogin[0])

	if err != nil {
		if strings.Compare(err.Error(), "not found") == 0 {
			respondWithError(w, http.StatusNotFound, "User not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	var check int = 0
	for i, s := range user.User.StarredRepositories.Edges {
		if s.NodeRepositories.ID == repository.ID {
			user.User.StarredRepositories.Edges[i].NodeRepositories = repository
			check++
		}
	}
	for _, s := range user.User.StarredRepositories.Edges {
		println(s.NodeRepositories.Name)
	}

	if check == 0 {
		respondWithError(w, http.StatusBadRequest, "Esse respositorio não faz parte dos itens curtidos")
		return
	}

	if err := daoTags.UpdateTags(userLogin[0], user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, "ok")
}

func getJson(r *http.Response, target interface{}) error {
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
func getJsonMensage(r *http.Response, target string) error {
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
