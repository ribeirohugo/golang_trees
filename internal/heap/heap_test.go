package heap

import (
	"testing"
)

func TestHeap_IsEmpty(t *testing.T) {
	heapTest := Heap{}

	result := heapTest.IsEmpty()
	if !result {
		t.Errorf("Wrong IsEmpty return. Got: %v; Want: %v.", result, true)
	}

	heapTest.Insert(5)
	result = heapTest.IsEmpty()
	if result {
		t.Errorf("Wrong IsEmpty return. Got: %v; Want: %v.", result, false)
	}
}

func TestHeap_Size(t *testing.T) {
	heapTest := Heap{}

	size := heapTest.Size()
	expected := 0
	if size != expected {
		t.Errorf("Wrong Heap Size method return. Got: %d; Want: %d.", size, expected)
	}

	heapTest.Insert(5)
	size = heapTest.Size()
	expected = 1
	if size != expected {
		t.Errorf("Wrong Heap Size method return. Got: %d; Want: %d.", size, expected)
	}

	heapTest.Insert(7)
	heapTest.Insert(9)
	size = heapTest.Size()
	expected = 3
	if size != expected {
		t.Errorf("Wrong Heap Size method return. Got: %d; Want: %d.", size, expected)
	}

	min := heapTest.RemoveMin()
	size = heapTest.Size()
	expected = 2
	if size != expected {
		t.Errorf("Wrong Heap Size method return. Got: %d; Want: %d.", size, expected)
	}

	expected = 5
	if min != expected {
		t.Errorf("Wrong RemoveMin method return. Got: %d; Want: %d.", size, expected)
	}
}

func TestHeap_RecountSize(t *testing.T) {
	heapTest := Heap{}

	size := heapTest.RecountSize()
	expected := 0
	if size != expected {
		t.Errorf("Wrong Heap Size method return. Got: %d; Want: %d.", size, expected)
	}

	heapTest.Insert(5)
	size = heapTest.RecountSize()
	expected = 1
	if size != expected {
		t.Errorf("Wrong Heap Size method return. Got: %d; Want: %d.", size, expected)
	}

	heapTest.Insert(7)
	heapTest.Insert(9)
	size = heapTest.RecountSize()
	expected = 3
	if size != expected {
		t.Errorf("Wrong Heap Size method return. Got: %d; Want: %d.", size, expected)
	}

	min := heapTest.RemoveMin()
	size = heapTest.RecountSize()
	expected = 2
	if size != expected {
		t.Errorf("Wrong Heap Size method return. Got: %d; Want: %d.", size, expected)
	}

	expected = 5
	if min != expected {
		t.Errorf("Wrong RemoveMin method return. Got: %d; Want: %d.", size, expected)
	}
}

func TestHeap_RemoveMin(t *testing.T) {
	heapTest := Heap{}

	value := heapTest.RemoveMin()
	if value != nil {
		t.Errorf("Wrong Heap Size method return. It should return nil.")
	}

	heapTest.Insert(9)
	size := heapTest.Size()
	expected := 1
	if size != expected {
		t.Errorf("Wrong Heap Size method return. Got: %d; Want: %d.", size, expected)
	}

	min := heapTest.RemoveMin()
	size = heapTest.Size()
	expected = 0
	if size != expected {
		t.Errorf("Wrong Heap Size method return. Got: %d; Want: %d.", size, expected)
	}

	expected = 9
	if min != expected {
		t.Errorf("Wrong RemoveMin method return. Got: %d; Want: %d.", size, expected)
	}

	if heapTest.root != nil {
		t.Errorf("RemoveMin didn't removed last element. Got: %v; Want: %v.", heapTest.root, nil)
	}

	if !heapTest.IsEmpty() {
		t.Errorf("RemoveMin didn't removed last element. IsEmpty method should return true. Got: %v; Want: %v.", heapTest.IsEmpty(), true)
	}
}
