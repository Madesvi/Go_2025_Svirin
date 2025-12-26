package tasks

import "fmt"

type Node struct {
	name string
	time int
	dist float32
	next *Node
}

type Route struct {
	head   *Node
	length int
}

// Next возвращает следующий узел
func (n *Node) Next() *Node {
	return n.next
}

// insertAtHead вставляет новый узел в начало списка
func (r *Route) insertAtHead(name string, time int, dist float32) {
	nodeHead := &Node{name, time, dist, nil}
	if r.head == nil {
		r.head = nodeHead
	} else {
		nodeHead.next = r.head
		r.head = nodeHead
	}
	r.length += 1
}

// calcDistance рассчитывает общее время и расстояние от узла X до узла Y
func (r *Route) calcDistance(x *Node, y *Node) (int, float32, error) {
	if r.head == nil {
		return 0, 0, fmt.Errorf("маршрут пуст")
	}

	// Находим узел X
	current := r.head
	foundX := false
	for current != nil {
		if current == x {
			foundX = true
			break
		}
		current = current.next
	}

	if !foundX {
		return 0, 0, fmt.Errorf("узел X не найден в маршруте")
	}

	// Считаем время и расстояние от X до Y
	totalTime := 0
	totalDist := float32(0.0)
	foundY := false

	for current != nil {
		totalTime += current.time
		totalDist += current.dist

		if current == y {
			foundY = true
			break
		}
		current = current.next
	}

	if !foundY {
		return 0, 0, fmt.Errorf("узел Y не найден после узла X")
	}

	return totalTime, totalDist, nil
}

func (r *Route) findNodeByName(name string) *Node {
	current := r.head
	for current != nil {
		if current.name == name {
			return current
		}
		current = current.next
	}
	return nil
}

func SolutionTasks19() {
	list := Route{nil, 0}
	list.insertAtHead("С", 15, 18)
	list.insertAtHead("B", 5, 6)
	list.insertAtHead("A", 0, 0)

	// Демонстрация метода Next()
	fmt.Println("Демонстрация метода Next():")
	node := list.head
	for node != nil {
		fmt.Printf("Узел: %s, Время: %d, Расстояние: %.2f\n",
			node.name, node.time, node.dist)
		node = node.Next()
	}

	// Демонстрация метода calcDistance()
	fmt.Println("\nДемонстрация calcDistance():")
	nodeA := list.findNodeByName("A")
	nodeB := list.findNodeByName("B")
	nodeC := list.findNodeByName("С")

	if nodeA != nil && nodeC != nil {
		totalTime, totalDist, err := list.calcDistance(nodeA, nodeC)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Printf("От A до C: Время = %d, Расстояние = %.2f\n",
				totalTime, totalDist)
		}
	}

	if nodeB != nil && nodeC != nil {
		totalTime, totalDist, err := list.calcDistance(nodeB, nodeC)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Printf("От B до C: Время = %d, Расстояние = %.2f\n",
				totalTime, totalDist)
		}
	}

}
