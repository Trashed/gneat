package genome

type Gene struct {
	// innovation/id, weight, input id, output id, enabled/disabled
	Innovation uint
	Weight     float64
	InNode     *Node
	OutNode    *Node
	Enabled    bool
	Recurrent  bool
}
