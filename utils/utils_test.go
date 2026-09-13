package utils

import (
	"testing"
)

func TestParsePortsSingleAndList(t *testing.T) {
	ports, err := ParsePorts("80,443,8080")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(ports) != 3 || ports[0] != 80 || ports[1] != 443 || ports[2] != 8080 {
		t.Errorf("unexpected parsed ports: %v", ports)
	}
}

func TestParsePortsRange(t *testing.T) {
	ports, err := ParsePorts("20-25")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(ports) != 6 || ports[0] != 20 || ports[5] != 25 {
		t.Errorf("unexpected range ports: %v", ports)
	}
}

func TestParsePortsTop100(t *testing.T) {
	ports, err := ParsePorts("top100")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(ports) == 0 {
		t.Errorf("expected non-empty top100 ports")
	}
}

func TestParsePortsInvalid(t *testing.T) {
	_, err := ParsePorts("abc")
	if err == nil {
		t.Errorf("expected error for non-integer port")
	}

	_, err = ParsePorts("100-50")
	if err == nil {
		t.Errorf("expected error for inverted range")
	}
}
