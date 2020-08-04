package dushengchen

/**
Submission:
    https://leetcode.com/submissions/detail/370911144/
*/
func calculateMinimumHP(dungeon [][]int) int {

	for i := len(dungeon) - 1; i >= 0; i-- {
		for j := len(dungeon[i]) - 1; j >= 0; j-- {
			if i == len(dungeon)-1 && j == len(dungeon[i])-1 {
				dungeon[i][j] = calMinLimit(dungeon[i][j], 1)
				continue
			} else if i == len(dungeon)-1 {
				dungeon[i][j] = calMinLimit(dungeon[i][j], dungeon[i][j+1])
				continue
			} else if j == len(dungeon[i])-1 {
				dungeon[i][j] = calMinLimit(dungeon[i][j], dungeon[i+1][j])
				continue
			}
			a := calMinLimit(dungeon[i][j], dungeon[i+1][j])
			b := calMinLimit(dungeon[i][j], dungeon[i][j+1])
			if a > b {
				dungeon[i][j] = b
			} else {
				dungeon[i][j] = a
			}
		}
	}
	return dungeon[0][0]
}

func calMinLimit(cost int, limit int) int {
	if rest := limit - cost; rest <= 0 {
		return 1
	}
	return rest
}
