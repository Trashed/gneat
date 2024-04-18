// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package gneat_test

import (
	"testing"

	"github.com/Trashed/gneat"
	"github.com/Trashed/gneat/activation"
	"github.com/Trashed/gneat/internal/assertions"
	"github.com/Trashed/gneat/mocks"
)

func TestCreateNode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nodeType     gneat.Nodetype
		activationFn gneat.ActivationFunc
	}{
		{
			nodeType:     gneat.NodeSensor,
			activationFn: activation.NonActivationFunc,
		},
		{
			nodeType:     gneat.NodeNeuron,
			activationFn: activation.NonActivationFunc,
		},
	}

	for i, test := range tests {
		node := gneat.NewNode(uint(i), test.nodeType, test.activationFn)

		assertions.Equals(t, test.nodeType, node.Type)
	}
}

func TestMutateActivationFn(t *testing.T) {
	t.Parallel()

	m := mocks.MockMutator{}

	n := gneat.NewNode(1, gneat.NodeNeuron, activation.NonActivationFunc)

	m.MutateActivationFn(n, activation.SigmoidFunc)

	const fnInputVal = 0.41
	assertions.Equals(t, n.ActivationFn(fnInputVal), activation.SigmoidFunc(fnInputVal))
}
