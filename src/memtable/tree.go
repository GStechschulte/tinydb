package memtable

import (
	"fmt"
	"os"
	"time"
)

type Color bool

const (
	Red   Color = true
	Black Color = false
)

const (
	keySize    = 24
	valSize    = 4
	colorSize  = 1
	leftSize   = 8
	rightSize  = 8
	parentSize = 8
	approxSize = keySize + valSize + colorSize + leftSize + rightSize + parentSize
)

type Node struct {
	key                 time.Time
	val                 float32
	color               Color
	left, right, parent *Node
}

type Tree struct {
	root      *Node
	totalSize uint16
}

func NewTree() *Tree {
	nilNode := &Node{color: Black, val: 0.0}
	return &Tree{root: nilNode, totalSize: 0}
}

func newNode(key time.Time, val float32) *Node {
	return &Node{key: key, val: val, color: Red}
}

func (t *Tree) Insert(key time.Time, val float32) {
	newNode := newNode(key, val)

	if t.root == nil {
		t.root = newNode
		return
	} else {
		t.insertNode(t.root, newNode)
	}

	t.fixInsert(newNode)

	// Update size of tree as a new node is inserted
	t.totalSize += approxSize

}

func (t Tree) Size() uint16 {
	return t.totalSize
}

// Recursive insertion of nodes
func (t *Tree) insertNode(root, node *Node) {
	if node.key.Before(root.key) {
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

// InOrderTraversal to return the red-black tree in sorted order
func InOrderTraversal(t *Tree) map[time.Time]float32 {
	result := make(map[time.Time]float32)

	var traverse func(node *Node)
	traverse = func(node *Node) {
		if node != nil {
			traverse(node.left)
			result[node.key] = node.val
			traverse(node.right)
		}
	}
	traverse(t.root)
	return result
}

// Flush the memtable to disk
func (t *Tree) Flush(path string) error {

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	entries := InOrderTraversal(t)

	return flushMemTable(entries, f)
}

func flushMemTable(entries map[time.Time]float32, f *os.File) error {

	for key, val := range entries {
		_, err := f.WriteString(fmt.Sprintf("%s:%f\n", key.Format(time.UnixDate), val))
		if err != nil {
			return err
		}
	}

	f.Sync()

	return nil
}

// Print the node value and color
func (t *Tree) printTree(node *Node, prefix string, isTail bool) {
	if node != nil {
		fmt.Printf("")
		if isTail {
			fmt.Printf("└── ")
			prefix += "    "
		} else {
			fmt.Printf("├── ")
			prefix += "│   "
		}
		color := "R"
		if node.color == Black {
			color = "B"
		}
		fmt.Printf("%s: %v (%s)\n", node.key, node.val, color)
		t.printTree(node.right, prefix, false)
		t.printTree(node.left, prefix, true)
	}
}

func (t *Tree) PrintTree() {
	t.printTree(t.root, "", true)
}
