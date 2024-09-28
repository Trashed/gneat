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

package genetics_test

import (
	"errors"
	"github.com/Trashed/gneat"
	"github.com/Trashed/gneat/genetics"
	"math/rand/v2"
	"testing"
)

func TestCreateConnGeneWithNilTrait(t *testing.T) {
	t.Parallel()

	var trait *genetics.Trait = nil
	expErr := gneat.ErrNilTraitParameter

	conn, err := genetics.CreateConnGene(rand.Uint(), nil, nil, trait, rand.Float64(), rand.Uint())
	if conn == nil && err == nil {
		t.Errorf("CreateConnGene shouldn't return both conn and err as nil")
	}
	if conn == nil && err != nil {
		if !errors.Is(err, expErr) {
			t.Errorf("error = %v, expected err: %v", err, expErr)
		}
	}
}

func TestCreateConnGeneWithNilNodes(t *testing.T) {
	t.Parallel()

	trait := genetics.CreateTrait(rand.Uint(), gneat.ACTIVATION_SIGMOID, rand.NormFloat64(), rand.NormFloat64(), rand.Float64(), true, "")
	var nodeIn *genetics.NodeGene = nil
	var nodeOut *genetics.NodeGene = nil

	expErr := gneat.ErrNilConnGeneNodeParameter

	_, err := genetics.CreateConnGene(rand.Uint(), nodeIn, nodeOut, trait, rand.Float64(), rand.Uint())
	if !errors.Is(err, expErr) {
		t.Errorf("error = %v, expected err: %v", err, expErr)
	}
}

func TestCreateConnGene(t *testing.T) {
	t.Parallel()

	trait := genetics.CreateTrait(rand.Uint(), gneat.ACTIVATION_SIGMOID, rand.NormFloat64(), rand.NormFloat64(), rand.Float64(), true, "")
	nodeIn, _ := genetics.CreateNode(1, genetics.INPUT_NODE, trait)
	nodeOut, _ := genetics.CreateNode(2, genetics.OUTPUT_NODE, trait)

	conn, err := genetics.CreateConnGene(1, nodeIn, nodeOut, trait, rand.Float64(), rand.Uint())
	if conn == nil && err == nil {
		t.Errorf("connection shouldn't be nil")
	}
}
