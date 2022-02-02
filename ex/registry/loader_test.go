package registry

import (
	"strings"
	"testing"
)

func TestAddressMutation(t *testing.T) {
	addr := &Address{
		ID:       "TEST",
		PAON:     "12",
		SAON:     "A",
		Street:   "Dewberry",
		Locality: "South WoodFord",
		Town:     "London",
		District: "London",
		County:   "London",
	}

	mutation := AddressMutation(addr)
	if mutation == "" {
		t.Fatalf("empty address mutation")
	}

	if !strings.Contains(mutation, "TEST") {
		t.Fatalf("Id missing from mutaiotn")
	}

	t.Log(mutation)
}
