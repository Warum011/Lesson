package lru

import "testing"

func TestLRUCacheBasic(t *testing.T) {

	cache := NewLRUCache[int, string](2)

	cache.Put(1, "A")

	if val, ok := cache.Get(1); !ok || val != "A" {
		t.Errorf("expected to get \"A\" for key 1, we got %v (ok=%v)", val, ok)
	}

	cache.Put(2, "B")

	if val, ok := cache.Get(2); !ok || val != "B" {
		t.Errorf("expected to receive \"B\" for key 2, received %v (ok=%v)", val, ok)
	}

	if _, ok := cache.Get(1); !ok {
		t.Error("expected key 1 to be present")
	}

	cache.Put(3, "C")

	if _, ok := cache.Get(2); ok {
		t.Error("expected key 2 to be deleted")
	}

	if val, ok := cache.Get(1); !ok || val != "A" {
		t.Errorf("expected to get \"A\" for key 1, we got %v (ok=%v)", val, ok)
	}
	if val, ok := cache.Get(3); !ok || val != "C" {
		t.Errorf("expected to get \"C\" for key 3, we got %v (ok=%v)", val, ok)
	}
}

func TestLRUCacheUpdate(t *testing.T) {
	cache := NewLRUCache[int, int](2)

	cache.Put(1, 100)
	cache.Put(2, 200)
	cache.Put(1, 101)
	if val, ok := cache.Get(1); !ok || val != 101 {
		t.Errorf("expected an updated value of 101 for key 1, and received %v (ok=%v)", val, ok)
	}

	cache.Put(3, 300)
	if _, ok := cache.Get(2); ok {
		t.Error("expected key 2 to be deleted after inserting the key 3")
	}
	if val, ok := cache.Get(3); !ok || val != 300 {
		t.Errorf("expected a value of 300 for key 3, we got %v (ok=%v)", val, ok)
	}
}
