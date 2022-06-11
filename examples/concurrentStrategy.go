package main

import (
	"fmt"
	"github.com/decima/gobserve"
	"log"
)

type VoteEvent struct {
	Result *int
}

func (v VoteEvent) Name() string {
	return "vote"
}

func main() {
	for i := 0; i < 50; i++ {
		func(i int) {
			gobserve.Subscribe("vote", func(v VoteEvent) error {
				if i%3 == 0 {
					*v.Result--
					fmt.Printf("I(%v) voted NO : %v\n", i, *v.Result)

				} else {
					*v.Result++
					fmt.Printf("I(%v) voted YES: %v\n", i, *v.Result)
				}
				return nil
			})
		}(i)
	}

	result := 0
	v := VoteEvent{
		Result: &result,
	}
	gobserve.DispatchConcurrent(v)

	log.Println(*v.Result)
}
