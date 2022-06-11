package gobserve

var globalEventDispatcher = NewDispatcher()

func Dispatch[T EventInterface](event T, strategy ProcessingStrategy) {
	globalEventDispatcher.Dispatch(event, strategy)
}

func DispatchSequential[T EventInterface](event T) {
	Dispatch(event, Sequential)
}

func DispatchPerPriority[T EventInterface](event T) {
	Dispatch(event, PerPriorityConcurrent)
}

func DispatchConcurrent[T EventInterface](event T) {
	Dispatch(event, Concurrent)
}

func Subscribe[T EventInterface](name string, action func(T) error) {
	SubscribeWithPriority(name, action, NormalPriority)
}

func SubscribeWithPriority[T EventInterface](name string, action func(T) error, priority int) {
	globalEventDispatcher.SubscribeWithPriority(name, func(event EventInterface) error {
		return action(event.(T))
	}, priority)
}
