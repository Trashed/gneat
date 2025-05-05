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
		t.Run("AddGene", func(t *testing.T) {
			innovationNum, isAdded := store.AddGene(test.gene)
			if innovationNum != test.expectedNum || isAdded != test.expectedAdd {
				t.Fatalf("Expected innovation number %d, got %d, isAdded: %v", test.expectedNum, innovationNum, isAdded)
			}

			if test.gene.Innovation != test.expectedNum {
				t.Fatalf("Expected gene innovation number %d, got %d", test.expectedNum, test.gene.Innovation)
			}
		})
	}
}
