package rand

import "math/rand/v2"

type Listable interface {
	List() []uint
	Len() int
}

// Item returns a random value from a collection. A collection that satisfies the Listable interface is passed
// to the function.
func Item(items Listable) uint {
	vals := items.List()
	randIdx := rand.UintN(uint(items.Len()))

	return vals[randIdx]
}
