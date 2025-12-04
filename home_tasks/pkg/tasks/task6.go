package tasks

import (
	"fmt"
)

func SolutionTasks6() {
	var a int64
	fmt.Print("Введите сторону квадрата:")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	s := a * a
	fmt.Printf("Площадь квадрата равна: %d\n", s)
}
