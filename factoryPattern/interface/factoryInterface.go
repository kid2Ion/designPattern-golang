package factory_interface

import "fmt"

type Person interface {
	SayMySelf()
}

type person struct {
	name  string
	age   int
	money int
}

type richPerson struct {
	person person
}

type poorPerson struct {
	person person
}

func (p *person) SayMySelf() {
	fmt.Println("I am Normal Person")
}

func (r *richPerson) SayMySelf() {
	fmt.Println("I am Rich Person")
}

func (p *poorPerson) SayMySelf() {
	fmt.Println("I am Poor Person")
}

func NewPerson(name string, age, money int) Person {
	p := person{
		name:  name,
		age:   age,
		money: money,
	}

	if money >= 100 {
		return &richPerson{
			person: p,
		}
	} else if 0 < money && money < 100 {
		return &p
	} else {
		return &poorPerson{
			person: p,
		}
	}
}

func main() {
	p := NewPerson("hiro", 22, 101)
	p.SayMySelf()
}
