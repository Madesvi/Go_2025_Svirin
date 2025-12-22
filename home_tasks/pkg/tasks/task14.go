package tasks

import "fmt"

func SolutionTasks14() {

	mass := [...]int{1, 2, 17, 54, 30, 89, 2, 1, 6, 2}
	// fmt.Printf("Array: %v with len: %v\n", mass, len(mass))

	// Воспользоваться range для index и value ?
	// Возможно использовать map где ключ это value а значение это индекс

	indexMap := make(map[int]([]int))

	for i, v := range mass {
		// fmt.Printf("Index:%v, Value:%v\n", i, v)
		indexMap[v] = append(indexMap[v], i)
	}
	fmt.Printf("Map after insert: %v\n", indexMap)
	//	Map after insert: map[1:[0 7] 2:[1 6 9] 6:[8] 17:[2] 30:[4] 54:[3] 89:[5]]

	for k, v := range indexMap {
		if len(v) == 1 {
			// fmt.Println("Минимальных расстояний нет")
		} else if len(v) >= 2 {
			// fmt.Printf("Для числа:%v можно найти минимальное расстояние\n", k)
			minDistance := len(mass) + 1
			var index1 int
			var index2 int
			for i := 0; i < len(v)-1; i++ {
				distance := v[i+1] - v[i]
				if distance < minDistance {
					minDistance = distance
					index1 = v[i]
					index2 = v[i+1]
				}
			}
			fmt.Printf("Для числа: %d минимальное расстояние по индексам: %d, %d\n", k, index1, index2)
		}
	}

}
