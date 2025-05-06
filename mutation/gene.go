package mutation

import (
	"math/rand"

	"github.com/Trashed/gneat/genetics"
)

// ApplyRandWeight mutates the weight of a gene by adding a random value.
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

// AddNewNode inserts a new node by splitting an existing gene.
// It creates a new gene and a new node, and returns them.
func AddNewNode(g *genetics.Gene, innovation *genetics.InnovationStore) ([]*genetics.Gene, *genetics.Node) {

	g.Enabled = false

	newNode := genetics.NewNode(genetics.NodeHidden)
	newGene1 := &genetics.Gene{
		InNode:    g.InNode,
		OutNode:   newNode,
		Weight:    g.Weight,
		Enabled:   true,
		Recurrent: false,
	}
	innovation.PushGene(newGene1)

	newGene2 := &genetics.Gene{
		InNode:    newNode,
		OutNode:   g.OutNode,
		Weight:    1.0,
		Enabled:   true,
		Recurrent: false,
	}
	innovation.PushGene(newGene2)

	return []*genetics.Gene{newGene1, newGene2}, newNode
}
