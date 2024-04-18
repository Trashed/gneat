// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package neurals

type NeuralNodeType uint

const (
	NodeBias NeuralNodeType = iota
	NodeInput
	NodeOutput
	NodeHidden
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
func (n *Node) ConnectFrom(out *Node, weight float64) *Link {
	return &Link{
		In:     n,
		Out:    out,
		Weight: weight,
	}
}

type NodeMutator interface {
	MutateActivationFn(*Node, ActivationFunc)
}
