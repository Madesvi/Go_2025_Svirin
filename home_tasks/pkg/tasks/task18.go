package tasks

import (
	"fmt"
)

func SolutionTasks18() {

	readers_books := [...]string{"id3", "id5", "id9", "id8", "id2", "id1"} // Даны читатели газет
	readers_magazines := [...]string{"id8", "id2", "id1", "id4", "id6", "id7", "id10"}

	bookReaders := make(map[string]bool)
	for _, reader := range readers_books {
		bookReaders[reader] = true
	}

	var commonReaders []string

	for _, reader := range readers_magazines {
		// Если читатель газет также есть в читателях книг
		if bookReaders[reader] {
			commonReaders = append(commonReaders, reader)
		}
	}

	fmt.Println("Пользователи, которые читают и книги, и газеты:")
	for _, reader := range commonReaders {
		fmt.Println(reader)
	}

	fmt.Printf("\nВсего таких пользователей: %d\n", len(commonReaders))
	fmt.Printf("Их id: %v\n", commonReaders)

}
