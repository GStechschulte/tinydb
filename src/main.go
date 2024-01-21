package main

type Node struct {
	val    int
	red    bool
	parent *Node
	left   *Node
	right  *Node
}

type RBTree struct {
	nil  *Node
	root *Node
}

func NewRBTree() *RBTree {
	nilNode := &Node{red: false, val: 0}
	return &RBTree{nil: nilNode, root: nilNode}
}

func (t *RBTree) Insert(val int) {
	// Initialize binary search insertion
	newNode := &Node{val: val, red: true, left: t.nil, right: t.nil}

	// Red-black tree insert algorithm
	var parent *Node
	current := t.root
	for current != t.nil {
		parent = current
		if newNode.val < current.val {
			current = current.left
		} else if newNode.val > current.val {
			current = current.right
		} else {
			return
		}
	}

	// Set the parent node and insert the new node into the tree
	newNode.parent = parent
	if parent == t.nil {
		t.root = newNode
	} else if newNode.val < parent.val {
		parent.left = newNode
	} else {
		parent.right = newNode
	}

	t.fixInsert(newNode)
}

func (t *RBTree) fixInsert(newNode *Node) {
	for newNode.parent.red {
		if newNode.parent == newNode.parent.parent.left {
			uncle := newNode.parent.parent.right
			if uncle.red {
				newNode.parent.red = false
				uncle.red = false
				newNode.parent.parent.red = true
				newNode = newNode.parent.parent
			} else {
				if newNode == newNode.parent.right {
					newNode = newNode.parent
					t.leftRotate(newNode)
				}
				newNode.parent.red = false
				newNode.parent.parent.red = true
				t.rightRotate(newNode.parent.parent)
			}
		} else {
			uncle := newNode.parent.parent.left
			if uncle.red {
				newNode.parent.red = false
				uncle.red = false
				newNode.parent.parent.red = true
				newNode = newNode.parent.parent
			} else {
				if newNode == newNode.parent.left {
					newNode = newNode.parent
					t.rightRotate(newNode)
				}
				newNode.parent.red = false
				newNode.parent.parent.red = true
				t.leftRotate(newNode.parent.parent)
			}
		}
	}

	// Maintain property that root node is always black
	t.root.red = false
}

func (t *RBTree) leftRotate(x *Node) {
	y := x.right
	x.right = y.left
	if y.left != t.nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == t.nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *RBTree) rightRotate(x *Node) {
	y := x.left
	x.left = y.right
	if y.right != t.nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == t.nil {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

func main() {

	tree := NewRBTree()

	// Run a loop to insert values into the tree
	for i := 0; i < 10; i++ {
		tree.Insert(i)
	}

}
