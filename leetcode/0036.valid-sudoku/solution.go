package leetcode

func isValidSudoku(board [][]byte) bool {
	colBuckets := make([][]bool, 9)
	boxBuckets := make([][]bool, 9)

	for idx, _ := range colBuckets {
		colBuckets[idx] = make([]bool, 9)
		boxBuckets[idx] = make([]bool, 9)
	}

	for i, row := range board {
		rowBucket := make([]bool, 9)
		boxIndexOffset := (i / 3) * 3

		for j, digit := range row {
			if digit == '.' {
				continue
			}

			num := digit - '1'

			if rowBucket[num] == true {
				return false
			}
			if colBuckets[j][num] == true {
				return false
			}

			boxIndex := boxIndexOffset + (j / 3)
			if boxBuckets[boxIndex][num] == true {
				return false
			}

			rowBucket[num] = true
			colBuckets[j][num] = true
			boxBuckets[boxIndex][num] = true
		}
	}

	return true
}
