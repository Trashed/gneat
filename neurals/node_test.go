// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package neurals_test

import (
	"testing"

	"github.com/Trashed/gneat/activation"
	"github.com/Trashed/gneat/internal/assertions"
	"github.com/Trashed/gneat/mocks"
	"github.com/Trashed/gneat/neurals"
)

func TestCreateNode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nodeType     neurals.Nodetype
		activationFn neurals.ActivationFunc
	}{
		{
			nodeType:     neurals.NodeSensor,
			activationFn: activation.NonActivationFunc,
		},
		{
			nodeType:     neurals.NodeNeuron,
			activationFn: activation.NonActivationFunc,
		},
	}

	for i, test := range tests {
		node := neurals.NewNode(uint(i), test.nodeType, test.activationFn)

		assertions.Equals(t, test.nodeType, node.Type)
	}
}

func TestMutateActivationFn(t *testing.T) {
	t.Parallel()

	m := mocks.MockMutator{}

	n := neurals.NewNode(1, neurals.NodeNeuron, activation.NonActivationFunc)

	m.MutateActivationFn(n, activation.SigmoidFunc)

	const fnInputVal = 0.41
	assertions.Equals(t, n.ActivationFn(fnInputVal), activation.SigmoidFunc(fnInputVal))
}
