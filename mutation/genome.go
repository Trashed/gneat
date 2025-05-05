package mutation

import (
	"errors"
	"math/rand"

	"github.com/Trashed/gneat/genetics"
)

func ApplyRandWeights(g *genetics.Genome, weightMutationRate float64, weightPerturbationStrength float64) {

	for _, gene := range g.Genes {
		if rand.Float64() >= weightMutationRate {
			ApplyRandWeight(gene, weightMutationRate, weightPerturbationStrength)
		}
	}
}

// AddNewGene creates a new gene between two existing nodes.
func AddNewGene(g *genetics.Genome, innovation *genetics.InnovationStore) error {

	// TODO: Get possible

	return errors.New("not implemented")
}

/*
// TODO: Write implementation
func AddNewRecurrentGene(g *genetics.Genome, innovation *genetics.InnovationStore) {})
*/
