package bipartitegraph

import "fmt"

import . "github.com/onsi/gomega/matchers/support/goraph/node"
import . "github.com/onsi/gomega/matchers/support/goraph/edge"

type BipartiteGraph struct {
	Left  NodeOrderedSet
	Right NodeOrderedSet
	Edges EdgeSet
}

func NewBipartiteGraph(leftValues, rightValues []interface{}, neighbours func(interface{}, interface{}) (bool, error)) (*BipartiteGraph, error) {
	left := NodeOrderedSet{}
<<<<<<< HEAD
	for i, v := range leftValues {
		left = append(left, Node{ID: i, Value: v})
	}

	right := NodeOrderedSet{}
	for j, v := range rightValues {
		right = append(right, Node{ID: j + len(left), Value: v})
=======
	for i := range leftValues {
		left = append(left, Node{Id: i})
	}

	right := NodeOrderedSet{}
	for j := range rightValues {
		right = append(right, Node{Id: j + len(left)})
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
	}

	edges := EdgeSet{}
	for i, leftValue := range leftValues {
		for j, rightValue := range rightValues {
			neighbours, err := neighbours(leftValue, rightValue)
			if err != nil {
				return nil, fmt.Errorf("error determining adjacency for %v and %v: %s", leftValue, rightValue, err.Error())
			}

			if neighbours {
<<<<<<< HEAD
				edges = append(edges, Edge{Node1: left[i].ID, Node2: right[j].ID})
=======
				edges = append(edges, Edge{Node1: left[i], Node2: right[j]})
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
			}
		}
	}

	return &BipartiteGraph{left, right, edges}, nil
}
<<<<<<< HEAD

// FreeLeftRight returns left node values and right node values
// of the BipartiteGraph's nodes which are not part of the given edges.
func (bg *BipartiteGraph) FreeLeftRight(edges EdgeSet) (leftValues, rightValues []interface{}) {
	for _, node := range bg.Left {
		if edges.Free(node) {
			leftValues = append(leftValues, node.Value)
		}
	}
	for _, node := range bg.Right {
		if edges.Free(node) {
			rightValues = append(rightValues, node.Value)
		}
	}
	return
}
=======
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
