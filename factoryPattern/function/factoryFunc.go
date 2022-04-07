package factory_func

type Person struct {
	Name  string
	Age   int
	Money int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name:  name,
		Age:   age,
		Money: 1000,
	}
}
