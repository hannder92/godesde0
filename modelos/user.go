package modelos

import "time"

type User struct {
	Id        int
	Name      string
	Age       int
	CreatedAt time.Time
	Status    bool
}

func (usuario *User) AddUser(id int, name string, age int, createdAt time.Time, status bool) {
	usuario.Id = id
	usuario.Name = name
	usuario.Age = age
	usuario.CreatedAt = createdAt
	usuario.Status = status
}
