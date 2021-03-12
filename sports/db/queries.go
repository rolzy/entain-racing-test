package db

const (
	sportsList = "list"
)

func getSportQueries() map[string]string {
	return map[string]string{
		sportsList: `
			SELECT 
				id, 
				type, 
				name, 
				advertised_start_time
			FROM sports
		`,
	}
}
