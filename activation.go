/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gneat

type ActivationFuncEnum uint

const (
	ACTIVATION_NONE ActivationFuncEnum = iota
	ACTIVATION_SIGMOID
	ACTIVATION_TANH
	ACTIVATION_RELU
	ACTIVATION_LEAKY_RELU
	ACTIVATION_LINEAR
)
