package jewels_and_stones

// Task: https://leetcode.com/problems/jewels-and-stones/

// You're given strings jewels representing the types of stones that are jewels, and stones representing the stones you have.
// Each character in stones is a type of stone you have. You want to know how many of the stones you have are also jewels.
//
// Letters are case sensitive, so "a" is considered a different type of stone from "A".
//
// Constraints:
// 1 <= jewels.length, stones.length <= 50
// jewels and stones consist of only English letters.
// All the characters of jewels are unique.
func numJewelsInStones(jewels string, stones string) int {
	jewelsArr := []rune(jewels)
	stonesArr := []rune(stones)
	jmap := make(map[rune]bool)
	for _, j := range jewelsArr {
		jmap[j] = true
	}
	var result int
	for _, s := range stonesArr {
		if jmap[s] {
			result++
		}
	}
	return result
}
