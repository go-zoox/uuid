package uuid

import "testing"

func TestUUID(t *testing.T) {
	id := UUID()
	if id == "" {
		t.Fatal("UUID() returned empty string")
	}

	if len(id) != 36 {
		t.Fatalf("UUID() returned %d characters, expected 36", len(id))
	}
}

func TestV4(t *testing.T) {
	id := V4()
	if id == "" {
		t.Fatal("UUID() returned empty string")
	}

	if len(id) != 36 {
		t.Fatalf("UUID() returned %d characters, expected 36", len(id))
	}
}

func TestV5(t *testing.T) {
	id := V5()
	if id == "" {
		t.Fatal("UUID() returned empty string")
	}

	if len(id) != 36 {
		t.Fatalf("UUID() returned %d characters, expected 36", len(id))
	}
}

func TestV1(t *testing.T) {
	id := V1()
	if id == "" {
		t.Fatal("UUID() returned empty string")
	}

	if len(id) != 36 {
		t.Fatalf("UUID() returned %d characters, expected 36", len(id))
	}
}

func TestV2(t *testing.T) {
	id := V2()
	if id == "" {
		t.Fatal("UUID() returned empty string")
	}

	if len(id) != 36 {
		t.Fatalf("UUID() returned %d characters, expected 36", len(id))
	}
}

func TestV3(t *testing.T) {
	id := V3()
	if id == "" {
		t.Fatal("UUID() returned empty string")
	}

	if len(id) != 36 {
		t.Fatalf("UUID() returned %d characters, expected 36", len(id))
	}
}
