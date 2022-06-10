package gobserve

type (
	ProcessingStrategy int
)

const (
	HighPriority   = 100
	NormalPriority = 0
	LowPriority    = -100

	Sequential ProcessingStrategy = iota
	Concurrent
	PerPriorityConcurrent
)
