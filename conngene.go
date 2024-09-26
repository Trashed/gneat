/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gneat

type ConnGene struct {
	id      uint
	nodeIn  uint
	nodeOut uint
	trait   *Trait
	weight  float64
	// enabled bool		// TODO: This is already specified withing the Trait, so this isn't probably needed here.
}
