package main

import (
	"strings"
	"testing"
)

func TestUuidV4(t *testing.T) {
	testUuidV4(t, 10)
	testUuidV4(t, 100)
	testUuidV4(t, 1000)
	testUuidV4(t, 10000)
}

func testUuidV4(t *testing.T, n int) {
	set := make(map[string]bool)
	var u string
	for i := 0; i < n; i++ {
		u = string(uuidV4())
		//log.Println(u)
		if !validV4(u) {
			t.Errorf("value '%s' not a valid UUIDv4", u)
		}
		if set[u] {
			t.Errorf("value '%s' has already been generated", u)
		}
		set[u] = true
	}
}

func BenchmarkUuidV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuidV4()
	}
}

func validV4(s string) bool { // based on github.com/rogpeppe/fastuuid
	if len(s) != 36 {
		return false
	}
	parts := strings.Split(s, "-")
	if len(parts) != 5 {
		return false
	}
	if len(parts[0]) != 8 ||
		len(parts[1]) != 4 || len(parts[2]) != 4 ||
		len(parts[3]) != 4 || len(parts[4]) != 12 {
		return false
	}
	for _, p := range parts {
		if !isHex(p) {
			return false
		}
	}
	return true
}

func isHex(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !('0' <= c && c <= '9' || 'a' <= c && c <= 'f') {
			return false
		}
	}
	return true
}
