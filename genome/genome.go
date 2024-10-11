/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package genome

import (
	"slices"
)

type Genome struct {
	Id         uint
	NodeTraits []*NodeTrait
	Nodes      []*Node
}

func getNodeTrait(nodeTraits []*NodeTrait, id uint) *NodeTrait {

	if i, found := slices.BinarySearchFunc(nodeTraits, &NodeTrait{Id: id}, func(a, b *NodeTrait) int {
		if a.Id > b.Id {
			return 1
		} else if a.Id < b.Id {
			return -1
		}
		return 0
	}); !found {
		return nil
	} else {
		return nodeTraits[i]
	}
}
