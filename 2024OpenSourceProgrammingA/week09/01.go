package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(100) + 1
	fmt.Printf("%d\n", target)

}
