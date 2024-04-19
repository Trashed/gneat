// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package mocks

import "github.com/Trashed/gneat"

type MockMutator struct {
}

func (mm MockMutator) MutateActivationFn(n *gneat.Node, fn gneat.ActivationFunc) {
	n.ActivationFn = fn
}
