package memtable

import (
	"fmt"
	"strings"
)

type Color bool

const (
	Red   Color = true
	Black Color = false
)

type Node struct {
	val    int
	color  Color
	parent *Node
	left   *Node
	right  *Node
}

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	nilNode := &Node{color: Black, val: 0}
	return &Tree{root: nilNode}
}

func newNode(val int) *Node {
	return &Node{val: val, color: Red}
}

func (t *Tree) Insert(val int) {
	newNode := newNode(val)

	if t.root == nil {
		t.root = newNode
		return
	} else {
		t.insertNode(t.root, newNode)
	}

	t.fixInsert(newNode)

}

// Recursive insertion of nodes
func (t *Tree) insertNode(root, node *Node) {
	if node.val < root.val {
		if root.left == nil {
			root.left = node
			node.parent = root
		} else {
			t.insertNode(root.left, node)
		}
	} else {
		if root.right == nil {
			root.right = node
			node.parent = root
		} else {
			t.insertNode(root.right, node)
		}
	}
}

func (t *Tree) fixInsert(node *Node) {
	for node != t.root && node.parent.color == Red {
		if node.parent == node.parent.parent.left {
			uncle := node.parent.parent.right
			if uncle != nil && uncle.color == Red {
				// Case 1: Parent and uncle are red
				node.parent.color = Black
				uncle.color = Black
				node.parent.parent.color = Red
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					// Case 2: Node is right child and parent is left child
					node = node.parent
					t.rotateLeft(node)
				}
				// Case 3: Node is left child and parent is left child
				node.parent.color = Black
				node.parent.parent.color = Red
				t.rotateRight(node.parent.parent)
			}
		} else {
			uncle := node.parent.parent.left
			if uncle != nil && uncle.color == Red {
				// Case 1: Parent and uncle are red (mirror case)
				node.parent.color = Black
				uncle.color = Black
				node.parent.parent.color = Red
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					// Case 2: Node is left child and parent is right child (mirror case)
					node = node.parent
					t.rotateRight(node)
				}
				// Case 3: Node is right child and parent is right child (mirror case)
				node.parent.color = Black
				node.parent.parent.color = Red
				t.rotateLeft(node.parent.parent)
			}
		}
	}
	// Maintain the root color property
	t.root.color = Black
}

func (t *Tree) rotateLeft(x *Node) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *Tree) rotateRight(x *Node) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	color := "B" // Black
	if n.color == Red {
		color = "R" // Red
	}
	return fmt.Sprintf("%d%s", n.val, color)
}

func (t *Tree) PrintTree() {
	if t.root == nil {
		fmt.Println("Tree is empty")
		return
	}

	queue := []*Node{t.root}
	var level []string

	for len(queue) > 0 {
		size := len(queue)
		level = []string{}

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.String())

			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}

		fmt.Println(strings.Join(level, " "))
	}
}
