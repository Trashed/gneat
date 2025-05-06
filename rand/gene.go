package rand

type ItemList interface {
	Len() int
}

// GeneItem returns a random ID of a genetic item.
func GeneItem(items ItemList) uint {
	return 0
}
