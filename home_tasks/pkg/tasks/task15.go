package tasks

import "fmt"

func SolutionTasks15() {

	mass := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	massAfter := mass[:]
	// fmt.Println(massAfter)
	newMass := []int{}
	for _, v := range massAfter {
		newMass = append(newMass, v+1)
	}
	fmt.Println(newMass)

}
