package gneat_test

import (
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/Trashed/gneat"
	g "github.com/Trashed/gneat/genome"
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
		initialGenome      *g.Genome
		expectedPopulation int
		expectedError      error
	}{
		{
			name:          "Nil initial genome",
			initialGenome: initialGenomeFromFile(t, fs, "empty_genome"),
			expectedError: g.ErrNilInitialGenome,
		},
		{
			name:          "Simple genome",
			initialGenome: initialGenomeFromFile(t, fs, "simple_genome"),
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n := gneat.Neat{}
			err := n.SeedPopulation(test.initialGenome)

			if test.expectedError != err {
				t.Fatalf("expected error \"%v\" but got \"%v\"\n", test.expectedError, err)
			}
		})
	}
}

func initialGenomeFromFile(t *testing.T, fs fs.FS, fileName string) *g.Genome {
	f, err := fs.Open(fileName)

	if err != nil {
		t.Fatal("reading genome from file failed: " + err.Error())
	}

	genome, err := g.ReadGenome(f)
	if err != nil {
		t.Fatalf("reading genome content failed: %v\n", err)
	}

	return genome
}
