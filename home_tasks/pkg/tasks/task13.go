package tasks

import "fmt"

func SolutionTasks13() {

	baseYear := 1984
	colors := []string{"зеленой", "красной", "желтой", "белой", "чёрной"}
	animals := []string{
		"крысы", "коровы", "тигра", "зайца",
		"дракона", "змеи", "лошади", "овцы",
		"обезьяны", "курицы", "собаки",
		"свиньи"}

	var enteredYear int
	fmt.Println("Введите год: ")
	fmt.Scan(&enteredYear)

	deltaYear := (enteredYear - baseYear) % 60
	colorIndex := (deltaYear / 2) % 5
	animalIndex := deltaYear % 12

	fmt.Printf("Год: %s %s\n", colors[colorIndex], animals[animalIndex])

	// var x string
	// var y string
	// x = "зеленой"
	// y = "крысы"
	//enter year
	// var enteredYear int
	// fmt.Println("Введите год: ")
	// fmt.Scan(&enteredYear)
	// fmt.Printf("Год %d - %s %s\n", enteredYear, x, y)

}

// В восточном календаре принят 60-летний цикл, состоящий из 12- летних подциклов,
// обозначаемых названиями цвета: зеленый, красный, желтый, белый и черный.
// В каждом подцикле годы носят названия животных: крысы, коровы, тигра, зайца, дракона,
// змеи, лошади, овцы, обезьяны, курицы, собаки и свиньи. По номеру года вывести его название,
// если 1984 год был началом цикла — годом зеленой крысы.
