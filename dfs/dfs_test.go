package dfs

import (
	"testing"

	"github.com/victorfernandesraton/grafos/node"
)

func TestIsValid(t *testing.T) {
	zero := &node.Node{
		Value: 0,
		Index: 0,
	}
	one := &node.Node{
		Value: 1,
		Index: 0,
	}
	two := &node.Node{
		Value: 2,
		Index: 0,
	}
	one.AddChildren(zero)
	one.AddChildren(two)
	three := &node.Node{
		Value: 3,
		Index: 0,
	}
	another := &node.Node{
		Value: "F",
		Index: 0,
	}
	two.AddChildren(three)
	two.AddChildren(another)
	res := Dfsexecution(one, &node.Output{})

	if res.LastLevel != 2 {
		t.Errorf("Not expected index %v", res.LastLevel)
	}

	if len(res.Queue) != 5 {
		t.Errorf("Unexpected query elements expected 5 have %v", len(res.Queue))
	}
}
