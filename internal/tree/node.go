package tree

type Node struct {
	left    *Node
	element int
	right   *Node
}

func NewNode(left *Node, element int, right *Node) Node {
	return Node{
		left:    left,
		element: element,
		right:   right,
	}
}

func (n *Node) GetLeft() *Node {
	return n.left
}

func (n *Node) GetRight() *Node {
	return n.right
}

func (n *Node) GetElement() int {
	return n.element
}

func (n *Node) SetElement(element int) {
	n.element = element
}

func (n *Node) SetLeft(node Node) {
	n.left = &node
}

func (n *Node) SetRight(node Node) {
	n.right = &node
}
