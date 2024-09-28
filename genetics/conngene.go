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

type ConnGene struct {
	nodeIn     *NodeGene
	nodeOut    *NodeGene
	trait      *Trait
	weight     float64
	innovation uint
	// enabled bool		// TODO: This is already specified within the Trait, so this isn't probably needed here.
}

func CreateConnGene(nodeIn *NodeGene, nodeOut *NodeGene, trait *Trait, weight float64, innovation uint) (*ConnGene, error) {

	if trait == nil {
		return nil, gneat.ErrNilTraitParameter
	}

	if nodeIn == nil || nodeOut == nil {
		return nil, gneat.ErrNilConnGeneNodeParameter
	}

	return &ConnGene{
		nodeIn:     nodeIn,
		nodeOut:    nodeOut,
		trait:      trait,
		weight:     weight,
		innovation: innovation,
	}, nil
}

func (g ConnGene) Innovation() uint {
	return g.innovation
}
