package leetcode

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	left := 0
	for ; left < len(arr)-1; left++ {
		if arr[left] == arr[left+1] {
			return false
		}
		if arr[left] > arr[left+1] {
			break
		}
	}

	if left == 0 || left == len(arr)-1 {
		return false
	}

	right := len(arr) - 1
	for ; right > 0; right-- {
		if arr[right] == arr[right-1] {
			return false
		}
		if arr[right] > arr[right-1] {
			break
		}
	}

	return left == right
}
