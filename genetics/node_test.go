package genetics_test

import (
	"testing"

	"github.com/Trashed/gneat/genetics"
)

func TestNewNode(t *testing.T) {
	t.Parallel()

	store := genetics.NewInnovationStore()

	node := store.CreateNode(genetics.NodeInput)
	if node.Id != 1 {
		t.Errorf("Expected node ID to be 1, got %d", node.Id)
	}

	node2 := store.CreateNode(genetics.NodeOutput)
	if node2.Id != 2 {
		t.Errorf("Expected node ID to be 2, got %d", node2.Id)
	}
}
