package main

import "fmt"

func main() {
	var cat Cat = NewCat("Som")
	var dog Dog = NewDog("Dum")

	var list []Animal
	list = append(list, &dog)
	list = append(list, &cat)

	for i := range list {
		list[i].makeSound()
	}

	// Expect to see 15
	arr := []int{1, 2, 3, 4, 5}
	result := reduce(arr, sum)
	fmt.Println(result)

	ll := NewLinkedList()
	for i := 0; i < 10; i++ {
		ll.Append(i + 1)
	}
	result = ll.reduce(sum)
	fmt.Println(result)

	x := new(Node)
	fmt.Println(x)
}

func sum(a, b int) int {
	return a + b
}
