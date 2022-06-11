package main

import (
	"github.com/decima/gobserve"
	"log"
	"time"
)

type RacePodium struct{}

func (r RacePodium) Name() string {
	return "podium"
}

func main() {

	gobserve.SubscribeWithPriority("podium", func(RacePodium) error {
		log.Println("LAURA : I finished ex-ego with SARAH")
		time.Sleep(2 * time.Second)

		return nil
	}, 600)

	gobserve.SubscribeWithPriority("podium", func(RacePodium) error {
		log.Println("BOB : I finished Last")
		time.Sleep(2 * time.Second)
		return nil
	}, 400)

	gobserve.SubscribeWithPriority("podium", func(RacePodium) error {
		log.Println("JOHN :  Almost first!")
		time.Sleep(2 * time.Second)

		return nil

	}, 800)
	gobserve.SubscribeWithPriority("podium", func(RacePodium) error {
		log.Println("KEN : C'mon barbie, let's go party, i'm first!")
		time.Sleep(2 * time.Second)

		return nil

	}, 1000)
	gobserve.SubscribeWithPriority("podium", func(RacePodium) error {
		log.Println("SARAH : I finished at the same time as LAURA")
		time.Sleep(2 * time.Second)

		return nil

	}, 600)

	gobserve.DispatchConcurrentPerPriority(RacePodium{})

}
