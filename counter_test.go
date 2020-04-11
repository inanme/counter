package counter

import "testing"

func TestCounter(t *testing.T) {
	if counter := Counter(); counter != 2 {
		t.Errorf("Counter should be one, instead %d", counter)
	}
}
