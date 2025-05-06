/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gneat

import (
	"errors"
	"log"

	"github.com/Trashed/gneat/genetics"
	"github.com/Trashed/gneat/mutation"
)

type Neat struct {
	Population genetics.Population
	Innovation *genetics.InnovationStore

	config       NeatConfig
	experiment   func()
	reporterFunc func()
}

func (n *Neat) SetExperiment(experimentFunc func()) {
	n.experiment = experimentFunc
}

func (n *Neat) SeedPopulation(initialGenome *genetics.Genome) error {
	if initialGenome == nil {
		return genetics.ErrNilInitialGenome
	}

	for i := range n.config.PopulationSize {
		genomeCopy := genetics.CopyGenome(initialGenome)
		// Assigning 0.5 as the weight mutation rate assures that about half of the connections are mutated.
		mutation.ApplyRandWeights(genomeCopy, 0.5, n.config.WeightPerturbationStrength)
		n.Population[i] = genomeCopy
	}

	return nil
}

func (n *Neat) Run(reporterFunc func()) error {
	n.reporterFunc = reporterFunc

	log.Println("I'm not implemented yet")

	return errors.New("not implemented, I can't run anything")
}

func Init(conf NeatConfig) *Neat {
	return &Neat{
		Population: make(genetics.Population, conf.PopulationSize),
		Innovation: genetics.NewInnovationStore(),

		config: conf,
	}
}
