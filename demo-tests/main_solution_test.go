package demo_tests

import "testing"

func TestDivide2By2(t *testing.T) {
	r := divide2By(2)
	if r != 1 {
		t.Errorf("An error occurred : 2/%d shouldn't be %d", -43, r)
	}
}

func TestDivide2ByLessThan0(t *testing.T) {
	r := divide2By(-43)
	if r != 0 {
		t.Errorf("An error occurred : 2/%d shouldn't be %d", 2, r)
	}
}

func TestDivide2By0(t *testing.T) {
	r := divide2By(0)
	if r != 9999999 {
		t.Errorf("An error occurred : 2/%d shouldn't be %d", 0, r)
	}
}
