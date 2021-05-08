# Golang Trees

Golang repository with Go Binary Trees examples and resources.

## 1. Binary Search Tree

Binary Search Tree (BST) data structure developed using Golang.

### 1.1. Methods

Following methods are currently available to manage a Binary Search Tree.

* ``NewTree()`` - Creates a new BST.
* ``Root()`` - Returns BST root node.
* ``IsEmpty()`` - Returns a ``true`` if the tree doesn't have any node.
* ``Smallest()`` - Returns the smallest element number in the BST.
* ``Insert(element int)`` - Inserts a new element in the BST.
* ``Remove(element int)`` - Removes an existing element from the BST.
* ``PreOrder()`` - Returns a ``[]int`` slice using pre-order algorithm.
* ``InOrder()`` - Returns a ``[]int`` slice using in-order algorithm.
* ``PosOrder()`` - Returns ``[]int`` slice using pos-order algorithm.
* ``String()`` - Returns tree readable``string`` output.
* ``Height()`` - Returns Tree height ``int`` value. Returns ``0`` if tree is empty.
* ``Diameter()`` - Returns Tree diameter ``int`` value. Returns ``0`` if tree is empty.
* ``GetNodesByLevel(level int)`` - Returns a ``[]*Node`` with all Nodes from the inserted level number.

### 1.2. Example

A new BST can be created as in the following example:

```
tree := NewTree()

tree.Insert(20)
```

## 2. Linked List

Linked List data structure developed using Golang.

### 2.1. Methods

Following methods are currently available to manage a Linked List.

* ``New()`` - Creates a new Linked List and returns its pointer.
* ``IsEmpty()`` - Returns a ``true`` if the Linked List doesn't have any element.
* ``Insert(element int)`` - Inserts a new element in the Linked List.
* ``RemoveMin(element int)`` - Removes the first element from Linked List and returns it value.
* ``Size()`` - Returns the number of elements in the List.
* ``RecountSize()`` - Uses an algorithm to calculate and return Linked List the number of elements.

### 2.2. Example

A new Linked List can be created as in the following example:

```
list := linked_list.New()

list.Insert(20)
```
