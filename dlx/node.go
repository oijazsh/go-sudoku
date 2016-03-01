package dlx

import "fmt"

type Node struct {
	possibility           int
	header                *Node
	up, down, left, right *Node
}

// NewNode returns a node representing a certain position in the matrix
func NewNode(poss int) *Node {
	return &Node{poss, nil, nil, nil, nil, nil}
}

func AddNode(row int, header *Node) *Node {
	n := NewNode(row)
	n.header = header
	n.down = header
	n.up = header.up
	header.up.down = n
	header.up = n
	return n
}

// AddHeader adds a new column header to the matrix and returns it
func AddHeader(root *Node) *Node {
	h := NewNode(-1)
	h.up, h.down, h.header = h, h, h
	h.left = root.left
	h.right = root
	root.left.right = h
	root.left = h
	return h
}

// NewRoot creates a new root node
func NewRoot() *Node {
	root := NewNode(-1)
	root.right = root
	root.left = root
	root.up = root
	root.down = root
	return root
}

// BuildRow connect a set of nodes in different columns into a row
func BuildRow(row []*Node) error {
	var poss = row[0].possibility
	for i := range row {

		if row[i].possibility != poss {
			return fmt.Errorf("dlx: All nodes in a row must represent the same possibility")
		}

		iLeft := i - 1
		iRight := (i + 1) % len(row)
		if iLeft < 0 {
			iLeft = len(row) - 1
		}
		row[i].left = row[iLeft]
		row[i].right = row[iRight]
	}
	return nil
}
