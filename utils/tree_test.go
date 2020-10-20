package utils

import (
	"testing"

	"github.com/victorfernandesraton/bfs-and-dfs/node"
)

func ExpectedValue(t *testing.T) {
	vlist := []*node.Node{}
	vlist = append(vlist, &node.Node{Value: 12})
	vlist = append(vlist, &node.Node{Value: 1})
	vlist = append(vlist, &node.Node{Value: "A"})
	vlist = append(vlist, &node.Node{Value: "F"})

	matrix := [][]uint8{}
	matrix = append(matrix, []uint8{0, 1, 0, 1})
	matrix = append(matrix, []uint8{1, 0, 1, 0})
	matrix = append(matrix, []uint8{0, 1, 0, 0})
	matrix = append(matrix, []uint8{1, 0, 0, 0})

	tree, err := GenerateTree(matrix, vlist)
	t.Run("expected error", func(t *testing.T) {
		if err != nil {
			t.Errorf("Unexpected error, %s", err)
		}
	})
	t.Run("expected tree root", func(t *testing.T) {
		if tree.Value != 12 {
			t.Errorf("Unexpected tree root, want 12 have %s", tree.Value)
		}
	})
	t.Run("expected children for tree rott", func(t *testing.T) {
		if tree.Children[0].Value != 1 {
			t.Errorf("Unexpected tree root, want 1 have %s", tree.Children[0].Value)
		}
	})
	t.Run("expected value for children of element tree", func(t *testing.T) {
		if tree.Children[1].Value != "A" {
			t.Errorf("Unexpected tree root, want A have %s", tree.Children[1].Value)
		}
	})
	t.Run("expected value (children of children)", func(t *testing.T) {
		if tree.Children[0].Children[0].Value != "F" {
			t.Errorf("Unexpected tree root, want F have %s", tree.Children[0].Children[0].Value)
		}
	})
}

func InValidMatrix(t *testing.T) {
	vlist := []*node.Node{}
	vlist = append(vlist, &node.Node{Value: 12})
	vlist = append(vlist, &node.Node{Value: 1})
	vlist = append(vlist, &node.Node{Value: "A"})
	vlist = append(vlist, &node.Node{Value: "F"})
	vlist = append(vlist, &node.Node{Value: "F"})

	matrix := [][]uint8{}
	matrix = append(matrix, []uint8{0, 1, 0, 1})
	matrix = append(matrix, []uint8{1, 0, 1, 0})
	matrix = append(matrix, []uint8{0, 1, 0, 0})
	matrix = append(matrix, []uint8{1, 0, 0, 0})
	matrix = append(matrix, []uint8{1, 0, 0, 0})

	if _, err := GenerateTree(matrix, vlist); err.Error() != "Invalid adjacent matrix" {
		t.Errorf("Not have expected error, expected :  Invalid adjacent matrix got, %s", err.Error())
	}
}

func InvalidSizeMatrixAndVertex(t *testing.T) {
	vlist := []*node.Node{}
	vlist = append(vlist, &node.Node{Value: 12})
	vlist = append(vlist, &node.Node{Value: 1})
	vlist = append(vlist, &node.Node{Value: "A"})
	vlist = append(vlist, &node.Node{Value: "F"})

	matrix := [][]uint8{}
	matrix = append(matrix, []uint8{0, 1, 0, 1})
	matrix = append(matrix, []uint8{1, 0, 1, 0})
	matrix = append(matrix, []uint8{0, 1, 0, 0})
	matrix = append(matrix, []uint8{1, 0, 0, 0})

	if _, err := GenerateTree(matrix, vlist); err.Error() != "Invalid adjacent matrix or vector list this is need same length" {
		t.Errorf("Not have expected error, expected :  Invalid adjacent matrix or vector list this is need same length got, %s", err.Error())
	}
}
