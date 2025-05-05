package genetics

func CopyGenome(source *Genome) *Genome {

	nodes, genes := copy(source.Genes)

	return &Genome{
		Id:    source.Id,
		Nodes: nodes,
		Genes: genes,
	}
}

func CopyGene(g *Gene) *Gene {
	return &Gene{
		Innovation: g.Innovation,
		Weight:     g.Weight,
		Enabled:    g.Enabled,
		Recurrent:  g.Recurrent,
		InNode:     CopyNode(g.InNode),
		OutNode:    CopyNode(g.OutNode),
	}
}

func CopyNode(n *Node) *Node {
	return &Node{
		Id:       n.Id,
		NodeType: n.NodeType,
	}
}

func copy(source map[uint]*Gene) (Nodes, map[uint]*Gene) {
	newNodes := make(Nodes, 0)
	newGenes := make(map[uint]*Gene, 0)

	for _, g := range source {

		if canInsert(g.InNode, newNodes) {
			newNodes[g.InNode.Id] = g.InNode
		}
		if canInsert(g.OutNode, newNodes) {
			newNodes[g.OutNode.Id] = g.OutNode
		}

		newGenes[g.Innovation] = CopyGene(g)
	}

	return newNodes, newGenes
}

func canInsert(n *Node, nodes Nodes) bool {
	for _, existing := range nodes {
		if n.Id == existing.Id {
			return false
		}
	}

	return true
}
