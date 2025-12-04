package tasks

import "fmt"

func SolutionTasks10() {

	// A четное или нет
	a := 3

	if a%2 == 0 {
		fmt.Printf("Число %d чётное - %t\n", a, true)
	} else {
		fmt.Printf("Число %d не чётное - %t\n", a, false)
	}

}
