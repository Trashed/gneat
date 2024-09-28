/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package innovation

import (
	"github.com/Trashed/gneat"
	"github.com/Trashed/gneat/genetics"
	"math/rand"
	"strconv"
	"strings"
)

// Record holds information about the unique connections within NEAT. Every time a new input-output node pair (or connection)
// is added, the innovation number is incremented by one. When an existing connection is introduced during evolution process,
// a [genetics.ConnGene] with an existing innovation number - that is checked against the record of connection - is returned.
type Record struct {
	innovation  uint
	connections map[string]uint
}

// NewInnovationRecord returns a new instance of Record.
func NewInnovationRecord() *Record {
	return &Record{
		innovation:  0,
		connections: make(map[string]uint),
	}
}

// GetConnection get an input and output node as parameters and checks whether a connection with that node pair already exists.
// If the connection is new, the innovation number is incremented by one, the node pair is added to the record and a new
// connection is returned. If the connection exists, the connection with existing innovation number matching the input-output
// node pair is returned.
func (r *Record) GetConnection(in *genetics.NodeGene, out *genetics.NodeGene) (*genetics.ConnGene, error) {
	sb := strings.Builder{}
	inNodeIdStr := strconv.FormatUint(uint64(in.Id()), 10)
	outNodeIdStr := strconv.FormatUint(uint64(out.Id()), 10)
	sb.WriteString(inNodeIdStr)
	sb.WriteString("->")
	sb.WriteString(outNodeIdStr)

	recordKey := sb.String()

	if innov, ok := r.connections[recordKey]; ok {
		return genetics.CreateConnGene(in, out, in.Trait(), rand.Float64(), innov)
	}

	r.innovation++
	r.connections[recordKey] = r.innovation

	if len(r.connections) != int(r.innovation) {
		return nil, gneat.ErrInnovationNumberLenMismatch
	}

	conn, err := genetics.CreateConnGene(in, out, in.Trait(), rand.Float64(), r.innovation)
	return conn, err
}

// Current returns the current innovation number.
func (r *Record) Current() uint {
	return r.innovation
}
