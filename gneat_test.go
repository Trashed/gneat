package gneat_test

import (
	"testing"

	"github.com/Trashed/gneat"
	g "github.com/Trashed/gneat/genome"
)

func TestSeedPopulation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name               string
		initialGenome      *g.Genome
		expectedPopulation int
		expectedError      error
	}{
		{
			name:          "Nil initial genome",
			initialGenome: nil,
			expectedError: g.ErrNilInitialGenome,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n := gneat.Neat{}
			err := n.SeedPopulation(test.initialGenome)

			if err == nil && err != g.ErrNilInitialGenome {
				t.Fatalf("SeedPopulation shouldn't accept a nil genome - expected error \"%s\"", g.ErrNilInitialGenome.Error())
			}
		})
	}
}
