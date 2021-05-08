package heap

type Heap struct {
	root *Element
	size int
}

type Element struct {
	value interface{}
	child *Element
}

func (h *Heap) IsEmpty() bool {
	if h.root == nil {
		return true
	}
	return false
}

func (h *Heap) Insert(value interface{}) {
	h.size++

	if h.root == nil {
		element := Element{
			value: value,
		}

		h.root = &element
		return
	}

	h.root.insert(value)
}

func (e *Element) insert(value interface{}) {

	if e.child == nil {
		e.child = &Element{
			value: value,
		}
		return
	}

	e.child.insert(value)
}

func (h *Heap) RemoveMin() interface{} {
	if h.root == nil {
		return nil
	}

	value := h.root.value

	if h.root.child == nil {
		h.root = nil
	} else {
		h.root = h.root.child
	}

	h.size--
	return value
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) RecountSize() int {
	if h.root == nil {
		return 0
	}

	return h.root.size(1)
}

func (e *Element) size(cont int) int {
	if e.child == nil {
		return cont
	}

	return e.child.size(cont + 1)
}
