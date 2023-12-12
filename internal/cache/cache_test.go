package cache

import (
	"slices"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	data := []struct {
		key   string
		value []byte
	}{
		{
			key:   "https://example.com",
			value: []byte("testdata"),
		},
		{
			key:   "https://example.com/path",
			value: []byte("moretestdata"),
		},
	}

	const interval = 10 * time.Millisecond
	cache := NewCache(interval)

	for _, d := range data {
		cache.Add(d.key, d.value)

		v, ok := cache.Get(d.key)
		if !ok {
			t.Fatal("expected to find key", d.key)
		}
		if !slices.Equal(v, d.value) {
			t.Fatalf("\ngot: %v\nwant: %v", v, d.value)
		}
	}

	time.Sleep(interval + time.Millisecond)
	for _, d := range data {
		v, ok := cache.Get(d.key)
		if ok {
			t.Log(v)
			t.Fatalf("key %q should have been reaped", d.key)
		}
	}
}
