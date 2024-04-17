package gneat

import (
	"github.com/Trashed/gneat/genetics"
)

type NeatOptions struct {
	Name string `json:"name"`

	// The power of a link weight mutation
	WeightMutPower float64 `json:"weightMutPower"`

	DisjointCoeff float64 `json:"disjointCoeff"`
	ExcessCoeff   float64 `json:"excessCoeff"`
	MutdiffCoeff  float64 `json:"mutdiffCoeff"`

	// This global tells compatibility threshold under which
	// two Genomes are considered the same species
	CompatThreshold float64 `json:"compatThreshold"`

	/* Globals involved in the epoch cycle - mating, reproduction, etc.. */

	// How much does age matter? Gives a fitness boost up to some young age (niching).
	// If it is 1, then young species get no fitness boost.
	// TODO: AgeSignificance float64 `json:"age_significance"`
	// Percent of average fitness for survival, how many get to reproduce based on survival_thresh * pop_size
	SurvivalThreshold float64 `json:"survivalThreshold"`

	// Probabilities of a non-mating reproduction
	MutateOnlyProb         float64 `json:"mutateOnlyProb"`
	MutateLinkWeightsProb  float64 `json:"mutateLinkWeightsProb"`
	MutateToggleEnableProb float64 `json:"mutateToggleEnableProb"`
	MutateGeneReenableProb float64 `json:"mutateGeneReenableProb"`
	MutateAddNodeProb      float64 `json:"mutateAddNodeProb"`
	MutateAddLinkProb      float64 `json:"mutate_add_link_prob"`
	// probability of mutation involving disconnected inputs connection
	MutateConnectSensors float64 `json:"mutateConnectSensors"`

	// Probabilities of a mate being outside species
	InterspeciesMateRate float64 `json:"interspeciesMateRate"`
	// MateMultipointProb    float64 `json:"mateMultipoint_prob"`
	// MateMultipointAvgProb float64 `json:"mate_multipoint_avg_prob"`
	// MateSinglepointProb   float64 `json:"mate_singlepoint_prob"`

	// Prob. of mating without mutation
	MateOnlyProb float64 `json:"mateOnlyProb"`
	// Probability of forcing selection of ONLY links that are naturally recurrent
	RecurOnlyProb float64 `json:"recurOnlyProb"`

	// Size of population
	PopSize int `json:"popSize"`
	// Age when Species starts to be penalized
	DropOffAge int `json:"dropoffAge"`
	// Number of tries mutate_add_link will attempt to find an open link
	NewLinkTries int `json:"newlinkTries"`

	// Tells to print population to file every n generations
	// PrintEvery int `json:"print_every"`

	// The number of babies to stolen off to the champions
	// BabiesStolen int `json:"babies_stolen"`

	// The number of runs to average over in an experiment
	NumEpochs int `json:"numEpochs"`

	// The number of generations to execute training
	NumGenerations int `json:"numGenerations"`

	// The epoch's executor type to apply (sequential, parallel)
	// EpochExecutorType EpochExecutorType `json:"epoch_executor"`
	// The genome compatibility testing method to use (linear, fast (make sense for large genomes))
	// GenCompatMethod GenomeCompatibilityMethod `json:"genome_compat_method"`

	// The neuron nodes activation functions list to choose from
	// NodeActivators []math.NodeActivationType `json:"-"`
	// The probabilities of selection of the specific node activator function
	// NodeActivatorsProb []float64 `json:"-"`

	// NodeActivatorsWithProbs the list of supported node activation with probability of each one
	// NodeActivatorsWithProbs []string `json:"node_activators"`

	// LogLevel the log output details level
	// LogLevel string `json:"log_level"`

	StartGenome *genetics.Genome `json:"startGenome"`
}
