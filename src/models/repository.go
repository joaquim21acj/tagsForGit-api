package models

type GitRepositories struct {
	Data struct {
		User struct {
			ID                  string `bson:"_id_user" json:"id_user"`
			StarredRepositories struct {
				Edges []struct {
					node Node `bson:"node" json:"node"`
				} `bson:"edges"json:"edges"`
			} `bson:"starredRepositories" json:"starredRepositories"`
		} `bson:"user" json:"user"`
	} `bson:"data" json:"data"`
}
