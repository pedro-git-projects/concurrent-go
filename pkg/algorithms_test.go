package pkg

import "testing"

func TestContains(t *testing.T) {
	s0 := []int{1, 2, 3, 4, 5, 6}

	c0 := Contains(s0, 1)
	if !c0 {
		t.Errorf("Expected %t Got %t", true, c0)
	}

	c1 := Contains(s0, -1)
	if c1 {
		t.Errorf("Expected %t Got %t", false, c1)
	}

	s1 := [][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}
	s2 := []int{1, 2, 3, 4, 5}
	s3 := []int{1, 2, 3, 4, 6}

	c2 := Contains(s1, s2)
	if !c2 {
		t.Errorf("Expected %t Got %t", true, c2)
	}

	c3 := Contains(s1, s3)
	if c3 {
		t.Errorf("Expected %t Got %t", false, c3)
	}

}
