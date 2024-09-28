/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gneat

import "errors"

var (
	ErrNilTraitParameter           = errors.New("trait can't be nil")
	ErrNilConnGeneNodeParameter    = errors.New("connection input or output node can't be nil")
	ErrInnovationNumberLenMismatch = errors.New("innovation number doesn't match with the length of the record")
)
