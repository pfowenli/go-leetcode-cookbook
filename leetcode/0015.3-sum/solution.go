/**

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
Notice that the solution set must not contain duplicate triplets.

Example 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]

Example 2:
Input: nums = []
Output: []

Example 3:
Input: nums = [0]
Output: []

Input
[-1,0,1,2,-1,-4,-2,-3,3,0,4]
Output
[[1,-1,0],[1,-1,0],[2,-3,1],[2,-2,0],[3,-4,1],[3,-3,0],[3,-2,-1],[4,-4,0],[4,-3,-1]]
Expected
[[-4,0,4],[-4,1,3],[-3,-1,4],[-3,0,3],[-3,1,2],[-2,-1,3],[-2,0,2],[-1,-1,2],[-1,0,1]]


*/

func main () {
	nums = [-1,0,1,2,-1,-4]
	results = [[-1,-1,2],[-1,0,1]]
}


func threeSum(nums []int) [][]int {
	sort.Slice(nums, func (a, b int) bool {
		return nums[a] < nums[b]
	})

	n := len(nums)
	results := [][]int{}

	for i := 2; i < n; i++ {
		for i < n - 1 && nums[i] == nums[i + 1] {
			i++
		}

		j := 0
		k := i - 1

		for j < k {

			for j < k - 1 && nums[j] == nums[j + 1] {
				j++
			}

			for k > j + 1 && nums[k] == nums[k - 1] {
				k--
			}

			if nums[i] + nums[j] + nums[k] == 0 {
				results = append(results, []int{nums[i], nums[j], nums[k]})
			}

		}
	}

	return results
}