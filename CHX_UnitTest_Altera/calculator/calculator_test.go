package calc

import "testing"

func TestAdd(t *testing.T) {
	actual := Add(1, 2)
	expect := 3
	if actual != expect {
		t.Error("Actual is not match with expected")
	}
}

func TestSub(t *testing.T) {
	actual := Sub(2, 1)
	expect := 1
	if actual != expect {
		t.Error("Actual is not match with expected")
	}
}