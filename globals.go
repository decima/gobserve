package gobserve

var globalEventDispatcher = NewDispatcher()

func Dispatch[T EventInterface](name string, event T) {
	globalEventDispatcher.Dispatch(event, Sequential)
}

func DispatchPerPriority[T EventInterface](name string, event T) {
	globalEventDispatcher.Dispatch(event, PerPriorityConcurrent)
}

func DispatchConcurrent[T EventInterface](name string, event T) {
	globalEventDispatcher.Dispatch(event, Concurrent)
}

func Subscribe[T EventInterface](name string, action func(T) error) {
	globalEventDispatcher.Subscribe(name, func(event EventInterface) error {
		return action(event.(T))
	})
}

func SubscribeWithPriority[T EventInterface](name string, action func(T) error, priority int) {
	globalEventDispatcher.SubscribeWithPriority(name, func(event EventInterface) error {
		return action(event.(T))
	}, priority)
}
