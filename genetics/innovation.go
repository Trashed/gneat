/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package genetics

import "strconv"

type storedGene struct {
	key        string
	innovation uint
	inNode     uint
	outNode    uint
}

// InnovationStore is a structure that manages the innovation numbers of genes.
type InnovationStore struct {
	innovationMap map[string]*storedGene
	innovationNum uint
}

// NewInnovationStore initializes a new InnovationStore.
// It creates a map to store genes and sets the innovation number to zero.
// This function should be called before using the store.
func NewInnovationStore() *InnovationStore {
	return &InnovationStore{
		innovationMap: make(map[string]*storedGene),
	}
}

// AddGene adds a gene to the innovation store and returns its innovation number.
// If the gene already exists, it returns the existing innovation number and false.
func (s *InnovationStore) AddGene(gene *Gene) (uint, bool) {
	if gene == nil {
		return 0, false
	}

	if storedGene, exists := s.innovationMap[toStringKey(gene)]; exists {
		gene.Innovation = storedGene.innovation
		return gene.Innovation, false
	}

	s.innovationNum++
	gene.Innovation = s.innovationNum
	geneKey := toStringKey(gene)
	s.innovationMap[geneKey] = &storedGene{
		key:        geneKey,
		innovation: s.innovationNum,
		inNode:     gene.InNode.Id,
		outNode:    gene.OutNode.Id,
	}

	return s.innovationNum, true
}

func (s *InnovationStore) Innovation() uint {
	return s.innovationNum
}

func toStringKey(gene *Gene) string {

	inNodeStr := strconv.FormatUint(uint64(gene.InNode.Id), 10)
	outNodeStr := strconv.FormatUint(uint64(gene.OutNode.Id), 10)

	return inNodeStr + "->" + outNodeStr
}
