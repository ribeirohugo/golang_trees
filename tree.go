package golang_trees

import (
	"strconv"
)

type Tree struct {
	root *Node
}

func NewTree() Tree {
	return Tree{
		root: nil,
	}
}

func (t *Tree) Root() *Node {
	return t.root
}

func (t *Tree) IsEmpty() bool {
	if t.root == nil {
		return true
	}
	return false
}

func (t *Tree) Smallest() *Node {
	if t.root == nil {
		return nil
	}

	return t.root.smallest()
}

func (n *Node) smallest() *Node {
	if n.GetLeft() == nil {
		return n
	}
	return n.GetLeft().smallest()
}

func (t *Tree) Insert(element int) {
	if t.IsEmpty() {
		node := Node{element: element}
		t.root = &node
		return
	}

	t.root.insert(element)
}

func (n *Node) insert(element int) {
	if n.element >= element {
		if n.left == nil {
			n.left = &Node{element: element}
		} else {
			n.left.insert(element)
		}
	} else {
		if n.right == nil {
			n.right = &Node{element: element}
		} else {
			n.right.insert(element)
		}
	}
}

func (t *Tree) Remove(element int) {
	if t.IsEmpty() {
		return
	}

	t.root = remove(element, t.root)
}

func remove(element int, n *Node) *Node {
	if n == nil {
		return nil
	}

	if n.element == element {

		if n.left == nil && n.right == nil {
			return nil
		} else if n.left == nil {
			//remove(element, n.right)
			return n.right
		} else if n.right == nil {
			//remove(element, n.left)
			return n.left
		}

		min := n.right.smallest()
		n.SetElement(min.element)
		n.SetRight(*remove(min.element, n.right))

	} else if n.element >= element {
		if n.left == nil {
			return nil
		}
		left := remove(element, n.left)
		if left == nil {
			return nil
		}
		n.SetLeft(*left)
	} else {
		if n.right == nil {
			return nil
		}
		right := remove(element, n.left)
		if right == nil {
			return nil
		}
		n.SetRight(*right)
	}
	return n
}

func (t *Tree) Find(element int) *Node {
	if t.root == nil {
		return nil
	}

	return t.root.find(element)
}

func (n *Node) find(element int) *Node {
	if n.element == element {
		return n
	} else if n.element > element && n.left != nil {
		return n.left.find(element)
	} else if n.element < element && n.right != nil {
		return n.right.find(element)
	}
	return nil
}

func (t *Tree) Height() int {
	if t.IsEmpty() {
		return -1
	}

	return t.root.height(0)
}

func (n *Node) height(level int) int {
	heightLeft := level
	if n.left != nil {
		heightLeft = n.left.height(heightLeft + 1)
	}

	heightRight := level
	if n.right != nil {
		heightRight = n.right.height(heightRight + 1)
	}

	if heightRight > heightLeft {
		return heightRight
	}
	return heightLeft
}

func (t *Tree) PreOrder() []int {
	var collection []int

	if t.root == nil {
		return collection
	}

	return t.root.preOrder(collection)
}

func (n *Node) preOrder(collection []int) []int {
	collection = append(collection, n.element)

	if n.left != nil {
		collection = n.left.preOrder(collection)
	}

	if n.right != nil {
		collection = n.right.preOrder(collection)
	}

	return collection
}

func (t *Tree) InOrder() []int {
	var collection []int

	if t.root == nil {
		return collection
	}

	if t.root.left != nil {
		collection = t.root.left.inOrder(collection)
	}

	collection = append(collection, t.root.element)

	if t.root.right != nil {
		collection = t.root.right.inOrder(collection)
	}

	return collection
}

func (n *Node) inOrder(collection []int) []int {

	if n.left != nil {
		collection = n.left.inOrder(collection)
	}

	collection = append(collection, n.element)

	if n.right != nil {
		collection = n.right.inOrder(collection)
	}

	return collection
}

func (t *Tree) PosOrder() []int {
	var collection []int

	if t.root == nil {
		return collection
	}

	if t.root.left != nil {
		collection = t.root.left.posOrder(collection)
	}

	if t.root.right != nil {
		collection = t.root.right.posOrder(collection)
	}

	collection = append(collection, t.root.element)

	return collection
}

func (n *Node) posOrder(collection []int) []int {

	if n.left != nil {
		collection = n.left.posOrder(collection)
	}

	if n.right != nil {
		collection = n.right.posOrder(collection)
	}

	collection = append(collection, n.element)

	return collection
}

func (t *Tree) String() string {
	if t.root == nil {
		return ""
	}

	return t.root.toString(0, "")
}

func (n *Node) toString(level int, str string) string {
	nextLevel := level + 1

	if n.right != nil {
		str = n.right.toString(nextLevel, str)
	}

	if level != 0 {
		for i := 0; i < level-1; i++ {
			str = str + "|\t"
		}
		str = str + "|-------" + strconv.Itoa(n.element) + "\n"
	} else {
		str = str + strconv.Itoa(n.element) + "\n"
	}

	if n.left != nil {
		str = n.left.toString(nextLevel, str)
	}

	return str
}

func (t *Tree) GetNodesByLevel(level int) []*Node {

	return t.root.getNodesByLevel(level, 0, []*Node{})
}

func (n *Node) getNodesByLevel(level int, currentLevel int, nodes []*Node) []*Node {
	nextLevel := currentLevel + 1

	resultNodes := nodes

	if level == currentLevel {
		resultNodes = append(resultNodes, n)
	}

	if n.left != nil {
		resultNodes = n.left.getNodesByLevel(level, nextLevel, resultNodes)
		//resultNodes = append(resultNodes, leftNodes[len(leftNodes)-1])
	}

	if n.right != nil {
		resultNodes = n.right.getNodesByLevel(level, nextLevel, resultNodes)
		//resultNodes = append(resultNodes, rightNodes[len(rightNodes)-1])
	}

	return resultNodes
}
