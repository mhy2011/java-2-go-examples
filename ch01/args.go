package main

import (
	"fmt"
	"os"
)

func main() {
	for i, val := range os.Args {
		fmt.Printf("Args[%d] = %s\n", i, val)
	}
}
