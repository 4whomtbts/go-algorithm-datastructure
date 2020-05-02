package main

import "fmt"

type NodeColorType bool

const (
	BLACK = true
	RED = false
)
type Node struct {
	color NodeColorType
	key int64
	parent *Node
	left *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) insertFix(node *Node) {
	for node.parent.color == RED {
		if node.parent == node.parent.parent.left {
			var uncle *Node = node.parent.parent.right
			if uncle != nil && uncle.color == RED {
				node.parent.color = BLACK
				uncle.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			}
		}
	}
}

func (t *Tree) insert(key int64) {
	var newNode *Node = &Node{RED, key, nil, nil, nil}
	var temp *Node = nil
	var curr *Node = t.root

	for curr != nil {
		temp = curr
		if newNode.key < curr.key {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	newNode.parent = temp
	if temp == nil {
		t.root = newNode
	} else if newNode.key < temp.key {
		temp.left = newNode
	} else {
		temp.right = newNode
	}
	newNode.color = RED
	newNode.parent = temp
	t.fixUp(newNode)
}

func (t *Tree) leftRotate(x *Node) {
	var y = x.right;
	if y.left != nil {
		y.left.parent = x
	}
	x.right = y.left;

	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	x.parent = y
	y.left = x
}

func (t *Tree) rightRotate(x *Node) {
	var y = x.left
	if y.right != nil {
		y.right.parent = x
	}
	x.left = y.right

	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	x.parent = y
	y.right = x
}

func (t *Tree) fixUp(node *Node) {
	if t.root == node {
		t.root.color = BLACK
		return
	}

	for node.parent != nil && node.parent.color == RED {
		if node.parent == node.parent.parent.left {
			var y = node.parent.parent.right
			if y != nil && y.color == RED {
				node.parent.color = BLACK
				y.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else if node == node.parent.right {
				node = node.parent
				t.leftRotate(node)
			} else {
				node.parent.color = BLACK
				node.parent.parent.color = RED
				t.rightRotate(node.parent.parent)
			}
		} else {
			var y = node.parent.parent.left
			if  y != nil && y.color == RED {
				node.parent.color = BLACK
				y.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else if node ==  node.parent.left {
				node = node.parent
				t.rightRotate(node)
				node.parent.color = BLACK
				node.parent.parent.color = RED
				t.leftRotate(node.parent.parent)
			} else {
				node.parent.color = BLACK
				node.parent.parent.color = RED
				t.leftRotate(node.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

func (t* Tree) transplant(u, v *Node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left =  v
	} else {
		u.parent.right = v
	}

	if v != nil {
		v.parent = u.parent
	}
}

func (t *Tree) deleteFixUp(node *Node) {
	for node != t.root && node.color == BLACK {
		if node == node.parent.left	 {
			var w *Node = node.parent.right
			if w.color == RED {
				w.color = BLACK
				node.parent.color =RED
				t.leftRotate(node.parent)
				w = node.parent.right
			}

			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				node = node.parent
			} else if w.right.color == BLACK {
				w.left.color = BLACK
				w.color = RED
				t.rightRotate(w)
				w = node.parent.right
			} else {
				w.color = node.parent.color
				node.parent.color = BLACK
				w.right.color = BLACK
				t.leftRotate(node.parent)
				node = t.root
			}
		} else {
			var w *Node = node.parent.left
			if w.color == RED {
				w.color = BLACK
				node.parent.color = RED
				t.rightRotate(node.parent)
				w = node.parent.left
			}

			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				node = node.parent
			} else if w.left.color == BLACK {
				w.right.color = BLACK
				w.color = RED
				t.leftRotate(w)
				w = node.parent.left
			} else {
				w.color = node.parent.color
				node.parent.color = BLACK
				w.left.color = BLACK
				t.rightRotate(node.parent)
				node = t.root
			}
		}
	}
	node.color = BLACK
}

func (t *Tree) getTreeMinimum(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (t *Tree) delete(key int64) bool {

	if t.root == nil {
		return false
	}

	var z *Node = t.root

	for z.key != key {
		if key > z.key {
			z = z.right
		} else {
			z = z.left
		}
	}

	var y *Node = z
	var x *Node = nil
	var yOriginalColor = y.color
	if z.left == nil {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == nil {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y = t.getTreeMinimum(z.right)
		yOriginalColor = y.color
		x = y.right
		if y.parent == z {
			x.parent = y
		} else {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	if yOriginalColor == BLACK {
		t.deleteFixUp(x)
	}

	return true
}

func (t *Tree) inOrder(node *Node) {
	if node == nil {
		return
	}

	t.inOrder(node.left)
	var color string
	if node.color == BLACK {
		color = "BLACK"
	} else {
		color = "RED"
	}
	fmt.Println("[" ,node.key, ",", color, "]")
	t.inOrder(node.right)
}

func main () {
	var tree *Tree = &Tree{}
	tree.insert(15)
	tree.insert(8)
	tree.insert(16)
	tree.insert(24)
	tree.insert(3)
	tree.insert(2)
	tree.delete(2)
	tree.delete(8)
	tree.inOrder(tree.root)
}