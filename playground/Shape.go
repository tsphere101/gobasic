package main

import "fmt"

type Dog struct {
	name  string
	sound string
}

func NewDog(name string) Dog {
	return Dog{name: name, sound: "Hong"}
}
func (dog *Dog) makeSound() {
	fmt.Println(dog.name + " : " + dog.sound)
}

type Cat struct {
	name  string
	sound string
}

func NewCat(name string) Cat {
	return Cat{name: name, sound: "Meawww"}
}

func (cat *Cat) makeSound() {
	fmt.Println(cat.name + " : " + cat.sound)
}

type Animal interface {
	makeSound()
}
