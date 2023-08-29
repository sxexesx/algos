package concat_arrays

// Task: https://leetcode.com/problems/concatenation-of-array/

// Given an integer array nums of length n, you want to create an array ans of length 2n where ans[i] == nums[i]
// and ans[i + n] == nums[i] for 0 <= i < n (0-indexed).
// Specifically, ans is the concatenation of two nums arrays.
// Return the array ans.

// Constraints:
// n == nums.length
// 1 <= n <= 1000
// 1 <= nums[i] <= 1000
func getConcatenation(nums []int) []int {
	res := make([]int, 0, 2*len(nums))
	for _, num := range nums {
		res = append(res, num)
	}
	for _, num := range nums {
		res = append(res, num)
	}
	return res
}
