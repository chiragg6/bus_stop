package main

import (
	"fmt"
	"os"

	"github.com/bus_stop/bus"
)

func main() {
	if len(os.Args) < 4 {
		// Checking length of passed arguments
		fmt.Println("Please send proper parameters")
		os.Exit(1)
	}
	bus.GetETATime(os.Args[1], os.Args[2], os.Args[3])
}
