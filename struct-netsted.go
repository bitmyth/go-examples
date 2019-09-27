// https://medium.com/rungo/structures-in-go-76377cc106a2
package main

import "fmt"

// You are allowed to mix some anonymous fields within named fields
type Employee struct {
	firstName, lastName string
	*Salary // ☛ Nested struct
	bool
}

type Salaried interface {
	getSalary() int
}

type Salary struct {
	basic     int
	insurance int
	allowance int
}

func (s *Salary) getSalary() int {
	return s.basic + s.insurance + s.allowance
}

func main() {
	ross := Employee{
		firstName: "Ross",
		lastName:  "Geller",
		Salary:    &Salary{1000, 10, 50},
		bool:      true,
	}
	// But the cool thing about Go is when we use an anonymous nested struct,
	// all the nested struct fields are automatically available in parent struct.
	ross.basic = 1200
	ross.insurance = 0
	ross.allowance = 0
	// If a nested anonymous struct has a same field (field name) as in the parent struct,
	// then inner struct field won’t get promoted. Only non-conflicting fields will get promoted.
	fmt.Println(ross.getSalary())
}
