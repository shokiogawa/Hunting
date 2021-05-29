package entity

type Practice struct {
	Id   int
	Name string
	Age  int
}

func NewPractice(id int, name string, age int) *Practice {
	p := new(Practice)
	p.Id = id
	p.Name = name
	p.Age = age
	return p
}
