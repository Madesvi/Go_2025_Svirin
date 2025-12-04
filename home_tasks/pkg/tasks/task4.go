package tasks

import "fmt"

func SolutionTasks4() {
	// x := 10
	// y := 15
	// z := 2

	x := 77
	y := 9
	z := 130

	switch {
	case x > y && x > z:
		fmt.Printf("Наибольшее число: %d\n", x)
	case y > x && y > z:
		fmt.Printf("Наибольшее число: %d\n", y)
	case z > y && z > x:
		fmt.Printf("Наибольшее число: %d\n", z)
	}
}
