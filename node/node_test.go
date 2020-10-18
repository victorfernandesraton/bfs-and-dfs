package node

import (
	"testing"
)

func TestStruct(t *testing.T) {
	n := &Node{
		Value: 12,
	}
	t.Run("expected not leaf", func(t *testing.T) {
		if n.IsLeaf() == false {
			t.Error("isLeaf Not expected")
		}
	})
	t.Run("expected not root", func(t *testing.T) {
		if n.IsRoot() != true {
			t.Errorf("isRoot not expected got %v", len(n.Parent))
		}
	})
}

func TestAllocation(t *testing.T) {
	n := &Node{
		Value: 2,
	}
	n.AddChildren(&Node{
		Value: 5,
	})

	t.Run("expected leaf", func(t *testing.T) {
		if n.IsLeaf() == true {
			t.Error("isLeaf Not expected")
		}
	})
	t.Run("expected not root", func(t *testing.T) {
		if n.IsRoot() == false {
			t.Error("isRoot not expected")
		}
	})
}

func TestReference(t *testing.T) {
	n := &Node{
		Value: 1,
	}
	m := &Node{
		Value: 3,
	}
	n.AddChildren(m)

	if n.IsLeaf() == true {
		t.Error("n isLeaf Not expected")
	}
	if n.IsRoot() == false {
		t.Error("n isRoot not expected")
	}
	if m.IsLeaf() == false {
		t.Error("m isLeaf Not expected")
	}
	if m.IsRoot() == true {
		t.Error("m isRoot not expected")
	}
	if m.Parent[0] != n {
		t.Error("m reference not found")
	}
	if n.Children[0] != m {
		t.Error("n reference not found")
	}
	if m.Index != 1 {
		t.Error("m index not valid")
	}
	if n.Index != 0 {
		t.Error("n index not valid")
	}
}
