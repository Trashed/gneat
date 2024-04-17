package main

import (
	"log"

	"github.com/Trashed/gneat"
	"github.com/Trashed/gneat/examples/xor"
)

func main() {
	neatConfig, err := gneat.FromConfig("../../../examples/xor/xor.json")
	if err != nil {
		log.Fatalf("error parsing configuration: %v\n", err)
	}

	gneat.New(neatConfig, xor.NewXORExperiment()).Start()
	/*.OnError(func(err error) {
		log.Fatalf("error when running NEAT: %v\n", err)
	})*/
}
