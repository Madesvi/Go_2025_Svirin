package tasks

import "fmt"

func SolutionTasks7() {

	// 0__(A)1______(B)4__________(C)7
	// AC = C - A = 7 - 1 = 6
	// BC = C - B = 7 - 4 = 3
	var a int64
	var b int64
	var c int64
	fmt.Println("Определение длины отрезков на прямой")
	fmt.Print("Введите координату точки A:")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Print("Введите координату точки B:")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Print("Введите координату точки C:")
	_, err = fmt.Scan(&c)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	ac := c - a
	fmt.Printf("Отрезок AC равен = %d\n", ac)
	bc := c - b
	fmt.Printf("Отрезок BC равен = %d\n", bc)

}
