package service

import (
	"testing"
)

func TestGenerateShortLink(t *testing.T) {
	link := GenerateShortLink("https://google.com")
	if len(link) != 10 {
		t.Errorf("expected length 10, got %d", len(link))
	}
}

func TestGenerateShortLinkSameURL(t *testing.T) {
	link1 := GenerateShortLink("https://test.ru")
	link2 := GenerateShortLink("https://test.ru")
	
	if link1 != link2 {
		t.Errorf("same URL links not equal")
	}
}

func TestGenerateShortLinkValidChars(t *testing.T) {
	link := GenerateShortLink("https://example.com")
	
	for i := 0; i < len(link); i++ {
		c := link[i]
		if !isValidChar(c) {
			t.Errorf("invalid char %c in link %s", c, link)
		}
	}
}

func isValidChar(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		c == '_'
}
