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

type Trait struct {
	id             uint
	activationFunc gneat.ActivationFuncEnum
	// Controls the range of weights that can be assigned to a connection gene. Some genes may allow a wider or narrower range of possible weight values.
	connectionWeightRange float64
	//  Represents the bias value for certain neurons (or even entire node groups). Instead of using a fixed bias (like 1), this trait allows the bias value to evolve over time.
	bias float64
	// Represents the strength or impact of recurrent connections, if the network includes recurrence (i.e., feedback loops).
	recurrentStrength float64
	// weightMutationRate float64 // TODO: Decide whether this would make a good trait property.
	connectionEnabled bool
	module            string // This is for tagging certain nodes or connections to a named group. Not to be mutated. Can be empty.
	// distanceMetric        DistanceMetricEnum
}

func CreateTrait(id uint, activationFuncName gneat.ActivationFuncEnum, connectionWeightRange float64, bias float64, recurrentStrength float64, enabled bool, module string) *Trait {
	return &Trait{
		id:                    id,
		activationFunc:        activationFuncName,
		connectionWeightRange: connectionWeightRange,
		bias:                  bias,
		recurrentStrength:     recurrentStrength,
		connectionEnabled:     enabled,
		module:                module,
	}
}
