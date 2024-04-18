// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package activation

import (
	"math"

	"github.com/Trashed/gneat/neurals"
)

var (
	NonActivationFunc neurals.ActivationFunc = nonActivationFunc
	SigmoidFunc       neurals.ActivationFunc = sigmoidFunc
)

func nonActivationFunc(x float64) float64 {
	return x
}

func sigmoidFunc(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}
