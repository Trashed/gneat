/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gneat

type Trait struct {
	id                    uint
	activationFunc        ActivationFuncEnum
	connectionWeightRange float64
	bias                  float64
	recurrentStrength     float64
	// weightMutationRate float64 // TODO: Decide whether this would make a good trait property.
	connectionEnabled bool
	module            string // This is for tagging certain nodes or connections to a named group. Not to be mutated. Can be empty.
	// distanceMetric        DistanceMetricEnum
}

func CreateTrait(id uint, activationFuncName ActivationFuncEnum, connectionWeightRange float64, bias float64, recurrentStrength float64, enabled bool, module string) *Trait {
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
