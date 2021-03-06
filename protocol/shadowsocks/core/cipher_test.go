package core

import (
	"testing"
)

func TestCipher(t *testing.T) {
	var set = map[string]string{
		"AES-256-GCM":            "Test1234",
		"CHACHA20-IETF-POLY1305": "Test1234",
		"DUMMY":                  "Test1234",
		"AEAD_AES_256_GCM":       "Test1234",
		"AEAD_CHACHA20_POLY1305": "Test1234",
	}

	for method, password := range set {
		_, err := NewCipher(method, password)
		if err != nil {
			t.Errorf("NewCipher Error: %v, Method: %v, Password: %v", err, method, password)
		}
	}
}
