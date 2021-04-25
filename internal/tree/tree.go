package tree

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

	t.root.remove(element)
}

func (n *Node) remove(element int) {
	if n.element == element {

		return
	} else if n.element >= element {
		n.GetLeft().remove(element)
	} else {
		n.GetRight().remove(element)
	}
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
