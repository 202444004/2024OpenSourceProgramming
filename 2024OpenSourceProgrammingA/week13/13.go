package main

import "fmt"

type visitor struct {
	age  int
	cost int
}

func caculateCost(visitors []visitor) int {
	totalcost := 0
	for _, v := range visitors {
		totalcost = totalcost + v.cost
	}
	return totalcost
}

func main() {
	var numvisitors int
	fmt.Println("How many visitors?")
	fmt.Scanln(&numvisitors)

	var vs []visitor
	vs = make([]visitor, numvisitors)

	for i := 0; i < numvisitors; i++ {
		var age int
		fmt.Print("input age:")
		fmt.Scan(&age)

		switch {
		case age <= 12:
			vs[i] = visitor{age: age, cost: 5000}
		case age >= 12 && age < 65:
			vs[i] = visitor{age: age, cost: 10000}
		default:
			vs[i] = visitor{age: age, cost: 7000}
		}
	}

	fmt.Printf("total price is %dwon\n", caculateCost(vs))
}
