package neat

type Options struct {
	// Ignore trait-related options, they are not going to be implemented, at least for now.
// trait_param_mut_prob 0.5
// trait_mutation_power 1.0
// linktrait_mut_sig 1.0
// nodetrait_mut_sig 0.5
	WeightMutPower float64  `json:"weighMutPower"`
	RecurProb float64  `json:"recurProb"`	// Prob. that a link mutation which doesn't have to be recurrent will be made recurrent 
	
	// These 3 global coefficients are used to determine the formula for
	// computating the compatibility between 2 genomes.  The formula is:
	// disjoint_coeff*pdg+excess_coeff*peg+mutdiff_coeff*mdmg.
	// See the compatibility method in the Genome class for more info
	// They can be thought of as the importance of disjoint Genes,
	// excess Genes, and parametric difference between Genes of the
	// same function, respectively. 
	DisjointCoeff float64 `json:"disjointCoeff"`
	ExcessCoeff float64 `json:"excessCoeff"`
	MutDiffCoeff float64 `json:"mutdiff_coeff"`

	CompatThreshold float64 `json:"compatThresh"`	// This global tells compatibility threshold under which two Genomes are considered the same species 
	AgeSignificance float64 `json:"ageSignificance"`	// How much does age matter?
	SurvivalThresh float64 `json:"survivalThresh"`	// Percent of ave fitness for survival 
    MutateOnlyProb float64 `json:"mutate_only_prob"`	// Prob. of a non-mating reproduction
	MutateLinkWeightsProb float64 `json:"mutate_link_weights_prob"`
	MutateToggleEnableProb	float64 `json:"mutate_toggle_enable_prob"`	// Probability for either enabling or disabling a gene
	MutateGeneReenableProb float64 `json:"mutate_gene_reenable_prob"`	// Probability for enabling a disabled gene
	MutateAddNodeProb float64 `json:"mutate_add_node_prob"`
	MutateAddLinkProb float64 `json:"mutate_add_link_prob"`
	InterspeciesMatingRate float64 `json:"interspecies_mating_rate"`	// Prob. of a mate being outside species
mate_multipoint_prob 0.6
mate_multipoint_avg_prob 0.4
mate_singlepoint_prob 0.0
mate_only_prob 0.2	// Prob. of mating without mutation 
recur_only_prob;  // Probability of forcing selection of ONLY links that are naturally recurrent 
pop_size;  // Size of population 
dropoff_age;  // Age where Species starts to be penalized 
newlink_tries;  // Number of tries mutate_add_link will attempt to find an open link 
print_every; // Tells to print population to file every n generations 
babies_stolen; // The number of babies to siphon off to the champions
num_runs 100
}
