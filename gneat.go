package gneat

import (
	"errors"
	"log"

	"github.com/Trashed/gneat/experiment"
)

type Gneat struct {
}

func (g *Gneat) Start() {
	log.Fatalln("not implemented")
}

func FromConfig(configPath string) (*NeatOptions, error) {
	return nil, errors.New("not implemented")
}

func New(conf *NeatOptions, experiment experiment.Evaluator) *Gneat {
	return &Gneat{}
}
