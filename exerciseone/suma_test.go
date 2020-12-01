package exerciseone

import "testing"

func Test_Add(t *testing.T) {
	want := 5
	got := Add(2, 3)

	//The result have to be five
	if got != want {
		t.Logf("ERROR: Se esperaba %d, se obtuvo %d", want, got)
		t.Fail()
	}
}

func Test_AddMultiple(t *testing.T) {
	want := 6
	got := AddMultiple(1, 2, 3)

	//The result have to be five
	if got != want {
		t.Logf("ERROR: Se esperaba %d, se obtuvo %d", want, got)
		t.Fail()
	}
}
