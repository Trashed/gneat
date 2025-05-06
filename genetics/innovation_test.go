/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package genetics_test

import (
	"errors"
	"testing"

	"github.com/Trashed/gneat/genetics"
)

func TestCreateNode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		nodeTypes []genetics.NodeType
	}{
		{
			name:      "1 input, 1 output",
			nodeTypes: []genetics.NodeType{genetics.NodeInput, genetics.NodeOutput},
		},
		{
			name:      "2 inputs, 1 output",
			nodeTypes: []genetics.NodeType{genetics.NodeInput, genetics.NodeInput, genetics.NodeOutput},
		},
		{
			name:      "5 inputs, 2 outputs",
			nodeTypes: []genetics.NodeType{genetics.NodeInput, genetics.NodeInput, genetics.NodeInput, genetics.NodeInput, genetics.NodeInput, genetics.NodeOutput, genetics.NodeOutput},
		},
		{
			name:      "3 inputs, 2 hidden nodes, 2 outputs",
			nodeTypes: []genetics.NodeType{genetics.NodeInput, genetics.NodeInput, genetics.NodeInput, genetics.NodeHidden, genetics.NodeHidden, genetics.NodeOutput, genetics.NodeOutput},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := genetics.NewInnovationStore()

			for i, nt := range test.nodeTypes {
				expId := uint(i + 1)
				n := store.CreateNode(nt)

				if n == nil {
					t.Fatalf("node creation failed, it shouldn't be nil")
				}

				if n.NodeType != nt {
					t.Errorf("expected node type to be %v but got %v\n", nt, n.NodeType)
				}

				if n.Id != expId {
					t.Errorf("expected node id to be %d but got %d\n", expId, n.Id)
				}
			}

		})
	}
}

func TestCreateGene(t *testing.T) {
	t.Parallel()

	store := genetics.NewInnovationStore()

	tests := []struct {
		name          string
		in            *genetics.Node
		out           *genetics.Node
		expInnovation uint
		expErr        error
	}{
		{
			name:   "input node is nil",
			in:     nil,
			out:    &genetics.Node{Id: 2, NodeType: genetics.NodeOutput},
			expErr: genetics.ErrNilNode,
		},
		{
			name:   "output node is nil",
			in:     &genetics.Node{Id: 1, NodeType: genetics.NodeInput},
			out:    nil,
			expErr: genetics.ErrNilNode,
		},
		{
			name:          "gene created successfully",
			in:            store.CreateNode(genetics.NodeInput),
			out:           store.CreateNode(genetics.NodeOutput),
			expInnovation: 1,
			expErr:        nil,
		},
		{
			name:          "existing gene used successfully",
			in:            store.GetNode(1),
			out:           store.GetNode(2),
			expInnovation: 1,
			expErr:        nil,
		},
		{
			name:          "new gene created successfully",
			in:            store.CreateNode(genetics.NodeInput),
			out:           store.CreateNode(genetics.NodeOutput),
			expInnovation: 2,
			expErr:        nil,
		},
		{
			name:          "new gene with hidden node created successfully",
			in:            store.CreateNode(genetics.NodeInput),
			out:           store.CreateNode(genetics.NodeHidden),
			expInnovation: 3,
			expErr:        nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g, err := store.CreateGene(test.in, test.out)

			if err != nil && !errors.Is(err, test.expErr) {
				t.Errorf("expected err of \"%v\" but got \"%v\"\n", test.expErr, err)
			} else if g != nil && err == nil {
				if g.InNode == nil || g.OutNode == nil {
					t.Error("gene created but input and/or output nodes are nil")
				}
				if g.Innovation != test.expInnovation {
					t.Errorf("gene created but innovation numbers don't match, expected %d but got %d\n", test.expInnovation, g.Innovation)
				}
			} else if g == nil && test.expErr == nil && err != nil {
				t.Fatalf("unexpected failure with error: %v\n", err)
			}
		})
	}
}

func TestAddGene(t *testing.T) {
	t.Parallel()

	store := genetics.NewInnovationStore()

	tests := []struct {
		gene        *genetics.Gene
		expectedNum uint
		expectedAdd bool
	}{
		{
			gene: &genetics.Gene{
				Weight:    0.5,
				InNode:    &genetics.Node{Id: 1},
				OutNode:   &genetics.Node{Id: 2},
				Enabled:   true,
				Recurrent: false,
			},
			expectedNum: 1,
			expectedAdd: true,
		},
		{
			gene: &genetics.Gene{
				Weight:    0.5,
				InNode:    &genetics.Node{Id: 1},
				OutNode:   &genetics.Node{Id: 2},
				Enabled:   true,
				Recurrent: false,
			},
			expectedNum: 1,
			expectedAdd: false,
		},
		{
			gene: &genetics.Gene{
				Weight:    0.5,
				InNode:    &genetics.Node{Id: 2},
				OutNode:   &genetics.Node{Id: 3},
				Enabled:   true,
				Recurrent: false,
			},
			expectedNum: 2,
			expectedAdd: true,
		},
		{
			gene: &genetics.Gene{
				Weight:    0.5,
				InNode:    &genetics.Node{Id: 2},
				OutNode:   &genetics.Node{Id: 3},
				Enabled:   true,
				Recurrent: false,
			},
			expectedNum: 2,
			expectedAdd: false,
		},
	}

	for _, test := range tests {
		t.Run("PushGene", func(t *testing.T) {
			innovationNum, isAdded := store.PushGene(test.gene)
			if innovationNum != test.expectedNum || isAdded != test.expectedAdd {
				t.Fatalf("Expected innovation number %d, got %d, isAdded: %v", test.expectedNum, innovationNum, isAdded)
			}

			if test.gene.Innovation != test.expectedNum {
				t.Fatalf("Expected gene innovation number %d, got %d", test.expectedNum, test.gene.Innovation)
			}
		})
	}
}
