package main

import (
	"github.com/decima/gobserve"
	"log"
	"time"
)

type helloMessage struct {
	From string
}

func (s helloMessage) Name() string {
	return "sayHello"
}

func (s helloMessage) Source() string {
	return s.From
}

func main() {

	gobserve.SubscribeWithPriority("sayHello", func(h helloMessage) error {
		log.Println("I lately received Hello from " + h.Source())
		time.Sleep(2 * time.Second)

		return nil
	}, -100)

	gobserve.SubscribeWithPriority("sayHello", func(h helloMessage) error {
		log.Println("I first received Hello from " + h.Source())
		time.Sleep(2 * time.Second)
		return nil
	}, 200)

	gobserve.SubscribeWithPriority("sayHello", func(h helloMessage) error {
		log.Println("I received Hello from " + h.Source())
		time.Sleep(2 * time.Second)

		return nil
	}, 100)

	gobserve.DispatchSequential(helloMessage{From: "Henri"})
}
