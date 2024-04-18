// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package neurals

import "errors"

var (
	ErrInvalidConnection error = errors.New("link isn't supported")
)

type NeuralNodeType uint

// NeuralNodeTypes that can be assigned to a node.
const (
	NodeBias NeuralNodeType = iota
	NodeInput
	NodeHidden
	NodeOutput
)

type ActivationFunc func(float64) float64

func NewNode(id uint, nodeType NeuralNodeType, fn ActivationFunc) *Node {
	return &Node{
		Id:           id,
		Type:         nodeType,
		ActivationFn: fn,
	}
}

type Node struct {
	Id              uint
	Type            NeuralNodeType
	ActivationFn    ActivationFunc
	ActivationValue float64
}

// ConnectFrom connects a node with a new [Link] from this node forward;
// this node is the starting point, the other node is on the other side
// of the new link. Returns the newly created [Link].
func (n *Node) ConnectFrom(out *Node, weight float64, isRecurrent bool) (*Link, error) {

	if isInvalidLink(n, out) {
		return nil, ErrInvalidConnection
	}

	return &Link{
		In:          n,
		Out:         out,
		Weight:      weight,
		IsRecurrent: isRecurrent,
	}, nil
}

func isInvalidLink(in, out *Node) bool {

	switch in.Type {
	case NodeInput:
		if in.Type == out.Type {
			return true
		}
	case NodeHidden:
		if in.Type > out.Type {
			return true
		}
	case NodeOutput:
		if out.Type == NodeInput {
			return true
		}
		if out.Type == NodeOutput && in.Id != out.Id {
			return true
		}
	}

	return false
}

/*func isRecurrent(in, out *Node) bool {

	if (in.Type == NodeOutput || in.Type == NodeHidden) &&
		in.Type == out.Type &&
		in.Id == out.Id {
		return true
	}

	if in.Type != NodeInput && out.Type != NodeInput &&
		in.Type > out.Type {
		return true
	}

	return false
}*/

type NodeMutator interface {
	MutateActivationFn(*Node, ActivationFunc)
}
