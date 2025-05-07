package mutation

import (
	"math/rand"

	"github.com/Trashed/gneat/genetics"
	gneatrand "github.com/Trashed/gneat/rand"
)

func ApplyRandWeights(g *genetics.Genome, weightMutationRate float64, weightPerturbationStrength float64) {

	for _, gene := range g.Genes {
		if rand.Float64() >= weightMutationRate {
			ApplyRandWeight(gene, weightMutationRate, weightPerturbationStrength)
		}
	}
}

// AddNewGene creates a new gene between two existing nodes.
func AddNewGene(g *genetics.Genome, innovation *genetics.InnovationStore, randFunc func(items gneatrand.Listable) uint) (*genetics.Gene, error) {
	randomId1 := randFunc(g.Nodes)
	randomId2 := randFunc(g.Nodes)

	nodeIn := g.Nodes[randomId1]
	nodeOut := g.Nodes[randomId2]

	gene, err := innovation.CreateGene(nodeIn, nodeOut)
	if err != nil {
		return nil, err
	}

	return gene, nil
}
