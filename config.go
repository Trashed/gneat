/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gneat

type NeatConfig struct {
	// The name of the current experiment. Can be empty.
	Name string `json:"name,omitempty"`
	// The number of genomes (individuals) in the population per generation.
	PopulationSize int `json:"populationSize"`
	// The probability that two genomes will undergo crossover (i.e., offspring is produced by combining two parents' genomes).
	CrossoverRate float64 `json:"crossoverRate"` // Typical value 0.75
	// The overall probability that any given genome will undergo mutation (e.g., weight mutation, connection mutation, etc.).
	MutationRate float64 `json:"mutationRate"` // Typical value 0.25
	// The probability that a connection's weight will be mutated.
	WeightMutationRate float64 `json:"weightMutationRate"` // Typical value 0.8
	// The range or magnitude by which connection weights can be perturbed (i.e., changed during mutation).
	WeightPerturbationStrength float64 `json:"weightPerturbationStrength"` // Typical value ±0.1 to ±2.0
	// The probability that a new node (neuron) will be added by splitting an existing connection.
	AddNodeMutationRate float64 `json:"addNodeMutationRate"` // Typical value 0.03
	// The probability that a new connection between two nodes will be added.
	AddConnectionMutationRate float64 `json:"addConnectionMutationRate"` // Typical value 0.05-0.3
	// The probability that an existing connection will be removed.
	RemoveConnectionMutationRate float64 `json:"removeConnectionMutationRate"` // Typical value 0.01 - 0.05
	// The probability that node's activation function is changed to another one.
	ChangeActivationFunctionMutationRate float64 `json:"changeActivationFunctionMutationRate"` // 0.01 - 0.03
	// The activation function used in neurons to calculate their output.
	ActivationFunctions []string `json:"activationFunctions"` // Supported: Sigmoid, Tanh, ReLU, Identity (Linear), Step function (Binary)
	// The threshold for determining whether two genomes are similar enough to belong to the same species.
	CompatibilityThreshold float64 `json:"compatibilityThreshold"` // Typical value 3.0
	// A coefficient used when calculating the genetic distance between two genomes based on excess genes (genes that appear in one genome but not in the other).
	ExcessGeneCoefficient float64 `json:"excessGeneCoefficient"` // Typical value 1.0
	// A coefficient used when calculating genetic distance based on disjoint genes (genes that are not aligned between two genomes).
	DisjointGeneCoefficient float64 `json:"disjointGeneCoefficient"` // Typical value 1.0
	// A coefficient used when calculating genetic distance based on differences in the weights of matching genes.
	WeightDifferenceCoefficient float64 `json:"weightDifferenceCoefficient"` // Typical value 0.4
	// The number of generations that a species can go without improving its best fitness before it is eliminated.
	SpeciesStagnationLimit int `json:"speciesStagnationLimit"` // Typical value 15-20
	// The fraction of each species allowed to reproduce. Typically, only the top-performing individuals in each species are allowed to reproduce.
	SurvivalThreshold float64 `json:"survivalThreshold"` // Typical value 0.2
	// The number of generations for which the NEAT algorithm will evolve the population.
	MaxGenerations int `json:"maxGenerations"` // Typical value 100-1000
	// The fraction of potential connections between input and output nodes that are present at initialization.
	InitialConnectionDensity float64 `json:"initialConnectionDensity"` // Typical value 0.05-0.2
	// The number of top individuals that are automatically carried over to the next generation without mutation or crossover.
	Elitism int `json:"elitism"` // Typical value 1-5
}
