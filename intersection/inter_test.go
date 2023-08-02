package intersection

import (
	"testing"
)

func TestIntersection_1(t *testing.T) {
	feedTags := []string{"best_items", "top_pv", "adult"}
	offersTags := []string{"best_items", "top_pv", "adult"}
	if res := GetIntersection(feedTags, offersTags); len(res) != 3 {
		t.Fatalf("wrong result: %v", res)
	}
}

func TestIntersection_2(t *testing.T) {
	feedTags := []string{"adult"}
	offersTags := []string{"best_items", "top_pv", "adult"}
	if res := GetIntersection(feedTags, offersTags); len(res) != 1 {
		t.Fatalf("wrong result: %v", res)
	}
}

func TestIntersection_3(t *testing.T) {
	feedTags := []string{"adult"}
	offersTags := []string{"best_items", "top_pv"}
	if res := GetIntersection(feedTags, offersTags); len(res) != 0 {
		t.Fatalf("wrong result: %v", res)
	}
}

func TestIntersection_4(t *testing.T) {
	var feedTags []string
	offersTags := []string{"best_items", "top_pv"}
	if res := GetIntersection(feedTags, offersTags); len(res) != 2 {
		t.Fatalf("wrong result: %v", res)
	}
}

func TestIntersection_5(t *testing.T) {
	feedTags := []string{"best_items", "top_pv"}
	var offersTags []string
	if res := GetIntersection(feedTags, offersTags); len(res) != 0 {
		t.Fatalf("wrong result: %v", res)
	}
}
