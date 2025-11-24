package task3

import "fmt"

func Solution() {
	// === task 3 ===
	age := 36.6
	temperature := 25

	// temp := age
	// age = float64(temperature)
	// temperature = int(temp)

	age, temperature = float64(temperature), int(age)

	fmt.Printf("age is: %v, temperature is: %v\n", age, temperature)

	// === Чтобы оставить дробную часть ===
	var age2 float64 = 36.6
	var temperature2 float64 = 25

	age2, temperature2 = temperature2, age2
	fmt.Printf("age2 is: %v, and temperature is: %v\n", age2, temperature2)

	// // === Попытка через string ===
	// str_age := fmt.Sprintf("%v", age)
	// fmt.Printf("String is: %s, with type: %T\n", str_age, str_age)
	// str_temperature := fmt.Sprintf("%v", temperature)
	// fmt.Printf("String is: %s, with type: %T\n", str_temperature, str_temperature)

	// var tmp string
	// tmp = str_age
	// // fmt.Println(tmp)
	// str_age = str_temperature
	// str_temperature = tmp

	// fmt.Println(str_age)
	// fmt.Println(str_temperature)

	// age, err := strconv.ParseFloat(str_age, 64)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("New var age is: %v\n", age)

	// // ===== Error below =====
	// temperature, err = strconv.Atoi(str_temperature)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("New var temp is: %v\n", temperature)

}
