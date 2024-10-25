package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(42)
	target := rand.Intn(100) + 1 // 1~101
	fmt.Printf("%d\n", target)

}
