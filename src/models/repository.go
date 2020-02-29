package models

// GitRepositories é a raiz dos dados que são retornados pela api do GitHub V4
type GitRepositories struct {
	Data struct {
		User User
	} `json:"data"`
}
