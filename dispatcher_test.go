package gobserve

import (
	"sync"
	"testing"
)

const DEMO_EVENT_NAME = "demo"

type TestEvent struct {
	mux       *sync.Mutex
	callStack *[]string
}

func (t TestEvent) Name() string {
	return DEMO_EVENT_NAME
}

func TestDispatch(t *testing.T) {

	dispatcher := NewDispatcher()

	dispatcher.SubscribeWithPriority(DEMO_EVENT_NAME,
		func(event EventInterface) error {
			return func(evt TestEvent) error {
				*evt.callStack = append(*evt.callStack, "1")

				return nil
			}(event.(TestEvent))
		}, 0)
	dispatcher.SubscribeWithPriority(DEMO_EVENT_NAME,
		func(event EventInterface) error {
			return func(evt TestEvent) error {
				*evt.callStack = append(*evt.callStack, "2")

				return nil
			}(event.(TestEvent))
		}, 0)
	dispatcher.SubscribeWithPriority(DEMO_EVENT_NAME,
		func(event EventInterface) error {
			return func(evt TestEvent) error {
				*evt.callStack = append(*evt.callStack, "3")

				return nil
			}(event.(TestEvent))
		}, 0)
	callStack := []string{}
	dispatcher.Dispatch(TestEvent{callStack: &callStack}, Sequential)

	expected := []string{"1", "2", "3"}

	if len(expected) != len(callStack) {
		t.Errorf("Expected String(%v) is not same as"+
			" actual string (%v)", expected, callStack)
	}
	for i := 0; i < len(expected); i++ {

		if expected[i] != callStack[i] {
			t.Errorf("Expected slice (%v) is not same as"+
				" actual slice (%v)", expected, callStack)
		}
	}

}

func TestDispatch2(t *testing.T) {

	dispatcher := NewDispatcher()

	dispatcher.SubscribeWithPriority(DEMO_EVENT_NAME,
		func(event EventInterface) error {
			return func(evt TestEvent) error {
				*evt.callStack = append(*evt.callStack, "1")

				return nil
			}(event.(TestEvent))
		}, 1)
	dispatcher.SubscribeWithPriority(DEMO_EVENT_NAME,
		func(event EventInterface) error {
			return func(evt TestEvent) error {
				*evt.callStack = append(*evt.callStack, "2")

				return nil
			}(event.(TestEvent))
		}, 2)
	dispatcher.SubscribeWithPriority(DEMO_EVENT_NAME,
		func(event EventInterface) error {
			return func(evt TestEvent) error {
				*evt.callStack = append(*evt.callStack, "3")

				return nil
			}(event.(TestEvent))
		}, 3)
	callStack := []string{}
	dispatcher.Dispatch(TestEvent{callStack: &callStack}, Sequential)

	expected := []string{"3", "2", "1"}

	if len(expected) != len(callStack) {
		t.Errorf("Expected String(%v) is not same as"+
			" actual string (%v)", expected, callStack)
	}
	for i := 0; i < len(expected); i++ {

		if expected[i] != callStack[i] {
			t.Errorf("Expected slice (%v) is not same as"+
				" actual slice (%v)", expected, callStack)
		}
	}
}

func TestDispatch3(t *testing.T) {

	dispatcher := NewDispatcher()

	dispatcher.SubscribeWithPriority(DEMO_EVENT_NAME,
		func(event EventInterface) error {
			return func(evt TestEvent) error {

				*evt.callStack = append(*evt.callStack, "1")

				return nil
			}(event.(TestEvent))
		}, 1)
	dispatcher.SubscribeWithPriority(DEMO_EVENT_NAME,
		func(event EventInterface) error {
			return func(evt TestEvent) error {
				*evt.callStack = append(*evt.callStack, "2")

				return nil
			}(event.(TestEvent))
		}, 2)
	dispatcher.SubscribeWithPriority(DEMO_EVENT_NAME,
		func(event EventInterface) error {
			return func(evt TestEvent) error {
				*evt.callStack = append(*evt.callStack, "3")

				return nil
			}(event.(TestEvent))
		}, 1)
	callStack := []string{}
	dispatcher.Dispatch(TestEvent{callStack: &callStack}, PerPriorityConcurrent)

	if 3 != len(callStack) {
		t.Errorf("Expected Len(3) is not same as"+
			" actual slice (%v)", callStack)
	}

	if "2" != callStack[0] {

		t.Errorf("Expected 2 to be first in the list"+
			" got (%v)", callStack)
	}

}

func TestDispatch4(t *testing.T) {
	dispatcher := NewDispatcher()

	dispatcher.SubscribeWithPriority(DEMO_EVENT_NAME,
		func(event EventInterface) error {
			return func(evt TestEvent) error {
				evt.mux.Lock()
				*evt.callStack = append(*evt.callStack, "1")
				evt.mux.Unlock()
				return nil
			}(event.(TestEvent))
		}, 1)
	dispatcher.SubscribeWithPriority(DEMO_EVENT_NAME,
		func(event EventInterface) error {
			return func(evt TestEvent) error {
				evt.mux.Lock()

				*evt.callStack = append(*evt.callStack, "2")
				evt.mux.Unlock()
				return nil
			}(event.(TestEvent))
		}, 2)
	dispatcher.SubscribeWithPriority(DEMO_EVENT_NAME,
		func(event EventInterface) error {
			return func(evt TestEvent) error {
				evt.mux.Lock()
				*evt.callStack = append(*evt.callStack, "3")
				evt.mux.Unlock()
				return nil
			}(event.(TestEvent))
		}, 1)
	callStack := []string{}
	mutex := sync.Mutex{}
	dispatcher.Dispatch(TestEvent{callStack: &callStack, mux: &mutex}, Concurrent)
	if 3 != len(callStack) {
		t.Errorf("Expected Len(3) is not same as"+
			" actual slice (%v)", callStack)
	}

}
