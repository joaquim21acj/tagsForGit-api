package models

type Tags struct {
	Data struct {
		User struct {
			ID                  string `json:"id"`
			StarredRepositories struct {
				Edges []struct {
					Node struct {
						ID          string `json:"id"`
						Description string `json:"description"`
						Languages   struct {
							Edges []struct {
								Node struct {
									Name string `json:"name"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"languages"`
						Name        string `json:"name"`
						ProjectsURL string `json:"projectsUrl"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"starredRepositories"`
		} `json:"user"`
	} `json:"data"`
}
