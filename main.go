package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Please send proper parameters")
		os.Exit(1)
	}
	metro.GetETATime(os.Args[1], os.Args[2], os.Args[3])
}
