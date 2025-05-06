/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package genetics_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Trashed/gneat/genetics"
)

const (
	validFileContent string = `genomestart 1
// node, id, type
node 1 0  // bias node
node 2 1  // input node
node 3 1  // input node
node 4 2  // output node
// connection gene, innovation/id, weight, input id, output id, enabled/disabled
gene 1 1.0 1 4 1  // bias -> output
gene 2 1.0 2 4 1  // input 1 -> output
gene 3 1.0 3 4 1  // input 2 -> output
genomeend 1`
)

var (
	nodeStore map[int]*genetics.Node = make(map[int]*genetics.Node)
	connStore map[int]*genetics.Gene = make(map[int]*genetics.Gene)
)

func TestFromFile(t *testing.T) {
	t.Parallel()

	type args struct {
		path        string
		fileContent string
	}
	tests := []struct {
		name     string
		args     args
		expected *genetics.Genome
		wantErr  bool
	}{
		{
			name: "valid startgenome",
			args: args{path: "validstartgenome", fileContent: validFileContent},
			expected: &genetics.Genome{
				Id: 1,
				Nodes: createNodeMap([]*genetics.Node{
					createNode(genetics.NodeBias),
					createNode(genetics.NodeInput),
					createNode(genetics.NodeInput),
					createNode(genetics.NodeOutput),
				}),
				Genes: createGeneMap([]*genetics.Gene{
					createGene(nodeStore[1], nodeStore[4], 1.0, true),
					createGene(nodeStore[2], nodeStore[4], 1.0, true),
					createGene(nodeStore[3], nodeStore[4], 1.0, true),
				}),
			},
			wantErr: false,
		},
	}

	tempDir := t.TempDir()

	for _, tt := range tests {

		genomeFilePath := tempDir + "/" + tt.args.path

		prepareGenomeFile(t, genomeFilePath, tt.args.fileContent)

		t.Run(tt.name, func(t *testing.T) {
			actual, err := genetics.FromFile(genomeFilePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if actual.Id != tt.expected.Id {
				t.Fatalf("ids doesn't match: expected %d but got %d\n", tt.expected.Id, actual.Id)
			}

			if err := nodesMatch(actual.Nodes, tt.expected.Nodes); err != nil {
				t.Fatal(err.Error())
			}

			if err := genesMatch(actual.Genes, tt.expected.Genes); err != nil {
				t.Fatal(err.Error())
			}
		})
	}
}

func prepareGenomeFile(t *testing.T, path, content string) {
	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("failed to create file for genome file: %v\n", err)
	}

	_, err = f.WriteString(content)
	if err != nil {
		t.Fatalf("failed to write content to file: %v\n", err)
	}
	if err = f.Close(); err != nil {
		t.Fatalf("error while closing the genome file: %s\n", err)
	}
}

func createNode(nt genetics.NodeType) *genetics.Node {
	id := len(nodeStore)
	id++

	n := &genetics.Node{Id: uint(id), NodeType: nt}
	nodeStore[id] = n
	return n
}

func createGene(inNode, outNode *genetics.Node, weight float64, enabled bool) *genetics.Gene {
	id := len(connStore)
	id++

	g := &genetics.Gene{Innovation: uint(id), Weight: weight, InNode: inNode, OutNode: outNode, Enabled: enabled, Recurrent: false}
	connStore[id] = g
	return g
}

func nodesMatch(actualNodes, expectedNodes genetics.Nodes) error {

	actLen, expLen := len(actualNodes), len(expectedNodes)
	if actLen != expLen {
		return fmt.Errorf("number of nodes doesn't match - expected %d but got %d\n", expLen, actLen)
	}

	for i, actNode := range actualNodes {
		expNode := expectedNodes[i]
		if actNode.Id != expNode.Id && actNode.NodeType != expNode.NodeType {
			return fmt.Errorf("nodes doesn't match - expected node with id %d and type %d but got id %d and type %d\n", expNode.Id, expNode.NodeType, actNode.Id, actNode.NodeType)
		}
	}

	return nil
}

func genesMatch(actualGenes, expectedGenes map[uint]*genetics.Gene) error {

	actLen, expLen := len(actualGenes), len(expectedGenes)
	if actLen != expLen {
		return fmt.Errorf("number of genes doesn't match - expected %d but got %d\n", expLen, actLen)
	}

	for i, actGene := range actualGenes {
		expGene := expectedGenes[i]
		if actGene.Innovation != expGene.Innovation && actGene.Weight != expGene.Weight && actGene.InNode.Id != expGene.InNode.Id && actGene.OutNode.Id != expGene.OutNode.Id && actGene.Enabled != expGene.Enabled && actGene.Recurrent != expGene.Recurrent {
			return fmt.Errorf("genes don't match - expected gene %+v but got gene %+v\n", expGene, actGene)
		}
	}

	return nil
}

func createNodeMap(nodes []*genetics.Node) map[uint]*genetics.Node {
	nodeMap := make(map[uint]*genetics.Node)

	for _, node := range nodes {
		nodeMap[node.Id] = node
	}

	return nodeMap
}

func createGeneMap(genes []*genetics.Gene) map[uint]*genetics.Gene {
	geneMap := make(map[uint]*genetics.Gene)

	for _, g := range genes {
		geneMap[g.Innovation] = g
	}

	return geneMap
}
