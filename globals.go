package gobserve

//globalEventDispatcher is an instance of the EventDispatch.
//its purpose is to simplify the usage of the event dispatcher.
var globalEventDispatcher = NewDispatcher()

//Subscribe attach an action to an event name given.
//Internally, this function, will attach the action to a normalPriority (0)
func Subscribe[T EventInterface](name string, action func(T) error) interface{} {
	return SubscribeWithPriority(name, action, NormalPriority)
}

//SubscribeWithPriority attach an action to an event name given and weights the priority
//The higher the priority the first will the action be trigger.
func SubscribeWithPriority[T EventInterface](name string, action func(T) error, priority int) interface{} {
	return globalEventDispatcher.SubscribeWithPriority(name, func(event EventInterface) error {
		return action(event.(T))
	}, priority)
}

//Unsubscribe removes the subscription previously made.
func Unsubscribe(sub interface{}) {
	globalEventDispatcher.Unsubscribe(sub)
}

//Dispatch sends an event to every same name events.
//Processing Strategies are defined for concurrency/sequential behaviors.
func Dispatch[T EventInterface](event T, strategy ProcessingStrategy) []error {
	return globalEventDispatcher.Dispatch(event, strategy)
}

//DispatchSequential sends an event through every action in a sequential mode.
func DispatchSequential[T EventInterface](event T) []error {
	return Dispatch(event, Sequential)
}

//DispatchConcurrentPerPriority processes every priority sequentially, but with concurrent working in the tasks.
func DispatchConcurrentPerPriority[T EventInterface](event T) []error {
	return Dispatch(event, PerPriorityConcurrent)
}

func DispatchConcurrent[T EventInterface](event T) []error {
	return Dispatch(event, Concurrent)
}
