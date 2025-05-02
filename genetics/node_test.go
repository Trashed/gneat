package genetics_test

import (
	"testing"

	"github.com/Trashed/gneat/genetics"
)

func TestNewNode(t *testing.T) {
	t.Parallel()

	node := genetics.NewNode(genetics.NodeInput)
	if node.Id != 1 {
		t.Errorf("Expected node ID to be 1, got %d", node.Id)
	}

	node2 := genetics.NewNode(genetics.NodeOutput)
	if node2.Id != 2 {
		t.Errorf("Expected node ID to be 2, got %d", node2.Id)
	}
}
