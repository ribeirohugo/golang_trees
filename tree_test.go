package golang_trees

import (
	"reflect"
	"testing"
)

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

func TestNode_Nodes(t *testing.T) {
	testTree := Tree{}
	testTree.Insert(20)

	expected := 5
	leftNode := Node{}
	leftNode.SetElement(expected)

	if leftNode.element != expected {
		t.Errorf("Wrong SetElement element. Got: %d; Want: %d.", leftNode.element, expected)
	}

	expected = 21
	rightNode := NewNode(nil, expected, nil)
	if rightNode.element != expected {
		t.Errorf("Wrong NewNode element returned. Got: %d; Want: %d.", rightNode.element, expected)
	}

	testTree.root.SetLeft(leftNode)
	testTree.root.SetRight(rightNode)

	expected = 5
	element := testTree.root.left.element
	if expected != element {
		t.Errorf("Wrong SetLeft element returned. Got: %d; Want: %d.", element, expected)
	}

	expected = 21
	element = testTree.root.right.element
	if expected != element {
		t.Errorf("Wrong SetRight element returned. Got: %d; Want: %d.", element, expected)
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

func TestTree_PreOrder(t *testing.T) {
	testTree := Tree{}

	var expected []int

	if !reflect.DeepEqual(testTree.PreOrder(), expected) {
		t.Errorf("Wrong PreOrder() empty output returned. \nGot: %d; \nWant: %d.", testTree.PreOrder(), expected)
	}

	testTree.Insert(20)
	testTree.Insert(10)
	testTree.Insert(21)
	testTree.Insert(22)
	testTree.Insert(11)
	testTree.Insert(9)

	collection := testTree.PreOrder()

	expectedCollection := []int{20, 10, 9, 11, 21, 22}

	if !reflect.DeepEqual(collection, expectedCollection) {
		t.Errorf("Wrong PreOrder output returned. \nGot: %d; \nWant: %d.", collection, expectedCollection)
	}
}

func TestTree_InOrder(t *testing.T) {
	testTree := Tree{}

	var expected []int

	if !reflect.DeepEqual(testTree.InOrder(), expected) {
		t.Errorf("Wrong InOrder() empty output returned. \nGot: %d; \nWant: %d.", testTree.InOrder(), expected)
	}

	testTree.Insert(20)
	testTree.Insert(10)
	testTree.Insert(21)
	testTree.Insert(22)
	testTree.Insert(11)
	testTree.Insert(9)

	collection := testTree.InOrder()

	expectedCollection := []int{9, 10, 11, 20, 21, 22}

	if !reflect.DeepEqual(collection, expectedCollection) {
		t.Errorf("Wrong InOrder output returned. \nGot: %d; \nWant: %d.", collection, expectedCollection)
	}
}

func TestTree_PosOrder(t *testing.T) {
	testTree := Tree{}

	var expected []int

	if !reflect.DeepEqual(testTree.PosOrder(), expected) {
		t.Errorf("Wrong PosOrder() empty output returned. \nGot: %d; \nWant: %d.", testTree.PosOrder(), expected)
	}

	testTree.Insert(20)
	testTree.Insert(10)
	testTree.Insert(21)
	testTree.Insert(22)
	testTree.Insert(11)
	testTree.Insert(9)

	collection := testTree.PosOrder()
	expectedCollection := []int{9, 11, 10, 22, 21, 20}

	if !reflect.DeepEqual(collection, expectedCollection) {
		t.Errorf("Wrong PosOrder output returned. \nGot: %d; \nWant: %d.", collection, expectedCollection)
	}
}
