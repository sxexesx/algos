package revers_numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRevNumber_0(t *testing.T) {
	if res := revNumber(2, 1); assert.Equal(t, uint64(2), res) != true {
		t.Fail()
	}
}

func TestRevNumber_1(t *testing.T) {
	if res := revNumber(1002, 4); assert.Equal(t, uint64(2001), res) != true {
		t.Fail()
	}
}

func TestRevNumber_2(t *testing.T) {
	if res := revNumber(123456789, 9); assert.Equal(t, uint64(987654321), res) != true {
		t.Fail()
	}
}

func TestRevNumber_3(t *testing.T) {
	if res := revNumber(123456789123456789, 18); assert.Equal(t, uint64(987654321987654321), res) != true {
		t.Fail()
	}
}

func TestRevNumber_4(t *testing.T) {
	if res := revNumber(1844674407370955161, 19); assert.Equal(t, uint64(1615590737044764481), res) != true {
		t.Fail()
	}
}

func TestRevNumber_5(t *testing.T) {
	if res := revNumber(900, 3); assert.Equal(t, "009", res) != true {
		t.Fail()
	}
}
