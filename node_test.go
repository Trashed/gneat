// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package gneat_test

import (
	"fmt"
	"testing"

	"github.com/Trashed/gneat"
	"github.com/Trashed/gneat/activation"
	"github.com/Trashed/gneat/internal/assertions"
	"github.com/Trashed/gneat/mocks"
)

func TestCreateNode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nodeType     gneat.NeuralNodeType
		activationFn gneat.ActivationFunc
	}{
		{
			nodeType:     gneat.NodeInput,
			activationFn: activation.NonActivationFunc,
		},
		{
			nodeType:     gneat.NodeOutput,
			activationFn: activation.SigmoidFunc,
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

	n := gneat.NewNode(1, gneat.NodeHidden, activation.NonActivationFunc)

	m.MutateActivationFn(n, activation.SigmoidFunc)

	const fnInputVal = 0.41
	assertions.Equals(t, n.ActivationFn(fnInputVal), activation.SigmoidFunc(fnInputVal))
}

func TestConnectFromSuccess(t *testing.T) {
	t.Parallel()

	inNode := gneat.NewNode(1, gneat.NodeInput, activation.NonActivationFunc)
	outNode := gneat.NewNode(2, gneat.NodeOutput, activation.SigmoidFunc)
	hiddenNode := gneat.NewNode(4, gneat.NodeHidden, activation.SigmoidFunc)

	tests := []struct {
		in                *gneat.Node
		out               *gneat.Node
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

	inNode := gneat.NewNode(1, gneat.NodeInput, activation.NonActivationFunc)
	outNode := gneat.NewNode(2, gneat.NodeOutput, activation.SigmoidFunc)
	hiddenNode := gneat.NewNode(4, gneat.NodeHidden, activation.SigmoidFunc)

	tests := []struct {
		in     *gneat.Node
		out    *gneat.Node
		expErr error
	}{
		// recurrent link to the input node itself
		{
			in:     inNode,
			out:    inNode,
			expErr: gneat.ErrInvalidConnection,
		},
		// link to other input node
		{
			in:     inNode,
			out:    gneat.NewNode(2, gneat.NodeInput, activation.NonActivationFunc),
			expErr: gneat.ErrInvalidConnection,
		},
		// link from hidden node to input
		{
			in:     hiddenNode,
			out:    inNode,
			expErr: gneat.ErrInvalidConnection,
		},
		// link from output to input
		{
			in:     outNode,
			out:    inNode,
			expErr: gneat.ErrInvalidConnection,
		},
		{ // link from output to other output
			in:     outNode,
			out:    gneat.NewNode(4, gneat.NodeOutput, activation.SigmoidFunc),
			expErr: gneat.ErrInvalidConnection,
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
