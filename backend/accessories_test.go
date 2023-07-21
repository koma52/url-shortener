package backend

import "testing"

func TestMakeShortCode(t *testing.T) {
	h := makeShortcode("http://example.com")

	if h != "a9b9f04336ce0181a08e774e01113b31" {
		t.Errorf("Expected hash is a9b9f04336ce0181a08e774e01113b31, but got %v", h)
	}
}
