package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0

	for idx, val := range flowerbed {
		if val == 1 {
			continue
		}

		if idx > 0 && flowerbed[idx-1] == 1 {
			continue
		}

		if idx < len(flowerbed)-1 && flowerbed[idx+1] == 1 {
			continue
		}

		count++

		if count >= n {
			return true
		}

		flowerbed[idx] = 1
	}

	return count >= n
}
