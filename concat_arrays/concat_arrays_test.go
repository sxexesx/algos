package concat_arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*
Example 1:

Input: nums = [1,2,1]
Output: [1,2,1,1,2,1]
Explanation: The array ans is formed as follows:
- ans = [nums[0],nums[1],nums[2],nums[0],nums[1],nums[2]]
- ans = [1,2,1,1,2,1]

Example 2:
Input: nums = [1,3,2,1]
Output: [1,3,2,1,1,3,2,1]
Explanation: The array ans is formed as follows:
- ans = [nums[0],nums[1],nums[2],nums[3],nums[0],nums[1],nums[2],nums[3]]
- ans = [1,3,2,1,1,3,2,1]
*/
func TestConcatArray_1(t *testing.T) {
	r := getConcatenation([]int{1, 2, 3})
	require.Equal(t, []int{1, 2, 3, 1, 2, 3}, r)
}

func TestConcatArray_2(t *testing.T) {
	r := getConcatenation([]int{1, 3, 2, 1})
	require.Equal(t, []int{1, 3, 2, 1, 1, 3, 2, 1}, r)
}
