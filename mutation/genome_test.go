package mutation_test

import (
	"testing"

	"github.com/Trashed/gneat/mutation"
	"github.com/Trashed/gneat/rand"

	"github.com/Trashed/gneat/genetics"
)

func TestAddNewGene(t *testing.T) {
	t.Parallel()

	store := genetics.NewInnovationStore()

	in1 := store.CreateNode(genetics.NodeInput)
	in2 := store.CreateNode(genetics.NodeInput)
	out := store.CreateNode(genetics.NodeOutput)
	hidden := store.CreateNode(genetics.NodeHidden)

	// TODO: Handle error?
	gene1, _ := store.CreateGene(in1, out)
	gene2, _ := store.CreateGene(in2, hidden)
	gene3, _ := store.CreateGene(hidden, out)

	genome := &genetics.Genome{
		Id: 1,
		Nodes: map[uint]*genetics.Node{
			in1.Id:    in1,
			in2.Id:    in2,
			out.Id:    out,
			hidden.Id: hidden,
		},
		Genes: map[uint]*genetics.Gene{
			gene1.Innovation: gene1,
			gene2.Innovation: gene2,
			gene3.Innovation: gene3,
		},
	}

	gene4, err := mutation.AddNewGene(genome, store, mockRandItems([]uint{1, 4}))
	if err != nil {
		t.Fatalf("unexpected error: %v\n", err)
	}

	if gene4.InNode.Id != in1.Id && gene4.OutNode.Id != hidden.Id {
		t.Errorf("node IDs in new gene don't match expected values, expected in=%d and out=%d but got in=%d and out=%d\n", in1.Id, hidden.Id, gene4.InNode.Id, gene4.OutNode.Id)
	}
}

func mockRandItems(ids []uint) func(items rand.Listable) uint {
	indexList := ids
	currInternalIndex := 0

	return func(items rand.Listable) uint {
		val := indexList[currInternalIndex]
		currInternalIndex++

		return val
	}
}
