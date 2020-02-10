package models

func GetRepositories(userLogin string) string {
	var query = `	{
					user(login: "` + userLogin + `") {
					id
					starredRepositories(first: 10) {
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
