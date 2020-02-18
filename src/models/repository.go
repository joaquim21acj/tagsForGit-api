package models

type GitRepositories struct {
	Data struct {
		User struct {
			ID                  string `json:"id_user"`
			StarredRepositories struct {
				Edges []struct {
					NodeRepositories Node `json:"node"`
				} `json:"edges"`
			} `json:"starredRepositories"`
		} `json:"user"`
	} `json:"data"`
}
