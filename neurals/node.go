// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package neurals

type Nodetype uint

const (
	// Input nodes
	NodeSensor Nodetype = iota
	// Hidden and output nodes
	NodeNeuron
)

type ActivationFunc func(float64) float64

func NewNode(id uint, nodeType Nodetype, fn ActivationFunc) *Node {
	return &Node{
		Id:           id,
		Type:         nodeType,
		ActivationFn: fn,
	}
}

type Node struct {
	Id              uint
	Type            Nodetype
	ActivationFn    ActivationFunc
	ActivationValue float64
}

type NodeMutator interface {
	MutateActivationFn(*Node, ActivationFunc)
}
