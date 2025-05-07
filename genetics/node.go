/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package genetics

import (
	"slices"
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

type Nodes map[uint]*Node

func (ns Nodes) Len() int {
	return len(ns)
}

// List returns a list of node IDs.
func (ns Nodes) List() []uint {
	ids := make([]uint, 0)

	for id := range ns {
		ids = append(ids, id)
	}

	slices.SortFunc(ids, func(a uint, b uint) int {
		if a < b {
			return -1
		}

		if a > b {
			return 1
		}

		return 0
	})

	return ids
}

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
