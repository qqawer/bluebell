package snowflake

import (

	"testing"
)

func Test(t *testing.T) {
	err := Init("2024-12-13", 1)
	if err != nil {
		t.Errorf("Init failed: %v", err)
	} else {
		output := GenID()
		t.Logf("Init number: %d\n", output)

	}
}
