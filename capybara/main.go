package main

import "fmt"

func PrintInfo(u Capybara) {
	fmt.Println("капибебра: ", u.name, u.surname, "возрастом:", u.age)
}

type Capybara struct {
	name    string
	surname string
	age     int
}

func (c Capybara) printC() {
	fmt.Println("капибебра: ", c.name, c.surname, "возрастом:", c.age)

}

func main() {
	fmt.Println("hi capy")

	capy := Capybara{
		name:    "capik",
		surname: "capikov",
		age:     3,
	}

	PrintInfo(capy)
	//PrintInfo(c)

}
