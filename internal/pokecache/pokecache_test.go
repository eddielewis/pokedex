package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	interval := time.Second
	cache := NewCache(interval)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	interval := time.Second
	cache := NewCache(interval)

	cache.Add("key1", []byte("val1"))
	actual, ok := cache.Get("key1")
	if !ok {
		t.Error("key1 not found")
	}
	if string(actual) != "val1" {
		t.Error("value doesn't match")
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Microsecond)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s should have been reaped", keyOne)
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval / 2)

	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("%s should not have been reaped", keyOne)
	}
}
