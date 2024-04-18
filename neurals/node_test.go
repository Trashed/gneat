// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package neurals_test

import (
	"fmt"
	"testing"

	"github.com/Trashed/gneat/activation"
	"github.com/Trashed/gneat/internal/assertions"
	"github.com/Trashed/gneat/mocks"
	"github.com/Trashed/gneat/neurals"
)

func TestCreateNode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nodeType     neurals.NeuralNodeType
		activationFn neurals.ActivationFunc
	}{
		{
			nodeType:     neurals.NodeInput,
			activationFn: activation.NonActivationFunc,
		},
		{
			nodeType:     neurals.NodeOutput,
			activationFn: activation.SigmoidFunc,
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

	n := neurals.NewNode(1, neurals.NodeHidden, activation.NonActivationFunc)

	m.MutateActivationFn(n, activation.SigmoidFunc)

	const fnInputVal = 0.41
	assertions.Equals(t, n.ActivationFn(fnInputVal), activation.SigmoidFunc(fnInputVal))
}

func TestConnectFromSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  *neurals.Node
		out *neurals.Node
		// isRecurrent bool
	}{
		{
			in:  neurals.NewNode(1, neurals.NodeInput, activation.NonActivationFunc),
			out: neurals.NewNode(2, neurals.NodeOutput, activation.SigmoidFunc),
		},
		{
			in:  neurals.NewNode(2, neurals.NodeHidden, activation.SigmoidFunc),
			out: neurals.NewNode(3, neurals.NodeOutput, activation.SigmoidFunc),
		},
	}

	for testNum, test := range tests {
		t.Run(fmt.Sprintf("TestConnectFrom_%d", testNum), func(t *testing.T) {
			link := test.in.ConnectFrom(test.out, 0.0)

			assertions.Assert(t, link != nil, "returned link shouldn't be nil")
			assertions.Equals(t, link.In, test.in)
			assertions.Equals(t, link.Out, test.out)
			// assertions.Equals(t, link.IsRecurrent, test.isRecurrent)
		})
	}
}

// TODO: TestConnectFromFailure
// - Don't connect output to output
// - Don't connect input to input
