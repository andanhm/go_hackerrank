package main

import "fmt"

type person struct {
	age int
}
func (p person) NewPerson(initialAge int) person {
	//Add some more code to run some checks on initialAge
	if (initialAge > 0) {
		p.age = initialAge
	} else {
		p.age = 0
		fmt.Println("Age is not valid, setting age to 0.")
	}
	return p
}

func (p person) amIOld() {
	var personAge int =p.age
	if (personAge < 13) {
		fmt.Println("You are young.")
	} else if (personAge >= 13 && personAge < 18) {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are old.")
	}

}

func (p person) yearPasses() person {
	//Increment the age of the person in here
	p.age = p.age + 1
	return p
}

func main() {
	var T, age int

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&age)
		p := person{age: age}
		p = p.NewPerson(age)
		p.amIOld()
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}