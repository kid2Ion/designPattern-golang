package singleton_test

import (
	"designPattern-golang/singleton"
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := singleton.GetInstance()
	if counter1 == nil {
		t.Error("expected pointer to Singleton after calling GetInstance(). not nil")
	}

	expectedCounter := counter1
	currentCount := counter1.CountUp()
	if currentCount != 1 {
		t.Errorf("After calling CountUp(). the count must be 1 but it is %d\n", currentCount)
	}

	counter2 := singleton.GetInstance()
	if counter2 != expectedCounter {
		t.Error("Expected same instance in counter2 but it got a different instance")
	}

	counter3 := singleton.GetInstance()
	currentCount = counter3.CountDown()
	if currentCount != 0 {
		t.Errorf("After calling must be 0 but it is %d\n", currentCount)
	}
}
