package tree

import "testing"

func TestTree(t *testing.T) {
	root := 20
	left := 19
	right := 21

	testTree := Tree{
		root: &Node{element: root},
	}

	testTree.Insert(left)
	testTree.Insert(right)

	rootElement := testTree.Root().GetElement()
	if rootElement != root {
		t.Errorf("Wrong root element. Got: %d; Want: %d.", rootElement, root)
	}

	leftElement := testTree.Root().GetLeft().GetElement()
	if leftElement != left {
		t.Errorf("Wrong left element. Got: %d; Want: %d.", leftElement, left)
	}

	rightElement := testTree.Root().GetRight().GetElement()
	if rightElement != right {
		t.Errorf("Wrong right element. Got: %d; Want: %d.", rightElement, right)
	}
}

func TestTree_IsEmpty(t *testing.T) {
	testTree := Tree{}

	if !testTree.IsEmpty() {
		t.Errorf("Wrong IsEmpty return value. Got: %t; Want: %t.", testTree.IsEmpty(), true)
	}
}

func TestTree_Smallest(t *testing.T) {
	testTree := Tree{}
	if testTree.Smallest() != nil {
		t.Errorf("Wrong smallest element return. It should return nil.")
	}

	testTree.Insert(20)
	testTree.Insert(5)
	testTree.Insert(15)
	testTree.Insert(17)
	testTree.Insert(19)

	expectedSmallest := 5
	smallestResult := testTree.Smallest().GetElement()
	if expectedSmallest != smallestResult {
		t.Errorf("Wrong smallest element return. Got: %d; Want: %d.", smallestResult, expectedSmallest)
	}
}

func TestTree_Find(t *testing.T) {
	testTree := Tree{}
	if testTree.Find(0) != nil {
		t.Errorf("Wrong find element return. It should return nil.")
	}

	testTree.Insert(20)
	testTree.Insert(5)
	testTree.Insert(15)
	testTree.Insert(17)
	testTree.Insert(19)

	findResult := testTree.Find(17)
	if findResult == nil {
		t.Errorf("Wrong find element return. It shouldn't return nil.")
	}

	findResult = testTree.Find(22)
	if findResult != nil {
		t.Errorf("Wrong find element return. It shouldn't return nil.")
	}
}

func TestTree_Height(t *testing.T) {
	testTree := Tree{}

	expected := -1
	heightResult := testTree.Height()
	if heightResult != expected {
		t.Errorf("Wrong height element return. Got: %d; Want: %d.", heightResult, expected)
	}

	testTree.Insert(20)
	testTree.Insert(5)
	testTree.Insert(15)

	expected = 2
	heightResult = testTree.Height()
	if heightResult != expected {
		t.Errorf("Wrong height element return. Got: %d; Want: %d.", heightResult, expected)
	}

	testTree.Insert(17)

	expected = 3
	heightResult = testTree.Height()
	if heightResult != expected {
		t.Errorf("Wrong height element return. Got: %d; Want: %d.", heightResult, expected)
	}

	testTree.Insert(18)

	expected = 4
	heightResult = testTree.Height()
	if heightResult != expected {
		t.Errorf("Wrong height element return. Got: %d; Want: %d.", heightResult, expected)
	}
}

func TestTree_Remove(t *testing.T) {
	testTree := Tree{}

	testTree.Insert(20)
	testTree.Insert(5)
	testTree.Insert(15)
	testTree.Insert(17)
	testTree.Insert(19)

	testTree.Remove(19)
	height := testTree.Height()
	if height == 4 {
		t.Errorf("Wrong height after removing element. Got: %d; Want: %d.", height, 4)
	}

	node := testTree.Find(19)
	if node != nil {
		t.Errorf("Wrong node returned after removing element. It shoud be nil.")
	}

	testTree.Remove(17)
	height = testTree.Height()
	if height == 3 {
		t.Errorf("Wrong height after removing element. Got: %d; Want: %d.", height, 3)
	}

	node = testTree.Find(17)
	if node != nil {
		t.Errorf("Wrong node returned after removing element. It shoud be nil.")
	}
}
