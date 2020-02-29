package models

// Node correspondente aos dados do repositorio
type Node struct {
	ID          string `bson:"id" json:"id_project"`
	Description string `bson:"description" json:"description"`
	Languages   struct {
		Edges []struct {
			Node struct {
				Name string `bson:"name" json:"name"`
			} `bson:"node" json:"node"`
		} `bson:"edges" json:"edges"`
	} `bson:"languages" json:"languages"`
	Name        string `bson:"name" json:"name"`
	ProjectsURL string `bson:"projectsUrl" json:"projectsUrl"`
	Tags        []Tag  `bson:"tags" json:"tags"`
}
