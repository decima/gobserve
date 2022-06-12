package main

import (
	"github.com/decima/gobserve"
	"log"
)

type hello2Message struct {
}

func (s hello2Message) Name() string {
	return "sayHello2"
}

func main() {

	gobserve.SubscribeWithPriority("sayHello2", func(h hello2Message) error {
		log.Println("hello from subscriber with prio 3")
		return nil
	}, 3)

	subscription := gobserve.SubscribeWithPriority("sayHello2", func(h hello2Message) error {
		log.Println("hello from subscriber with prio 2")

		return nil
	}, 2)

	gobserve.SubscribeWithPriority("sayHello2", func(h hello2Message) error {
		log.Println("hello from subscriber with prio 1")

		return nil
	}, 1)

	gobserve.DispatchSequential(hello2Message{})
	gobserve.Unsubscribe(subscription)
	gobserve.DispatchSequential(hello2Message{})
}
