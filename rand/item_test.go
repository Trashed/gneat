package rand_test

import (
	"fmt"
	"testing"

	"github.com/Trashed/gneat/genetics"
	"github.com/Trashed/gneat/rand"
)

func TestItem(t *testing.T) {
	t.Parallel()

	const count = 1200

	store := genetics.NewInnovationStore()
	nodes := make(genetics.Nodes)

	// Create a large number of input, hidden and output nodes
	for i := range count {

		var n *genetics.Node
		if i < count/3 {
			n = store.CreateNode(genetics.NodeInput)
		} else if i > count/3 && i < (count/3)*2 {
			n = store.CreateNode(genetics.NodeHidden)
		} else {
			n = store.CreateNode(genetics.NodeOutput)
		}

		nodes[n.Id] = n
	}

	// Verify that the random function works by running it multiple times (meaning, it shouldn't fail or panic)
	for range 1000 {
		randVal := rand.Item(nodes)
		fmt.Printf("random value = %d\n", randVal)
	}
}
