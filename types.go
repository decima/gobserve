package gobserve

type (
	//ProcessingStrategy the type of the
	ProcessingStrategy int
)

const (
	HighPriority = 100
	//NormalPriority default priority
	NormalPriority = 0
	LowPriority    = -100

	//Sequential runs every action sequentially.
	Sequential ProcessingStrategy = iota
	//Concurrent runs every action at once, regardless priority
	Concurrent
	//PerPriorityConcurrent runs every actions concurrently regarding every priority. The priority processing is sequential.
	PerPriorityConcurrent
)
