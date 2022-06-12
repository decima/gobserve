package gobserve

import "github.com/google/uuid"

type Subscriber[T EventInterface] struct {
	eventName  string
	priority   int
	uniqueId   string
	dispatcher *EventDispatcher[T]
}

//Subscribe attach an action to an event name given.
//Internally, this function, will attach the action to a normalPriority (0)
func (e *EventDispatcher[T]) Subscribe(name string, action func(T) error) Subscriber[T] {
	return e.SubscribeWithPriority(name, action, NormalPriority).(Subscriber[T])
}

//SubscribeWithPriority attach an action to an event name given and weights the priority
//The higher the priority the first will the action be trigger.
func (e *EventDispatcher[T]) SubscribeWithPriority(name string, action func(T) error, priority int) interface{} {
	if _, ok := e.subscriptions[name]; !ok {
		e.subscriptions[name] = map[int]map[string]actionType[T]{}
	}
	if _, ok := e.subscriptions[name][priority]; !ok {
		e.subscriptions[name][priority] = map[string]actionType[T]{}
	}

	uniqueId := uuid.NewString()
	e.subscriptions[name][priority][uniqueId] = action

	return Subscriber[T]{
		eventName: name,
		priority:  priority,
		uniqueId:  uniqueId,
	}
}

//Unsubscribe removes the subscription previously made.
func (e *EventDispatcher[T]) Unsubscribe(rawSub interface{}) {
	sub := rawSub.(Subscriber[T])
	if _, ok := e.subscriptions[sub.eventName][sub.priority][sub.uniqueId]; ok {
		delete(e.subscriptions[sub.eventName][sub.priority], sub.uniqueId)
	}
	if len(e.subscriptions[sub.eventName][sub.priority]) == 0 {
		delete(e.subscriptions[sub.eventName], sub.priority)
	}
	if len(e.subscriptions[sub.eventName]) == 0 {
		delete(e.subscriptions, sub.eventName)
	}
}
