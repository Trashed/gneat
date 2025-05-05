/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package genetics

type Genome struct {
	Id    uint
	Nodes Nodes
	Genes map[uint]*Gene // Connections between Nodes
}

func (g *Genome) PushGene(gene *Gene) {

	if g.Nodes == nil {
		g.Nodes = make(Nodes)
	}

	if g.Genes == nil {
		g.Genes = make(map[uint]*Gene)
	}

	g.Nodes[gene.InNode.Id] = gene.InNode
	g.Nodes[gene.OutNode.Id] = gene.OutNode

	g.Genes[gene.Innovation] = gene
}
