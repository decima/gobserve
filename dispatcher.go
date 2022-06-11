package gobserve

import (
	"sync"
)

// This structure contains a map of every events registered in the event Dispatcher
type EventDispatcher[T EventInterface] struct {
	subscriptions map[string]map[int][]func(T) error
}

// NewDispatcher creates a new eventDispatcher for EventInterfaces
func NewDispatcher() EventDispatcher[EventInterface] {
	return EventDispatcher[EventInterface]{
		subscriptions: map[string]map[int][]func(EventInterface) error{},
	}
}

//Dispatch sends an event through every concerned event subscribers
func (e *EventDispatcher[T]) Dispatch(
	event T,
	processingStrategy ProcessingStrategy,
) []error {
	errorListMutex := sync.Mutex{}
	errorList := []error{}
	if _, ok := e.subscriptions[event.Name()]; !ok {
		return nil
	}

	actionsPerPriority := sortIntMap(e.subscriptions[event.Name()])

	wg := sync.WaitGroup{}

	for _, actionList := range actionsPerPriority {

		for _, action := range actionList {
			wg.Add(1)

			go func(action func(T) error) {
				err := action(event)
				if err != nil {
					errorListMutex.Lock()
					errorList = append(errorList, err)
					errorListMutex.Unlock()
				}
				wg.Done()
			}(action)
			if processingStrategy == Sequential {
				wg.Wait()
			}
		}

		if processingStrategy == PerPriorityConcurrent {
			wg.Wait()
		}
	}
	if processingStrategy == Concurrent {
		wg.Wait()
	}

	if len(errorList) == 0 {
		return nil
	}
	return errorList
}
