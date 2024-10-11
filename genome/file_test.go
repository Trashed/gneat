/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package genome_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/Trashed/gneat/activation"
	"github.com/Trashed/gneat/genome"
)

const (
	validFileContent string = `genomestart 1
// nodetrait, id, activation func, bias value
nodetrait 1 0 1.0   // node trait for bias node
nodetrait 2 0 0.0   // node trait for input nodes
nodetrait 3 1 0.0   // node trait for output node; has a SIGMOID activation function marked by value 1
// node, id, type, trait
node 1 0 1      // bias node
node 2 1 2      // input node
node 3 1 2      // input node
node 4 2 3      // output node
// connection trait, id, enabled, recurrent strength, connection weight range, module
conntrait 1 1 0.0 1.0 1.0
conntrait 2 1 0.0 1.0 2.0
// connection gene, innovation/id, weight, input id, output id, connection trait
conngene 1 1.0 1 4 1    // bias -> output
conngene 2 1.0 2 4 2    // input 1 -> output
conngene 3 1.0 3 4 2    // input 2 -> output
genomeend 1`
)

func TestFromFile(t *testing.T) {
	t.Parallel()

	type args struct {
		path        string
		fileContent string
	}
	tests := []struct {
		name    string
		args    args
		want    *genome.Genome
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid startgenome",
			args: args{path: "validstartgenome", fileContent: validFileContent},
			want: &genome.Genome{
				Id: 1,
				NodeTraits: []*genome.NodeTrait{
					{Id: 1, ActivationFunc: activation.ActivationNone, Bias: 1.0},
					{Id: 2, ActivationFunc: activation.ActivationNone, Bias: 0.0},
					{Id: 3, ActivationFunc: activation.ActivationSigmoid, Bias: 0.0},
				},
				Nodes: []*genome.Node{
					{Id: 1, NodeType: genome.NodeBias, Trait: &genome.NodeTrait{Id: 1, ActivationFunc: activation.ActivationNone, Bias: 1.0}},
					{Id: 2, NodeType: genome.NodeInput, Trait: &genome.NodeTrait{Id: 2, ActivationFunc: activation.ActivationNone, Bias: 0.0}},
					{Id: 3, NodeType: genome.NodeInput, Trait: &genome.NodeTrait{Id: 2, ActivationFunc: activation.ActivationNone, Bias: 0.0}},
					{Id: 4, NodeType: genome.NodeOutput, Trait: &genome.NodeTrait{Id: 3, ActivationFunc: activation.ActivationSigmoid, Bias: 0.0}},
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
			got, err := genome.FromFile(genomeFilePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromFile() = %v, want %v", got, tt.want)
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
