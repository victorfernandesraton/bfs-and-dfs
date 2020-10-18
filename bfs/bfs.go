package bfs

import (
	"github.com/victorfernandesraton/grafos/node"
)

// BfsExecution é a função que pega um vertice de uma arvore qualquer e implementa o algoritimo
func BfsExecution(n *node.Node, out *node.Output) *node.Output {
	// Marca o vertice raiz como usado
	if n.Used == false {
		n.Used = true
		out.Queue = append(out.Queue, n) // adiciona o vetice raiz cna listga de queues
	}
	// Verifica o nivel do vertice de acordo com seu parent
	if out.LastLevel < n.Index {
		out.LastLevel = n.Index
	}
	if len(n.Children) > 0 {
		// se o vertice atual tiver adjacencia, percorre estas
		out.LastLevel = n.Index           // adiciona um level na marcação dos levels
		for _, item := range n.Children { // para cada vertice verifica se esse foi usado
			if item.Used == false {
				// Marca como visitado
				item.Used = true
				out.Queue = append(out.Queue, item) // adiciona o vetice raiz cna listga de queues
			}
		}
		// Percorre executando bfs
		for _, item := range n.Children {
			BfsExecution(item, out)
		}
	}
	// a kista de querys e retornada como saida
	return out
}
