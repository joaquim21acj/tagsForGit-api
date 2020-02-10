package models

type Tags struct {
	Data struct {
		User struct {
			ID                  string `bson:"id_user" json:"id"`
			StarredRepositories struct {
				Edges []struct {
					Node struct {
						ID          string `bson: "id_repository" json:"id"`
						Description string `bson: "description" json:"description"`
						Languages   struct {
							Edges []struct {
								Node struct {
									Name string `bson:"name_language" json:"name"`
								} `bson:"node_language" json:"node"`
							} `bson:"edges_language" json:"edges"`
						} `bson:"languages" json:"languages"`
						Name        string `bson:"name_repository" json:"name"`
						ProjectsURL string `bson:"projectsURL" json:"projectsUrl"`
					} `bson:"node_repository" json:"node"`
				} `bson:"edges_repository" json:"edges"`
			} `bson:"starredRepositories" json:"starredRepositories"`
		} `bson:"user" json:"user"`
	} `bson:"data" json:"data"`
}
