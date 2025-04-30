package gneat_test

import (
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/Trashed/gneat"
	g "github.com/Trashed/gneat/genetics"
)

func TestSeedPopulation(t *testing.T) {
	t.Parallel()

	fs := fstest.MapFS{
		"empty_genome": &fstest.MapFile{
			Data: []byte(""),
		},
		"simple_genome": &fstest.MapFile{
			Data: []byte(`genomestart 1
// node, id, type
node 1 0  // bias node
node 2 1  // input node
node 3 1  // input node
node 4 2  // output node
// connection gene, innovation/id, weight, input id, output id, enabled/disabled
gene 1 1.0 1 4 1  // bias -> output
gene 2 1.0 2 4 1  // input 1 -> output
gene 3 1.0 3 4 1  // input 2 -> output
genomeend 1`),
		},
	}

	tests := []struct {
		name               string
		genomeName         string
		expectedPopulation int
		expectedError      error
	}{
		{
			name:          "Nil initial genome",
			genomeName:    "empty_genome",
			expectedError: g.ErrNilInitialGenome,
		},
		{
			name:          "Simple genome",
			genomeName:    "simple_genome",
			expectedError: nil,
		},
	}

	for _, test := range tests {

		initialGenome, err := initialGenomeFromFile(fs, test.genomeName)

		if err != nil && err == g.ErrEmptyGenomeFile {
			continue
		}

		t.Run(test.name, func(t *testing.T) {

			conf := gneat.NeatConfig{
				PopulationSize: 100,
			}

			n := gneat.Init(conf)

			err = n.SeedPopulation(initialGenome)

			if err != nil && test.expectedError != err {
				t.Fatalf("unexpected failure in seeding the population: %v\n", err)
			}

			if n.Population.Size() != conf.PopulationSize {
				t.Fatalf("expected population size to be %d but got %d\n", n.Population.Size(), conf.PopulationSize)
			}
		})
	}
}

func initialGenomeFromFile(fs fs.FS, fileName string) (*g.Genome, error) {
	f, err := fs.Open(fileName)

	if err != nil {
		return nil, err
	}

	genome, err := g.ReadGenome(f)
	if err != nil {
		return nil, err
	}

	return genome, nil
}
