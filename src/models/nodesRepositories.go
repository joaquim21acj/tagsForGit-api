package models

type Node struct {
	ID          string `json:"id_project"`
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
	tags        []Tags `json:"tags"`
}
