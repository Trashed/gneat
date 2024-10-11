/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Trashed/gneat/config"
)

func TestFromFile(t *testing.T) {
	t.Parallel()

	testConfName := "testConf.json"
	testName := "Test"
	testPopulationSize := 100
	testMutationRate := 0.05
	testConfJsonStr := fmt.Sprintf(`{
	    "name": "%s",
		"populationSize": %d,
		"crossoverRate": 0.75,
		"mutationRate": 0.25,
		"weightMutationRate": 0.8,
		"weightPerturbationStrength": 0.1,
		"addNodeMutationRate": 0.03,
		"addConnectionMutationRate": %.2f,
		"removeConnectionMutationRate": 0.01,
		"changeActivationFunctionMutationRate": 0.01,
		"activationFunctions": ["Sigmoid", "Tanh"],
		"compatibilityThreshold": 3.0,
		"excessGeneCoefficient": 1.0,
		"disjointGeneCoefficient": 1.0,
		"weightDifferenceCoefficient": 0.4,
		"speciesStagnationLimit": 15,
		"survivalThreshold": 0.2,
		"maxGenerations": 100,
		"initialConnectionDensity": 0.05,
		"elitism": 1
	}`, testName, testPopulationSize, testMutationRate)

	dir := t.TempDir()

	testConfPath := dir + "/" + testConfName

	f, err := os.Create(testConfPath)
	if err != nil {
		t.Fatalf("failed to create %s on temp dir %s: %v\n", testConfName, dir, err)
	}

	_, err = f.WriteString(testConfJsonStr)
	if err != nil {
		t.Fatalf("failed to write test JSON string to file: %v\n", err)
	}

	err = f.Close()
	if err != nil {
		t.Fatalf("error occured while closing the file: %v\n", err)
	}

	conf := config.FromFile(testConfPath)
	if conf.Name != testName && conf.PopulationSize != testPopulationSize && conf.AddConnectionMutationRate != testMutationRate {
		t.Fatalf("test conf JSON was not properly parsed")
	}
}
