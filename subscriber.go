package gobserve

//Subscribe attach an action to an event name given.
//Internally, this function, will attach the action to a normalPriority (0)
func (e *EventDispatcher[T]) Subscribe(name string, action func(T) error) {
	e.SubscribeWithPriority(name, action, NormalPriority)
}

//SubscribeWithPriority attach an action to an event name given and weights the priority
//The higher the priority the first will the action be trigger.
func (e *EventDispatcher[T]) SubscribeWithPriority(name string, action func(T) error, priority int) {
	if _, ok := e.subscriptions[name]; !ok {
		e.subscriptions[name] = map[int][]func(T) error{}
	}
	if _, ok := e.subscriptions[name][priority]; !ok {
		e.subscriptions[name][priority] = []func(T) error{}
	}
	e.subscriptions[name][priority] = append(e.subscriptions[name][priority], action)
}
