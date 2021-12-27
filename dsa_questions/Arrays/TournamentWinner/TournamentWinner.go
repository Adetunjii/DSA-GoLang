 package dsaquestions

const HOME_TEAM_WON = 1
const AWAY_TEAM_WON = 0


//Optimal Solution O(N) Time, O(K) Space Time

func TournamentWinner(competitions [][]string, results []int) string{
	teamTable := make(map[string]int, len(results))
	currentHighestValue := 0
	var currentHighestKey string


	for index := range competitions{
		if results[index] == HOME_TEAM_WON {
			winningTeam := competitions[index][0]
			teamTable[winningTeam] += 3
		}
		if results[index] == AWAY_TEAM_WON {
			winningTeam := competitions[index][1]
			teamTable[winningTeam] += 3
		} 
	}


	for key, value := range teamTable {
		if value > currentHighestValue {
			currentHighestValue = value
			currentHighestKey = key
		}
	}

	return currentHighestKey
}