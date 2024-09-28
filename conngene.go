/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gneat

type ConnGene struct {
	id         uint
	nodeIn     *NodeGene
	nodeOut    *NodeGene
	trait      *Trait
	weight     float64
	innovation uint
	// enabled bool		// TODO: This is already specified within the Trait, so this isn't probably needed here.
}

func CreateConnGene(id uint, nodeIn *NodeGene, nodeOut *NodeGene, trait *Trait, weight float64, innovation uint) (*ConnGene, error) {

	if trait == nil {
		return nil, ErrNilTraitParameter
	}

	if nodeIn == nil || nodeOut == nil {
		return nil, ErrNilConnGeneNodeParameter
	}

	return &ConnGene{
		id:         id,
		nodeIn:     nodeIn,
		nodeOut:    nodeOut,
		trait:      trait,
		weight:     weight,
		innovation: innovation,
	}, nil
}
