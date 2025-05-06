package mutation_test

import (
	"testing"

	"github.com/Trashed/gneat/genetics"
	"github.com/Trashed/gneat/mutation"
)

func TestAddNewNode(t *testing.T) {
	t.Parallel()

	const newNodeId = 3
	store := genetics.NewInnovationStore()

	gene, err := store.CreateGene(store.CreateNode(genetics.NodeInput), store.CreateNode(genetics.NodeOutput))

	if err != nil {
		t.Fatalf("gene creation should succeed, but got error: %v\n", err)
	}

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
