/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gneat

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
