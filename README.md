Gobserve
=====

A simple event-dispatcher library.

[![Go Reference](https://pkg.go.dev/badge/github.com/decima/gobserve.svg)](https://pkg.go.dev/github.com/decima/gobserve)

## Install

to add the library to your project simply run :

```bash
go get github.com/decima/gobserve
```

## example usage :

```go
package main

import (
	"github.com/decima/gobserve"
	"log"
)

type helloWorld int

func (s helloWorld) Name() string {
	return "sayHello"
}

func main() {
	gobserve.Subscribe("sayHello", func(h helloWorld) error {
		log.Println("hello from 1")
		return nil
	})
	gobserve.Subscribe("sayHello", func(h helloWorld) error {
		log.Println("hello from 2")
		return nil
	})

	gobserve.DispatchConcurrent(helloWorld(1))

}
```

By default you can use the Subscribe and Dispatch methods in order to use a globally defined event Subscriber. 
But you can create your own dispatcher for more precise and independant control.
