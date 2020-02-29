package models

// GetRepositories recebe o login do usuário e retorna a query para o graphiql
func GetRepositories(userLogin string) string {
	var query = `	query {
					user(login: "` + userLogin + `") {
					id
					login
					starredRepositories(first: 100) {
						edges {
						node {
							id
							description
							languages(first: 10) {
							edges {
								node {
								name
								}
							}
							}
							name
							projectsUrl
						}
						}
					}
					}
				}
				`

	return query
}

// func GetRepositoriesString () string{
// 	return `{"query":
// 	"{\n    user(login: \"joaquim21acj\")
// 	 {\n      id\n       starredRepositories(first: 10)
// 		 {\n        edges {\n          node {\n            id\n
// 					   description\n            languages(first: 10) {\n
// 								   edges {\n                node {\n
// 											 name\n                }\n
// 													}\n            }\n
// 														 name\n            projectsUrl\n
// 															   }\n
// 																   }\n
// 																    }\n    }\n  }"}'`
// }
