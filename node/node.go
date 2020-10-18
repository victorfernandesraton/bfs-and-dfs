package node

// Node define a estrutura de um vertice podendo ter childresn, parents e o seu indice
// OBS: o indice é atribuido ao vertice poois este é usado nos algoritimos de bfs e dfss
type Node struct {
	Value    interface{}
	Index    uint64
	Children []*Node
	Parent   []*Node
	Used     bool
}

// Output represeenta a saida encontrada com a queue e o ultimo level registtrado
type Output struct {
	Queue     []*Node
	LastLevel uint64
}

// IsLeaf as a function to detect if this elemt is leaf
func (n *Node) IsLeaf() bool {
	return len(n.Children) <= 0
}

// IsRoot verifica se o vertice corrente é uma raiz, ou seja sem parents
func (n *Node) IsRoot() bool {
	return len(n.Parent) == 0
}

// AddChildren aceita um outro vertice e adiciona este a lista de vertices filhos
func (n *Node) AddChildren(m *Node) {
	if m.Index == 0 {
		m.Index = n.Index + 1
	}
	m.Parent = append(m.Parent, n)
	n.Children = append(n.Children, m)
}
