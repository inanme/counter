package v2

import "testing"

func TestCounter(t *testing.T) {
	if counter := Counter(); counter != 3 {
		t.Errorf("Counter should be one, instead %d", counter)
	}
}
