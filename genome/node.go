/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package genome

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

type Nodes []*Node

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
