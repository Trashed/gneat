/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package genetics

var (
	idIncrementer func() uint = func() func() uint {
		id := uint(0)
		return func() uint {
			id++
			return id
		}
	}()
)

type NodeType uint

const (
	NodeBias NodeType = iota
	NodeInput
	NodeOutput
	NodeHidden
)

type Node struct {
	Id       uint
	NodeType NodeType
}

// NewNode creates a new node with a unique ID and the specified type.
// TODO: Pass the ID as an arument rather than creating it within NewNode.
// TODO: Refactor InnovationStore to include node creation and ID assignment?
func NewNode(nodeType NodeType) *Node {
	return &Node{
		Id:       idIncrementer(),
		NodeType: nodeType,
	}
}

type Nodes map[uint]*Node

func (ns Nodes) fetch(id uint) *Node {
	// TODO: Optimize this with binary search
	// TODO: Binary search expects the slice to be sorted -> implement logic where the Nodes slice is sorted each time a new node is pushed
	for _, n := range ns {
		if n.Id == id {
			return n
		}
	}

	return nil
}
