package when_this

import "testing"

func TestIsEmptyString(t *testing.T) {
	if v := IsEmptyString("").UseThisString("ABC"); v != "ABC" {
		t.Error("Expect", "ABC", v)
	}
}
