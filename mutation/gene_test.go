package mutation_test

import (
	"testing"

	"github.com/Trashed/gneat/genetics"
	"github.com/Trashed/gneat/mutation"
)

func TestAddNewNode(t *testing.T) {
	t.Parallel()

	// Create a new gene
	gene := &genetics.Gene{
		InNode:    genetics.NewNode(genetics.NodeInput),
		OutNode:   genetics.NewNode(genetics.NodeOutput),
		Weight:    0.5,
		Enabled:   true,
		Recurrent: false,
	}

	const newNodeId = 3

	store := genetics.NewInnovationStore()
	if _, isAdded := store.PushGene(gene); !isAdded {
		t.Error("Expected to add gene to innovation store, but it was not added")
	}

	// Add a new node to the gene
	newGenes, newNode := mutation.AddNewNode(gene, store)

	if newGenes == nil {
		t.Error("Expected a new gene, got nil")
	}

	if newNode == nil {
		t.Error("Expected a new node, got nil")
	}

	if newGenes[0].OutNode.Id != newNodeId {
		t.Errorf("Expected new gene's OutNode ID to be %d, got %d", newNodeId, newGenes[0].OutNode.Id)
	}

	if newGenes[1].InNode.Id != newNodeId {
		t.Errorf("Expected new gene's InNode ID to be %d, got %d", newNodeId, newGenes[1].InNode.Id)
	}
}
