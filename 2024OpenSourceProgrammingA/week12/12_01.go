package main

import (
	"fmt"
	"time"
)

func main() {
	//var dates [3]time.Time
	//dates[1] = time.Unix(1447920000, 0)
	//fmt.Println(dates[1])

	/*dates := [3]time.Time{
		time.Unix(1447920000, 0),
		time.Unix(1447920001, 0), // ← " , " 시험?!
	}
	fmt.Println(dates[0], dates[1], dates[2])*/

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1947920000, 0),
	}
	//fmt.Println(dates[0], dates[1], dates[2])
	//fmt.Println(dates)
	//fmt.Printf("")

	/*for i := 0; i < len(dates); i++ { //안전함
		fmt.Println(i, dates[i])
	}*/

	for i, v := range dates { //안전함
		fmt.Println(i, v)
	}

}
