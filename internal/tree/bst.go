package tree

type BST struct {
	root *Node
}

func NewBST() BST {
	return BST{
		root: nil,
	}
}

func (t *BST) Root() *Node {
	return t.root
}

func (t *BST) IsEmpty() bool {
	if t.root == nil {
		return true
	}
	return false
}

func (t *BST) Insert(element int) {
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
