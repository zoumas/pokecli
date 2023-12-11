package cache

import (
	"slices"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	const interval = time.Second
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

	cache := NewCache(interval)

	for _, d := range data {
		cache.Add(d.key, d.value)
		v, ok := cache.Get(d.key)
		if !ok {
			t.Fatal("expected to find key")
		}
		if !slices.Equal(v, d.value) {
			t.Fatal("expected to find value")
		}
	}

	time.Sleep(interval)
	for _, d := range data {
		v, ok := cache.Get(d.key)
		if ok {
			t.Log(v)
			t.Fatal("key should have been reaped")
		}
	}
}
