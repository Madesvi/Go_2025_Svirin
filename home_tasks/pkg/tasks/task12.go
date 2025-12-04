package tasks

import "fmt"

func SolutionTasks12() {

	// todo: Единицы массы пронумерованы следующим образом:
	// 1 — килограмм, 2 — миллиграмм, 3 — грамм,
	// 4 — тонна, 5 — центнер. Дан номер единицы массы и
	// масса тела M в этих единицах (вещественное число).
	// Вывести массу данного тела в килограммах

	massType := map[int]string{
		1: "килограмм",
		2: "милиграмм",
		3: "грамм",
		4: "тонн",
		5: "центнер",
	}
	// fmt.Println(massType)
	// var bodyMass int

	bodyMassNum := 1
	bodyMass := 5

	for k, v := range massType {
		if k == bodyMassNum {
			fmt.Printf("Масса тела bodyMass равна: %d %s\n", bodyMass, v)
		}
	}

	// 	switch bodyMassNum {
	// case k:
	// 	{
	// 		fmt.Printf("Масса тела bodyMass равна: %d %s\n", bodyMass, v)
	// 	}
	// }
}
