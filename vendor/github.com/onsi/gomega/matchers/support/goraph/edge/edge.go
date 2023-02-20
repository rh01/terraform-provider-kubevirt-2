package edge

import . "github.com/onsi/gomega/matchers/support/goraph/node"

type Edge struct {
<<<<<<< HEAD
	Node1 int
	Node2 int
=======
	Node1 Node
	Node2 Node
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
}

type EdgeSet []Edge

func (ec EdgeSet) Free(node Node) bool {
	for _, e := range ec {
<<<<<<< HEAD
		if e.Node1 == node.ID || e.Node2 == node.ID {
=======
		if e.Node1 == node || e.Node2 == node {
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
			return false
		}
	}

	return true
}

func (ec EdgeSet) Contains(edge Edge) bool {
	for _, e := range ec {
		if e == edge {
			return true
		}
	}

	return false
}

func (ec EdgeSet) FindByNodes(node1, node2 Node) (Edge, bool) {
	for _, e := range ec {
<<<<<<< HEAD
		if (e.Node1 == node1.ID && e.Node2 == node2.ID) || (e.Node1 == node2.ID && e.Node2 == node1.ID) {
=======
		if (e.Node1 == node1 && e.Node2 == node2) || (e.Node1 == node2 && e.Node2 == node1) {
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
			return e, true
		}
	}

	return Edge{}, false
}

func (ec EdgeSet) SymmetricDifference(ec2 EdgeSet) EdgeSet {
	edgesToInclude := make(map[Edge]bool)

	for _, e := range ec {
		edgesToInclude[e] = true
	}

	for _, e := range ec2 {
		edgesToInclude[e] = !edgesToInclude[e]
	}

	result := EdgeSet{}
	for e, include := range edgesToInclude {
		if include {
			result = append(result, e)
		}
	}

	return result
}
