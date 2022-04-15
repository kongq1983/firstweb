package dto

import "fmt"

type Employee struct {
	Id       int
	Username string
	Name     string
	Age      int
}

// requestBody接收，字段需要大写，否则值是接收不到的
//type Employee1 struct {
//	Id    int
//	Username string
//	Name  string
//	Age int
//}

func (s *Employee) String() string {
	return fmt.Sprintf("[id => %d username=%s , name => %v, age => %v]", s.Id, s.Username, s.Name, s.Age)
}
