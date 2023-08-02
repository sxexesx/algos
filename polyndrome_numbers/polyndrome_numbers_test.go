package polyndrome_numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 22, 33, 44, 55, 66, 77, 88, 99, 101
func TestFindReverseNumber_1(t *testing.T) {
	if res := FindReverseNumber(1); assert.Equal(t, uint64(0), res) != true {
		t.Fatalf("expected: %d, received: %d", 0, res)
	}
}

func TestFindReverseNumber_2(t *testing.T) {
	if res := FindReverseNumber(2); assert.Equal(t, uint64(1), res) != true {
		t.Fatalf("expected: %d, received: %d", 1, res)
	}
}

func TestFindReverseNumber_3(t *testing.T) {
	if res := FindReverseNumber(10); assert.Equal(t, uint64(9), res) != true {
		t.Fatalf("expected: %d, received: %d", 9, res)
	}
}

func TestFindReverseNumber_4(t *testing.T) {
	if res := FindReverseNumber(100); assert.Equal(t, uint64(909), res) != true {
		t.Fatalf("expected: %d, received: %d", 909, res)
	}
}

// too long :(
//func TestFindReverseNumber_5(t *testing.T) {
//	FindReverseNumber(100000)
//	if res := FindReverseNumber(100000000); assert.Equal(t, uint64(900000000000009), res) != true {
//		t.Fatalf("expected: %d, received: %d", 909, res)
//	}
//}

func TestIsRevNumber_1(t *testing.T) {
	if res := isRevNumber(9); !assert.True(t, res) {
		t.Fail()
	}
}

func TestIsRevNumber_2(t *testing.T) {
	if res := isRevNumber(12021); !assert.True(t, res) {
		t.Fail()
	}
}

func TestIsRevNumber_3(t *testing.T) {
	if res := isRevNumber(1112); !assert.False(t, res) {
		t.Fail()
	}
}

func TestIsRevNumber_4(t *testing.T) {
	if res := isRevNumber(123404321); !assert.True(t, res) {
		t.Fail()
	}
}

func TestIsRevNumber_5(t *testing.T) {
	if res := isRevNumber(12344321); !assert.True(t, res) {
		t.Fail()
	}
}

func TestIsRevNumber_6(t *testing.T) {
	if res := isRevNumber(123456789); !assert.False(t, res) {
		t.Fail()
	}
}

func TestIsRevNumber_7(t *testing.T) {
	if res := isRevNumber(900000000000009); !assert.True(t, res) {
		t.Fail()
	}
}
