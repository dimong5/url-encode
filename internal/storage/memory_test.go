package storage

import "testing"

func TestURLStore(t *testing.T) {
    store := NewURLStore()
    
    store.Save("abc123", "https://google.com")
    url, err := store.Get("abc123")
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    if url != "https://google.com" {
        t.Errorf("expected https://google.com, got %s", url)
    }
}