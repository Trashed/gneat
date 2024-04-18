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

	inNode := neurals.NewNode(1, neurals.NodeInput, activation.NonActivationFunc)
	outNode := neurals.NewNode(2, neurals.NodeOutput, activation.SigmoidFunc)
	hiddenNode := neurals.NewNode(4, neurals.NodeHidden, activation.SigmoidFunc)

	tests := []struct {
		in                *neurals.Node
		out               *neurals.Node
		shouldBeRecurrent bool
	}{
		{
			in:  inNode,
			out: outNode,
		},
		{
			in:  hiddenNode,
			out: outNode,
		},
		{ // link back to the output node itself
			in:                outNode,
			out:               outNode,
			shouldBeRecurrent: true,
		},
		{
			in:                outNode,
			out:               hiddenNode,
			shouldBeRecurrent: true,
		},
		{ // link back to the hiddne node itself
			in:                hiddenNode,
			out:               hiddenNode,
			shouldBeRecurrent: true,
		},
	}

	for testNum, test := range tests {
		t.Run(fmt.Sprintf("TestConnectFromSuccess_%d", testNum), func(t *testing.T) {
			link, err := test.in.ConnectFrom(test.out, 0.0, test.shouldBeRecurrent)

			assertions.Ok(t, err)
			assertions.Assert(t, link != nil, "returned link shouldn't be nil")
			assertions.Equals(t, link.In, test.in)
			assertions.Equals(t, link.Out, test.out)
			assertions.Equals(t, link.IsRecurrent, test.shouldBeRecurrent)
		})
	}
}

func TestConnectFromFailure(t *testing.T) {
	t.Parallel()

	inNode := neurals.NewNode(1, neurals.NodeInput, activation.NonActivationFunc)
	outNode := neurals.NewNode(2, neurals.NodeOutput, activation.SigmoidFunc)
	hiddenNode := neurals.NewNode(4, neurals.NodeHidden, activation.SigmoidFunc)

	tests := []struct {
		in     *neurals.Node
		out    *neurals.Node
		expErr error
	}{
		// recurrent link to the input node itself
		{
			in:     inNode,
			out:    inNode,
			expErr: neurals.ErrInvalidConnection,
		},
		// link to other input node
		{
			in:     inNode,
			out:    neurals.NewNode(2, neurals.NodeInput, activation.NonActivationFunc),
			expErr: neurals.ErrInvalidConnection,
		},
		// link from hidden node to input
		{
			in:     hiddenNode,
			out:    inNode,
			expErr: neurals.ErrInvalidConnection,
		},
		// link from output to input
		{
			in:     outNode,
			out:    inNode,
			expErr: neurals.ErrInvalidConnection,
		},
		{ // link from output to other output
			in:     outNode,
			out:    neurals.NewNode(4, neurals.NodeOutput, activation.SigmoidFunc),
			expErr: neurals.ErrInvalidConnection,
		},
	}

	for testNum, test := range tests {
		t.Run(fmt.Sprintf("TestConnectFromFailure_%d", testNum), func(t *testing.T) {
			_, err := test.in.ConnectFrom(test.out, 0.0, true)

			assertions.Assert(t, err != nil, "returned error should't be nil")
			assertions.Equals(t, err, test.expErr)
		})
	}
}
