package entity

type Status struct {
	id   int
	name string
}

func NewStatus(id int, name string) *Status {
	status := new(Status)
	status.id = id
	status.name = name
	return status
}
