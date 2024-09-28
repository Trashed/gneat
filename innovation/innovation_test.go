/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package innovation_test

import (
	"github.com/Trashed/gneat"
	"github.com/Trashed/gneat/genetics"
	"github.com/Trashed/gneat/innovation"
	"math/rand/v2"
	"testing"
)

func TestGetConnection(t *testing.T) {
	t.Parallel()

	innovationRecord := innovation.NewInnovationRecord()

	tests := []struct {
		name          string
		expInnovation uint
		nodeInId      uint
		nodeOutId     uint
	}{
		{
			name:          "add first connection",
			expInnovation: 1,
			nodeInId:      1,
			nodeOutId:     2,
		},
		{
			name:          "add second connection",
			expInnovation: 2,
			nodeInId:      2,
			nodeOutId:     2,
		},
		{
			name:          "add a similar connection",
			expInnovation: 2,
			nodeInId:      2,
			nodeOutId:     2,
		},
		{
			name:          "add third connection",
			expInnovation: 3,
			nodeInId:      2,
			nodeOutId:     1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			trait := genetics.CreateTrait(1, gneat.ACTIVATION_SIGMOID, rand.Float64(), rand.Float64(), rand.Float64(), true, "")
			nodeIn, _ := genetics.CreateNode(test.nodeInId, genetics.INPUT_NODE, trait)
			nodeOut, _ := genetics.CreateNode(test.nodeOutId, genetics.OUTPUT_NODE, trait)

			conn, err := innovationRecord.GetConnection(nodeIn, nodeOut)
			if err != nil {
				t.Errorf("error while getting connection for innovation")
			}

			if conn.Innovation() != test.expInnovation && innovationRecord.Current() != test.expInnovation {
				t.Errorf("got innovation = %d, expected = %d\n", conn.Innovation(), test.expInnovation)
			}
		})
	}

}
