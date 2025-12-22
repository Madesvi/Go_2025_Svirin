package tasks

import (
	"fmt"
	"math/rand"
)

type mainWorkers interface {
	aboutPerson() string
	TakeDamage(damage int)
	GetHP() int
}

type Person struct {
	Name   string
	Level  int
	Skills int
	Class  int
	Hp     int
}

type AtributesForPerson struct {
	Warrior    bool
	Mage       bool
	Archer     bool
	PersonInfo Person
}

func (a *AtributesForPerson) aboutPerson() string {
	name := a.PersonInfo.Name
	level := a.PersonInfo.Level
	hp := a.PersonInfo.Hp
	info := fmt.Sprintf("Персонаж: %s, (Level: %d), HP: %d", name, level, hp)
	return info
}

func (a *AtributesForPerson) GetHP() int {
	hp := a.PersonInfo.Hp
	return hp
}

func (a *AtributesForPerson) TakeDamage(damage int) {
	a.PersonInfo.Hp -= damage
	if a.PersonInfo.Hp < 0 {
		a.PersonInfo.Hp = 0
	}
	fmt.Printf("%s получает %d урона. Осталось HP: %d\n", a.PersonInfo.Name, damage, a.PersonInfo.Hp)
}

func (p *Person) attack(target mainWorkers, damage int) {
	fmt.Printf("%s атакует с уроном %d!\n", p.Name, damage)
	target.TakeDamage(damage)
}

func SolutionTasks16() {

	person1 := &AtributesForPerson{
		Warrior: true,
		PersonInfo: Person{
			Name:   "Kekik",
			Level:  55,
			Skills: 1000,
			Hp:     100,
		},
	}
	person2 := &AtributesForPerson{
		Mage: true,
		PersonInfo: Person{
			Name:   "Fookik",
			Level:  75,
			Skills: 1500,
			Hp:     100,
		},
	}
	// fmt.Println(person1.aboutPerson())
	// fmt.Println(person1.GetHP())
	// fmt.Println(person2.aboutPerson())
	// fmt.Println(person2.GetHP())

	randomNum := rand.Intn(100) + 1
	// fmt.Printf("Случайное число от 1 до 99: %d\n", randomNum)
	person1.PersonInfo.attack(person2, randomNum)
	randomNum = rand.Intn(100) + 1
	// fmt.Printf("Случайное число от 1 до 99: %d\n", randomNum)
	person2.PersonInfo.attack(person1, randomNum)
	randomNum = rand.Intn(100) + 1
	// fmt.Printf("Случайное число от 1 до 99: %d\n", randomNum)
	person1.PersonInfo.attack(person2, randomNum)
	randomNum = rand.Intn(100) + 1
	// fmt.Printf("Случайное число от 1 до 99: %d\n", randomNum)
	person2.PersonInfo.attack(person1, randomNum)
}
