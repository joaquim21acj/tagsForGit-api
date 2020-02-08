package models

func GetRepositories(userLogin string) string {
	var query = `query:
				{
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
