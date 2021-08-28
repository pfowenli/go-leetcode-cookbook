package leetcode

func duplicateZeros(arr []int) {
	left := 0
	right := len(arr) - 1

	for left < right {
		if arr[left] == 0 {
			right--
		}
		left++
	}

	if right == left {
		arr[len(arr)-1] = arr[left]
		left--
		right = len(arr) - 2
	} else {
		left = right
		right = len(arr) - 1
	}

	for left < right {
		if arr[left] == 0 {
			arr[right] = 0
			right--
		}

		arr[right] = arr[left]
		right--
		left--
	}
}
