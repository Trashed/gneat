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
	Trait    *NodeTrait
}
