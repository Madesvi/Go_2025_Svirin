package tasks

import "fmt"

func SolutionTasks11() {

	// month := make(map[int]string)
	// fmt.Println(month)
	// month[1] = "Январь"
	// fmt.Println(month)

	month := map[int]string{
		1:  "Январь",
		2:  "Февраль",
		3:  "Март",
		4:  "Апрель",
		5:  "Май",
		6:  "Июнь",
		7:  "Июль",
		8:  "Август",
		9:  "Сентябрь",
		10: "Октябрь",
		11: "Ноябрь",
		12: "Декабрь",
	}
	fmt.Println(month)

	fmt.Println("Введите число месяца")
	var scan int
	fmt.Scan(&scan)
	for k, v := range month {
		var season string
		switch scan {
		case 1, 2, 12:
			{
				season = "Зима"
			}
		case 3, 4, 5:
			{
				season = "Весна"
			}
		case 6, 7, 8:
			{
				season = "Лето"
			}
		case 9, 10, 11:
			{
				season = "Осень"
			}
		}
		if scan == k {
			fmt.Printf("Вы ввели месяц: %s, Сезон: %s\n", v, season)
		}

	}
}
