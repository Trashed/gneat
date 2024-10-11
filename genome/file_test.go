/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package genome_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Trashed/gneat/genome"
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
	nodeStore map[int]*genome.Node = make(map[int]*genome.Node)
	connStore map[int]*genome.Gene = make(map[int]*genome.Gene)
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
		expected *genome.Genome
		wantErr  bool
	}{
		{
			name: "valid startgenome",
			args: args{path: "validstartgenome", fileContent: validFileContent},
			expected: &genome.Genome{
				Id: 1,
				Nodes: []*genome.Node{
					createNode(genome.NodeBias),
					createNode(genome.NodeInput),
					createNode(genome.NodeInput),
					createNode(genome.NodeOutput),
				},
				Genes: []*genome.Gene{
					createGene(nodeStore[1], nodeStore[4], 1.0, true),
					createGene(nodeStore[2], nodeStore[4], 1.0, true),
					createGene(nodeStore[3], nodeStore[4], 1.0, true),
				},
			},
			wantErr: false,
		},
	}

	tempDir := t.TempDir()

	for _, tt := range tests {

		genomeFilePath := tempDir + "/" + tt.args.path

		prepareGenomeFile(t, genomeFilePath, tt.args.fileContent)

		t.Run(tt.name, func(t *testing.T) {
			actual, err := genome.FromFile(genomeFilePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			/*if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("FromFile() = %v, want %v", actual, tt.expected)
			}*/
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

func createNode(nt genome.NodeType) *genome.Node {
	id := len(nodeStore)
	id++

	n := &genome.Node{Id: uint(id), NodeType: nt}
	nodeStore[id] = n
	return n
}

func createGene(inNode, outNode *genome.Node, weight float64, enabled bool) *genome.Gene {
	id := len(connStore)
	id++

	g := &genome.Gene{Innovation: uint(id), Weight: weight, InNode: inNode, OutNode: outNode, Enabled: enabled, Recurrent: false}
	return g
}

func nodesMatch(actualNodes, expectedNodes genome.Nodes) error {

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

func genesMatch(actualGenes, expectedGenes []*genome.Gene) error {

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
