// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package gneat

// Link is a structure that connects two nodes together in the neural network.
//
// In the case of input layer nodes, those nodes only have outgoing connections.
type Link struct {
	Weight      float64
	In          *Node
	Out         *Node
	IsRecurrent bool
}

func NewLink(in, out *Node, weight float64, isRecurrent bool) *Link {
	return &Link{
		In:          in,
		Out:         out,
		Weight:      weight,
		IsRecurrent: isRecurrent,
	}
}
