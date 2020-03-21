package main

import (
	"fmt"
)

var queue chan Alert

func main() {
	queue = make(chan Alert)
	startWeb()

	for alert := range queue {
		fmt.Println("Got alert:", alert)
	}
}
