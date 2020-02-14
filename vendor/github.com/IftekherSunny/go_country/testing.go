package country

import "testing"

func assert(t *testing.T, expect interface{}, actual interface{}, message string) {
	if expect != actual {
		t.Error(message)
	}
}
