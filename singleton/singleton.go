package singleton

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) CountUp() int {
	s.count++
	return s.count
}

func (s *singleton) CountDown() int {
	s.count--
	return s.count
}
