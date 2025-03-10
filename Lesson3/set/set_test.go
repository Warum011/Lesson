package set

import (
	"testing"
)

func TestAddContains(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	if !s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
		t.Errorf("set was expected to contain 1, 2, 3, got: %v", s)
	}

	if s.Contains(4) {
		t.Errorf("expected that the set does not contain 4, got: %v", s)
	}
}

func TestRemove(t *testing.T) {
	s := NewSet[int](1, 2, 3)
	s.Remove(2)

	if s.Contains(2) {
		t.Errorf("after removing 2, the set should not contain 2, got: %v", s)
	}
}

func TestUnion(t *testing.T) {
	s1 := NewSet[int](1, 2, 3)
	s2 := NewSet[int](3, 4, 5)
	union := s1.Union(s2)

	expectedElements := []int{1, 2, 3, 4, 5}
	for _, v := range expectedElements {
		if !union.Contains(v) {
			t.Errorf("union does not contain the expected element %d", v)
		}
	}
	if len(union) != 5 {
		t.Errorf("expected union length = 5, received %d", len(union))
	}
}

func TestIntersection(t *testing.T) {
	s1 := NewSet[int](1, 2, 3, 4)
	s2 := NewSet[int](3, 4, 5, 6)
	intersection := s1.Intersection(s2)

	expectedElements := []int{3, 4}
	for _, v := range expectedElements {
		if !intersection.Contains(v) {
			t.Errorf("intersection must contain the element %d", v)
		}
	}
	if len(intersection) != 2 {
		t.Errorf("expected intersection length = 2, received%d", len(intersection))
	}
}

func TestDifference(t *testing.T) {
	s1 := NewSet[int](1, 2, 3, 4)
	s2 := NewSet[int](3, 4, 5, 6)
	difference := s1.Difference(s2)

	expectedElements := []int{1, 2}
	for _, v := range expectedElements {
		if !difference.Contains(v) {
			t.Errorf("difference must contain the element %d", v)
		}
	}
	if len(difference) != 2 {
		t.Errorf("expected difference length = 2, received %d", len(difference))
	}
}
