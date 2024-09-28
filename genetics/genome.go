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

import "errors"

type Nodes []*NodeGene

func (n Nodes) Len() int {
	return len(n)
}

// Genome (or genotype) is the blueprint from which the neural network is constructed.
type Genome struct {
	id          uint
	nodes       []*NodeGene
	connections []*ConnGene
}

func (g Genome) Inputs() Nodes {
	return nil
}

// CreateInitialGenome creates a new genome. The new genome gets an id, number of inputs and output nodes as parameters.
func CreateInitialGenome(id uint, inputs int, outputs int /*fullyConnected bool*/) (*Genome, error) {

	// TODO: add a bias node and connections to outputs manually
	// TODO: create nodes for inputs and outputs

	return nil, errors.New("not yet implemented")
}
