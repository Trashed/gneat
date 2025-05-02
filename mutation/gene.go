package mutation

import (
	"math/rand"

	"github.com/Trashed/gneat/genetics"
)

func ApplyRandWeight(g *genetics.Gene, weightMutationRate float64, weightPerturbationStrength float64) {
	randVal := rand.Float64()
	negativeMultiplier := func() float64 {
		if rand.Float64() < 0.5 {
			return -1
		}
		return 1
	}()

	g.Weight += negativeMultiplier * (randVal * weightPerturbationStrength)
}
