package jewels_and_stones

import "testing"

// Example 1:
// Input: jewels = "aA", stones = "aAAbbbb"
// Output: 3
//
// Example 2:
//
// Input: jewels = "z", stones = "ZZ"
// Output: 0
func Test_1(t *testing.T) {
	if r := numJewelsInStones("aA", "aAAbbbb"); r != 3 {
		t.Fatalf("wrong answer! should be 3, but was %d", r)
	}
}

func Test_2(t *testing.T) {
	if r := numJewelsInStones("z", "ZZ"); r != 0 {
		t.Fatalf("wrong answer! should be 0, but was %d", r)
	}
}

func Test_3(t *testing.T) {
	if r := numJewelsInStones("z", "ZslrnvowppvwivzbksjbvkjsrnlknslZ"); r != 1 {
		t.Fatalf("wrong answer! should be 1, but was %d", r)
	}
}
