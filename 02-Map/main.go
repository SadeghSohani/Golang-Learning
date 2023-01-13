package main

import "fmt"

func main() {
	//var ages = make(map[string]int)
	ages := make(map[string]int)
	ages["Sadegh"] = 24
	ages["Jimi"] = 25
	fmt.Println(ages["Sadegh"])
	delete(ages, "Jimi")
	fmt.Println(ages)

	lastName := map[string]string{
		"Sadegh": "Sohani",
	}
	fmt.Println(lastName)

	if v, ok := lastName["Sadegh"]; ok == true {
		fmt.Println(v)
	}

}
