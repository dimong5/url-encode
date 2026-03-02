package service

import "testing"

//тест на детермированность и длину короткой ссылки

func TestGenerateShortLink(t *testing.T) {

	link1 := GenerateShortLink("https://ozon.ru")
	link2 := GenerateShortLink("https://ozon.ru")
	
	if link1 != link2 {
		t.Errorf(": %s != %s", link1, link2)
	}
	
	if len(link1) != 10 {
		t.Errorf("expected length 10, got %d", len(link1))
	}
}