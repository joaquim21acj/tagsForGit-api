package models

func getRepositories(userLogin string) string {
	var query = `
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
