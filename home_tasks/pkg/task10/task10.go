package task10

import "fmt"

func Solution() {

	// A четное или нет
	a := 2

	if a%2 == 0 {
		fmt.Printf("Число %d чётное - %t\n", a, true)
	} else {
		fmt.Printf("Число %d не чётное - %t\n", a, false)
	}

}
