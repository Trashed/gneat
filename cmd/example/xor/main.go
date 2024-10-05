/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"fmt"
	"github.com/Trashed/gneat"
	"github.com/Trashed/gneat/config"
	"log"
)

// TODO: WIP implementation, may possibly and probably will change in the future
func main() {
	neat := gneat.Init(config.FromFile("xor.json")) // TODO: Inject configuration; returns a NEAT object that wraps the components needed for evolving the neural network and running the experiment.
	fmt.Printf("neat object: %+v\n", neat)
	neat.SetExperiment(func() {
		log.Println("this activation function is a stub")
	}) // TODO: Inject experimentFunc to NeatCtx
	// TODO: Read starter genome from file and generate starter population from starter genome
	neat.SeedPopulation("empty genome file")
	// TODO: Run experiment - inject reporterFunc for reporting metrics about training and evaluation
	log.Fatalf("failed to run neat experiment: %v\n", neat.Run(func() {}))
}
