package lfqueue

import (
	"testing"
)

func TestQueueBasic(t *testing.T) {
	var i int
	capacity := 100
	q := NewQueue(capacity)
	for i = 0; i < q.Capacity(); i++ {
		ok := q.Put(i)
		if !ok {
			t.Fatalf("put %d returns false", i)
		}
	}

	// Full queue
	ok := q.Put(i)
	if ok {
		t.Fatalf("put %d returns true", i)
	}

	for i = 0; i < capacity; i++ {
		elem, ok := q.Get()
		if !ok {
			t.Fatalf("get %d returns false", i)
		}
		if elem.(int) != i {
			t.Fatalf("Get wrong value [%d] = %d", i, elem.(int))
		}
	}
}

func TestDroppingQueueBasic(t *testing.T) {
	var i int
	capacity := 100
	q := NewDroppingQueue(capacity)
	for i = 0; i < q.Capacity(); i++ {
		ok := q.Put(i)
		if !ok {
			t.Fatalf("put %d returns false", i)
		}
	}

	// Full queue - drop last element
	ok := q.Put(i)
	if !ok {
		t.Fatalf("put %d returns false", i)
	}

	for i = 0; i < capacity; i++ {
		elem, ok := q.Get()
		if !ok {
			t.Fatalf("get %d returns false", i)
		}
		// queue overfloed, so first element must dropped
		if elem.(int) != i+1 {
			t.Fatalf("Get wrong value [%d] = %d", i, elem.(int))
		}
	}
}
