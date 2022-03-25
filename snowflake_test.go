package uuid

import "testing"

func TestSnowflake(t *testing.T) {
	id := Snowflake(0)
	if id == 0 {
		t.Fatal("Snowflake() returned 0")
	}

	t.Log("snowflake id:", id)
}
