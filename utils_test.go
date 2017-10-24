package main

import "testing"

const remoteAddr = "192.168.0.1:12345"
const expectedIP = "192.168.0.1"

func BenchmarkSplitIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		splitIP(remoteAddr)
	}
}

func BenchmarkLastIndexIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lastIndexIP(remoteAddr)
	}
}

func TestSplitIP(t *testing.T) {
	got := splitIP(remoteAddr)
	if got != expectedIP {
		t.Fatalf("expected %q, got %q", expectedIP, got)
	}
}

func TestLastIndexIP(t *testing.T) {
	got := lastIndexIP(remoteAddr)
	if got != expectedIP {
		t.Fatalf("expected %q, got %q", expectedIP, got)
	}
}
