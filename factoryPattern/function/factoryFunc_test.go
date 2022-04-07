package factory_func_test

import (
	factory_func "designPattern-golang/factoryPattern/function"
	"fmt"
	"testing"
)

func TestNewPerson(t *testing.T) {
	name := "hiro"
	age := 22
	p := factory_func.NewPerson(name, age)
	testCase := fmt.Sprintf("%s is %d", p.Name, p.Age)
	if testCase != "hiro is 22" {
		t.Errorf("test failed")
	}
}
