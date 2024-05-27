package main

import (
	"fmt"
	"github.com/gavinin/go-state-machine"
	"time"
)

var stateManager = state.NewStateManager[Example]()

type Example struct {
	Content string
}

func main() {
	testChan := make(chan Example, 0)

	managerEvent := stateManager.NewStateManagerEvent(state.ADD, "Example", testChan, 5*time.Second, func(ts chan Example) {
		ts <- Example{
			"This is example.",
		}
	})

	managerEvent.SetTimes(1)

	stateManager.SendEvent(managerEvent)

	content := <-testChan
	fmt.Print(content.Content)

}
