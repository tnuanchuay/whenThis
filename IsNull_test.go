package whenThis

import "testing"

func TestIsNullDoThis(t *testing.T) {
	variable := "someValues"

	IsNull(nil).
		DoThis(func() {
			variable = "changed"
		})

	if variable != "changed" {
		t.Error("expect", "changed", "actual", variable)
	}
}
