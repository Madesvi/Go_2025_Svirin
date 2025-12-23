package tasks

import (
	"fmt"
)

func SolutionTasks17() {

	// Все пользователи
	var all_users = [...]string{"id3", "id5", "id9", "id8",
		"id2", "id1", "id4", "id6", "id7", "id10"}

	var offline_users = [...]string{"id3", "id9", "id7", "id2",
		"id4", "id6"}

	var online_users []string

	onlineMap := make(map[string]bool)

	for _, user := range offline_users {
		onlineMap[user] = true
	}

	for _, user := range all_users {
		if !onlineMap[user] {
			online_users = append(online_users, user)
		}
	}
	fmt.Println(online_users)
}
