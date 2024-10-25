package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1
	fmt.Printf("%d\n", answer)

	win := false
	for guesses := 0; guesses < 3; guesses++ {
		fmt.Printf("%d숫자입력: ")
		r := bufio.NewReader(os.Stdin)
		i, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		i = strings.TrimSpace(i)
		guess, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guess)

		if answer == guess {
			fmt.Println("정답")
			win = true
			break
		} else if answer > guess {
			fmt.Println("LOW")
		} else if answer < guess {
			fmt.Println("High")
		}
	}
	if win {
		fmt.Println("win.")
	} else {
		fmt.Println("lose. %d\n", answer)
	}
}
