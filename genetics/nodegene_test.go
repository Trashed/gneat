/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package genetics_test

import (
	"errors"
	"github.com/Trashed/gneat/genetics"

	// "errors"
	"github.com/Trashed/gneat"
	"math/rand/v2"
	"testing"
)

func TestCreateNodeGene(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		id       uint
		nodeType genetics.NodeType
		trait    *genetics.Trait
		err      error
	}{
		{
			name:     "create node with nil trait",
			id:       rand.Uint(),
			nodeType: genetics.INPUT_NODE,
			trait:    nil,
			err:      gneat.ErrNilTraitParameter,
		},
		{
			name:     "create valid node with trait",
			id:       rand.Uint(),
			nodeType: genetics.HIDDEN_NODE,
			trait:    genetics.CreateTrait(rand.Uint(), gneat.ACTIVATION_SIGMOID, 2.0, 1.0, 0.5, true, ""),
			err:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			node, actErr := genetics.CreateNode(test.id, test.nodeType, test.trait)

			// Expected error and actual error mismatch
			if test.err != nil && !errors.Is(actErr, test.err) {
				t.Errorf("%s got error: %v, expected error: %v", test.name, actErr, test.err)
			}

			if node == nil && actErr == nil {
				t.Errorf("%s - created node should not be nil", test.name)
			}
		})
	}
}
