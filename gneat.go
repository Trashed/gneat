/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gneat

import (
	"errors"
	"log"

	"github.com/Trashed/gneat/genome"
)

type Neat struct {
	config       NeatConfig
	experiment   func()
	reporterFunc func()
}

func (n *Neat) SetExperiment(experimentFunc func()) {
	n.experiment = experimentFunc
}

func (n *Neat) SeedPopulation(initialGenome *genome.Genome) {
	log.Printf("initial genome: %+v\n", initialGenome)
}

func (n *Neat) Run(reporterFunc func()) error {
	n.reporterFunc = reporterFunc

	log.Println("I can't run anything")

	return errors.New("not implemented, I can't run anything")
}

func Init(neatConf NeatConfig) *Neat {
	return nil
}
