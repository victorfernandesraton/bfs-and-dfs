package dfs

import "github.com/victorfernandesraton/bfs-and-dfs/node"

// Execution é a função que implementa o algoritimo DFS na estrutura de nós previamente criada
func Execution(n *node.Node, out *node.Output) *node.Output {
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
		for _, item := range n.Children { // para cada vertice verifica se esse foi usado
			if item.Used == false {
				// Marca como visitado
				item.Used = true
				out.Queue = append(out.Queue, item) // adiciona o vetice raiz cna listga de queues
				Execution(item, out)
			}
		}
	} else if len(n.Parent) > 0 { // Verificação nos vertices pais caso não tenha filho a ser  percorrido
		for _, item := range n.Parent { // para cada vertice verifica se esse foi usado
			if item.Used == false {
				// Marca como visitado
				item.Used = true
				out.Queue = append(out.Queue, item) // adiciona o vetice raiz cna listga de queues
				Execution(item, out)
			}
		}
	}
	return out
}
