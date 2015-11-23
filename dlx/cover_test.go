package dlx

import "testing"

// Dummy matrix for testing
// 1 1 0
// 0 0 1
// 0 1 1
func buildDummy() *Node {
	root := NewRoot()
	c := make([]*Node, 3)
	for i := range c {
		c[i] = AddHeader(root)
	}
	matrix := [][]int{{1, 1, 0}, {0, 0, 1}, {0, 1, 1}}
	for i, r := range matrix {
		var row []*Node
		for j, v := range r {
			if v != 0 {
				node := AddNode(i, c[j])
				row = append(row, node)
			}
		}
		BuildRow(row)
	}
	return root
}

// Test covering on r1c1 of the matrix
// Only r2c3 should remain
func TestCover(t *testing.T) {
	root := buildDummy()
	c1 := root.right
	c3 := root.left
	r1c1 := c1.down
	r2c3 := c3.down
	r3c3 := r2c3.down
	Cover(r1c1)
	if c1.left.right == c1 || c1.right.left == c1 {
		t.Fatal("c1 not removed")
	}
	if c3.down != r2c3 {
		t.Fatal("uncovered cell r2c3 removed")
	}
	if c3.up == r3c3 {
		t.Fatal("r3c3 not removed")
	}
}
