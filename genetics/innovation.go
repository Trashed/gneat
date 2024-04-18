package genetics

import "errors"

var (
	ErrInvalidGeneValue error = errors.New("can't provide empty or null value as gene")
)

// InnovationRecord is a data type that holds information about unique connections during the evolution process.
// When a unique connection emerges, the currentInnovation counter is incremented and the gene is added to the
// store. If the connection already exists in the store, it's innovation number is returned.
type InnovationRecord[C comparable] struct {
	geneStore         map[C]uint
	currentInnovation uint
}

// NewInnovationRecord returns a new [InnovationRecord].
func NewInnovationRecord[C comparable]() *InnovationRecord[C] {
	return &InnovationRecord[C]{
		geneStore:         map[C]uint{},
		currentInnovation: 0,
	}
}

// Store checks whether the gene provided already exists. If it doesn't exist, it is stored and the innovation
// number is increased. Otherwise the innovation number related to that gene is returned. If the provided gene
// is an empty value, zero value and an error is returned.
func (ir *InnovationRecord[C]) Store(gene C) (uint, error) {

	if isInvalid(gene) {
		return 0, ErrInvalidGeneValue
	}

	if innov, found := ir.geneStore[gene]; !found {
		ir.currentInnovation++
		ir.geneStore[gene] = ir.currentInnovation
		return ir.currentInnovation, nil
	} else {
		return innov, nil
	}
}

func isInvalid[C comparable](val C) bool {
	// https://stackoverflow.com/questions/70585852/return-default-value-for-generic-type/70589302#70589302
	return val == *new(C)
}
