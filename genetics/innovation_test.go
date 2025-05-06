/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package genetics_test

import (
	"testing"

	"github.com/Trashed/gneat/genetics"
)

func TestCreateNode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		nodeTypes []genetics.NodeType
		//expectedNodeCount int
	}{
		{
			name:      "1 input, 1 output",
			nodeTypes: []genetics.NodeType{genetics.NodeInput, genetics.NodeOutput},
			//expectedNodeCount: 2,
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
