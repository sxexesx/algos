package array_permutation

/*
Given a zero-based permutation nums (0-indexed), build an array ans of the same length where ans[i] = nums[nums[i]]
for each 0 <= i < nums.length and return it.
A zero-based permutation nums is an array of distinct integers from 0 to nums.length - 1 (inclusive).

Constraints:
1 <= nums.length <= 1000
0 <= nums[i] < nums.length
The elements in nums are distinct.
*/
func buildArray(nums []int) []int {
	res := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		m := nums[i]
		res = append(res, nums[m])
	}
	return res
}

// Space complexity: O(1)
func buildArray2(nums []int) []int {
	for i, n := range nums {
		nums[i] += (nums[n] % 1000) * 1000
	}

	for i, _ := range nums {
		nums[i] = nums[i] / 1000
	}

	return nums
}
