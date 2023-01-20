package cours

import "fmt"

type User struct {
	Name    string // public
	age     int    // private
	isAlive bool
}

type Animal struct {
	Name    string // public
	age     int    // private
	isAlive bool
}

type Be interface {
	IsAlive() bool
}

func (u *User) IsAlive() bool {
	return u.isAlive
}

func (u *Animal) IsAlive() bool {
	return u.isAlive
}

func (u *User) Hello() {
	fmt.Println("Bonjour je suis", u.Name)
}

func MakeUser(name string, age int) *User {
	return &User{
		Name:    name,
		age:     age,
		isAlive: true,
	}
}

func MakeAnimal(name string, age int) *Animal {
	return &Animal{
		Name:    name,
		age:     age,
		isAlive: true,
	}
}
