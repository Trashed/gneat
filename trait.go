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
	connectionEnabled     bool
	module                string // This is for tagging certain nodes or connections to a named group. Not to be mutated.
	// distanceMetric        DistanceMetricEnum
}
