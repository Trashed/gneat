package mutation

import (
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
