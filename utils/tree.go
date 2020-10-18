package utils

import (
	"errors"

	"github.com/victorfernandesraton/bfs-and-dfs/node"
)

func generateTree(matrix [][]uint8, n []*node.Node) (*node.Node, error) {
	if !IsAdjacence(matrix) {
		return nil, errors.New("Invalid adjacent matrix")
	}

	if len(matrix) != len(n) {
		return nil, errors.New("Invalid adjacent matrix or vector list this is need same length")
	}

	for kline, vline := range matrix {
		vertex := n[kline]
		for kcol, vcol := range vline {
			if vcol == 1 {
				vertex.AddChildren(n[kcol])
			}
		}
	}
	return n[0], nil
}
