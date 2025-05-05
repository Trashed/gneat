package genetics_test

import (
	"testing"

	"github.com/Trashed/gneat/genetics"
)

func TestPushGene(t *testing.T) {
	t.Parallel()

	const expNodes = 2
	const expGenes = 1
	const expInnovation = 1

	store := genetics.NewInnovationStore()
	gene := &genetics.Gene{
		InNode:    genetics.NewNode(genetics.NodeInput),
		OutNode:   genetics.NewNode(genetics.NodeOutput),
		Weight:    0.5,
		Enabled:   true,
		Recurrent: false,
	}

	if _, isAdded := store.AddGene(gene); !isAdded {
		t.Fatal("failed to add gene to innovation store")
	}

	g := &genetics.Genome{
		Id: 1,
	}
	g.PushGene(gene)

	if len(g.Genes) != expGenes {
		t.Errorf("expected %d genes, got %d", expGenes, len(g.Genes))
	}
	if len(g.Nodes) != expNodes {
		t.Errorf("expected %d nodes, got %d", expNodes, len(g.Nodes))
	}

	if _, ok := g.Genes[expInnovation]; !ok {
		t.Fatal("gene not found in genome")
	}
	/*if g.Genes[0].Innovation != expInnovation {
		t.Errorf("expected innovation %d, got %d", expInnovation, g.Genes[0].Innovation)
	}*/
}

func TestPushTwoGenes(t *testing.T) {
	t.Parallel()

	const expNodes = 3
	const expGenes = 2
	const expInnovation = 2

	store := genetics.NewInnovationStore()

	outNode := genetics.NewNode(genetics.NodeOutput)

	gene1 := &genetics.Gene{
		InNode:    genetics.NewNode(genetics.NodeInput),
		OutNode:   outNode,
		Weight:    0.5,
		Enabled:   true,
		Recurrent: false,
	}
	gene2 := &genetics.Gene{
		InNode:    genetics.NewNode(genetics.NodeInput),
		OutNode:   outNode,
		Weight:    0.5,
		Enabled:   true,
		Recurrent: false,
	}

	store.AddGene(gene1)
	store.AddGene(gene2)

	if expInnovation != store.Innovation() {
		t.Errorf("expected innovation %d, got %d", expInnovation, store.Innovation())
	}

	g := &genetics.Genome{
		Id: 1,
	}
	g.PushGene(gene1)
	g.PushGene(gene2)

	if len(g.Genes) != expGenes {
		t.Errorf("expected %d genes, got %d", expGenes, len(g.Genes))
	}
	if len(g.Nodes) != expNodes {
		t.Errorf("expected %d nodes, got %d", expNodes, len(g.Nodes))
	}
}
