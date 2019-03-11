package main

import "fmt"

func printMap() {
	m := make(map[string]int)
        m["kalpesh"] = 32
        m["ujwal"] = 27
        fmt.Println(": ", m)
}

func printStruct() {
	type developer struct {
		name string
		designation string
	}

	kalpesh := developer{"Kalpesh Tech Lead", "Tech Lead"}
	fmt.Println(kalpesh)
	chujwal := developer{name: "Chujwal",  designation: "Software Engineer"}
	fmt.Println(chujwal)
}

func main() {
	printMap()
	printStruct()
}

