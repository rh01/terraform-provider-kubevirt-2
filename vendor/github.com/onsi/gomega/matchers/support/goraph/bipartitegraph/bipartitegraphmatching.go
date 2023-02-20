package bipartitegraph

<<<<<<< HEAD
import (
	. "github.com/onsi/gomega/matchers/support/goraph/edge"
	. "github.com/onsi/gomega/matchers/support/goraph/node"
	"github.com/onsi/gomega/matchers/support/goraph/util"
)

// LargestMatching implements the Hopcroft–Karp algorithm taking as input a bipartite graph
// and outputting a maximum cardinality matching, i.e. a set of as many edges as possible
// with the property that no two edges share an endpoint.
=======
import . "github.com/onsi/gomega/matchers/support/goraph/node"
import . "github.com/onsi/gomega/matchers/support/goraph/edge"
import "github.com/onsi/gomega/matchers/support/goraph/util"

>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
func (bg *BipartiteGraph) LargestMatching() (matching EdgeSet) {
	paths := bg.maximalDisjointSLAPCollection(matching)

	for len(paths) > 0 {
		for _, path := range paths {
			matching = matching.SymmetricDifference(path)
		}
		paths = bg.maximalDisjointSLAPCollection(matching)
	}

	return
}

func (bg *BipartiteGraph) maximalDisjointSLAPCollection(matching EdgeSet) (result []EdgeSet) {
	guideLayers := bg.createSLAPGuideLayers(matching)
	if len(guideLayers) == 0 {
		return
	}

<<<<<<< HEAD
	used := make(map[int]bool)
=======
	used := make(map[Node]bool)
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")

	for _, u := range guideLayers[len(guideLayers)-1] {
		slap, found := bg.findDisjointSLAP(u, matching, guideLayers, used)
		if found {
			for _, edge := range slap {
				used[edge.Node1] = true
				used[edge.Node2] = true
			}
			result = append(result, slap)
		}
	}

	return
}

func (bg *BipartiteGraph) findDisjointSLAP(
	start Node,
	matching EdgeSet,
	guideLayers []NodeOrderedSet,
<<<<<<< HEAD
	used map[int]bool,
=======
	used map[Node]bool,
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
) ([]Edge, bool) {
	return bg.findDisjointSLAPHelper(start, EdgeSet{}, len(guideLayers)-1, matching, guideLayers, used)
}

func (bg *BipartiteGraph) findDisjointSLAPHelper(
	currentNode Node,
	currentSLAP EdgeSet,
	currentLevel int,
	matching EdgeSet,
	guideLayers []NodeOrderedSet,
<<<<<<< HEAD
	used map[int]bool,
) (EdgeSet, bool) {
	used[currentNode.ID] = true
=======
	used map[Node]bool,
) (EdgeSet, bool) {
	used[currentNode] = true
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")

	if currentLevel == 0 {
		return currentSLAP, true
	}

	for _, nextNode := range guideLayers[currentLevel-1] {
<<<<<<< HEAD
		if used[nextNode.ID] {
=======
		if used[nextNode] {
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
			continue
		}

		edge, found := bg.Edges.FindByNodes(currentNode, nextNode)
		if !found {
			continue
		}

		if matching.Contains(edge) == util.Odd(currentLevel) {
			continue
		}

		currentSLAP = append(currentSLAP, edge)
		slap, found := bg.findDisjointSLAPHelper(nextNode, currentSLAP, currentLevel-1, matching, guideLayers, used)
		if found {
			return slap, true
		}
		currentSLAP = currentSLAP[:len(currentSLAP)-1]
	}

<<<<<<< HEAD
	used[currentNode.ID] = false
=======
	used[currentNode] = false
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
	return nil, false
}

func (bg *BipartiteGraph) createSLAPGuideLayers(matching EdgeSet) (guideLayers []NodeOrderedSet) {
<<<<<<< HEAD
	used := make(map[int]bool)
=======
	used := make(map[Node]bool)
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
	currentLayer := NodeOrderedSet{}

	for _, node := range bg.Left {
		if matching.Free(node) {
<<<<<<< HEAD
			used[node.ID] = true
=======
			used[node] = true
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
			currentLayer = append(currentLayer, node)
		}
	}

	if len(currentLayer) == 0 {
		return []NodeOrderedSet{}
	}
	guideLayers = append(guideLayers, currentLayer)

	done := false

	for !done {
		lastLayer := currentLayer
		currentLayer = NodeOrderedSet{}

		if util.Odd(len(guideLayers)) {
			for _, leftNode := range lastLayer {
				for _, rightNode := range bg.Right {
<<<<<<< HEAD
					if used[rightNode.ID] {
=======
					if used[rightNode] {
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
						continue
					}

					edge, found := bg.Edges.FindByNodes(leftNode, rightNode)
					if !found || matching.Contains(edge) {
						continue
					}

					currentLayer = append(currentLayer, rightNode)
<<<<<<< HEAD
					used[rightNode.ID] = true
=======
					used[rightNode] = true
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")

					if matching.Free(rightNode) {
						done = true
					}
				}
			}
		} else {
			for _, rightNode := range lastLayer {
				for _, leftNode := range bg.Left {
<<<<<<< HEAD
					if used[leftNode.ID] {
=======
					if used[leftNode] {
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
						continue
					}

					edge, found := bg.Edges.FindByNodes(leftNode, rightNode)
					if !found || !matching.Contains(edge) {
						continue
					}

					currentLayer = append(currentLayer, leftNode)
<<<<<<< HEAD
					used[leftNode.ID] = true
=======
					used[leftNode] = true
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
				}
			}

		}

		if len(currentLayer) == 0 {
			return []NodeOrderedSet{}
		}
		guideLayers = append(guideLayers, currentLayer)
	}

	return
}
