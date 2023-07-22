package backend

import "testing"

func TestMakeShortCode(t *testing.T) {
	h := makeShortcode("http://example.com")

	if h != "a9b9f04336ce0181a08e774e01113b31" {
		t.Errorf("Expected hash is a9b9f04336ce0181a08e774e01113b31, but got %v", h)
	}
}

func TestIsURL(t *testing.T) {
	if isURL("http://example.com") != true {
		t.Errorf("Expected output is true, but got %v", isURL("http://example.com"))
	}

	if isURL("http//example.com") != false {
		t.Errorf("Expected output is false, but got %v", isURL("http://example.com"))
	}
}
