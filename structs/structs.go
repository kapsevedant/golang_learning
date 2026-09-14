package main

import "fmt"

type Person struct {
	name      string
	age       int
	salary    float32
	city      string
	isMarried bool
}

func newPerson(name string, age int, salary float32, city string) *Person {
	person := Person{
		name:   name,
		age:    age,
		salary: salary,
		city:   city,
	}

	return &person
}

func (P *Person) changeCity(city string) {
	P.city = city
}

func (P Person) returnAge() int {
	return P.age
}
func main() {
	// person := Person{
	// 	name:      "Vedant Kapse",
	// 	age:       23,
	// 	isMarried: false,
	// 	city:      "Pune",
	// }

	person := newPerson("Vedant Kapse", 23, 25000.00, "Pune")

	age := person.returnAge()
	fmt.Println("Name :", person.name)
	fmt.Println("Age :", age)
	fmt.Println("Person Details : ", person)
	person.changeCity("Nagpur")

	fmt.Println("Person Details : ", person)

	// second way of defining struct
	student := struct {
		name      string
		class     string
		isPresent bool
	}{"Abhijit Yadav", "10th", true}
	fmt.Println("Student Details : ", student)

	// struct Embedding

	type Address struct {
		city    string
		state   string
		pincode int
	}

	type Employee struct {
		Person
		Address
		salary float32
	}

	address := Address{
		city:    "Pune",
		state:   "Maharashtra",
		pincode: 411001,
	}

	employee := Employee{
		Person:  *person,
		Address: address,
		salary:  25000.00,
	}

	fmt.Println("Employee Details : ", employee)
}
