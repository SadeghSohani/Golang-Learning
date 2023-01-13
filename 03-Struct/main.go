package main

import "fmt"

type person struct {
	Name   string
	Family string
	Age    int
}

func main() {

	persons := map[string]person{
		"Sadegh": person{
			Name:   "Sadegh",
			Family: "Sohani",
			Age:    24,
		},
		"Reza": person{
			Name:   "Reza",
			Family: "Rezazadeh",
			Age:    26,
		},
	}

	fmt.Println(persons["Sadegh"])
	fmt.Println(persons["Sadegh"].Age)

}
