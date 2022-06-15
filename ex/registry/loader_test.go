package registry

import (
	"strings"
	"testing"
)

func TestAddressMutation(t *testing.T) {
	addr := &Area{
		ID:       "TEST",
		Locality: "South WoodFord",
		Town:     "London",
		District: "London",
		County:   "London",
	}

	mutation := AreaMutation(addr)
	if mutation == "" {
		t.Fatalf("empty address mutation")
	}

	if !strings.Contains(mutation, "TEST") {
		t.Fatalf("Id missing from mutaiotn")
	}

	t.Log(mutation)
}
