/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package genetics

import "github.com/Trashed/gneat"

type NodeType uint

const (
	BIAS_NODE NodeType = iota
	INPUT_NODE
	HIDDEN_NODE
	OUTPUT_NODE
)

type NodeGene struct {
	id       uint
	nodeType NodeType
	trait    *Trait
}

func CreateNode(id uint, nodeType NodeType, trait *Trait) (*NodeGene, error) {

	if trait == nil {
		return nil, gneat.ErrNilTraitParameter
	}

	return &NodeGene{
		id:       id,
		nodeType: nodeType,
		trait:    trait,
	}, nil
}

func (n NodeGene) Trait() *Trait {
	return n.trait
}

func (n NodeGene) Id() uint {
	return n.id
}
