package dlx

// Cover removes the row containing the given node, as well as all columns with
// a node on that row and all rows with a node on those columns
func Cover(row *Node) {
	coverCol(row)
	for cur := row.right; cur != row; cur = cur.right {
		coverCol(cur)
	}
}

// Uncover undoes the effects of Cover
func Uncover(row *Node) {
	for cur := row.left; cur != row; cur = cur.left {
		uncoverCol(cur)
	}
	uncoverCol(row)
}

func coverCol(n *Node) {
	// Detach column
	detachHor(n.header)
	for row := n.header.down; row.header != row; row = row.down {
		// Detach each row with a non-zero element in the column
		detachVert(row)
		for cur := row.left; cur != row; cur = cur.left {
			detachVert(cur)
		}
	}
}

func uncoverCol(n *Node) {
	// Reattach column
	reattachHor(n.header)
	for row := n.header.up; row.header != row; row = row.up {
		// Reattach each row with a non-zero element in the column
		for cur := row.right; cur != row; cur = cur.right {
			reattachVert(cur)
		}
		reattachVert(row)
	}
}

// Solve solves the exact cover problem represented by our matrix.
// If successful, the covered nodes are held in the solution slice.
func Solve(root *Node, solution *[]int) bool {
	if root.left == root {
		return true
	}
	head, size := smallestColumn(root)
	if size == 0 {
		return false
	}

	for row := head.down; row != head; row = row.down {
		Cover(row)
		if Solve(root, solution) {
			*solution = append(*solution, row.possiblity)
			return true
		}
		Uncover(row)
	}
	return false
}

// Find returns the leftmost node in the row representing the given possibility
func Find(possibility int, root *Node) *Node {
	for col := root.right; col != root; col = col.right {
		for cur := col.down; cur != col; cur = cur.down {
			if possibility == cur.possiblity {
				return cur
			}
		}
	}
	return nil
}

// smallestColumn returns the head node for the shortest column with its size
func smallestColumn(root *Node) (*Node, int) {
	min := 9999 // TODO: replace with maximum possible number of rows
	var minCol *Node
	for col := root.left; col != root; col = col.left {
		count := columnSize(col)
		if count < min {
			min = count
			minCol = col
		}
	}
	return minCol, min
}

func columnSize(n *Node) int {
	count := 0
	for cur := n.down; cur != n; cur = cur.down {
		count++
	}
	return count
}

func detachHor(n *Node) {
	n.right.left = n.left
	n.left.right = n.right
}

func reattachHor(n *Node) {
	n.right.left = n
	n.left.right = n
}

func detachVert(n *Node) {
	n.up.down = n.down
	n.down.up = n.up
}

func reattachVert(n *Node) {
	n.up.down = n
	n.down.up = n
}
