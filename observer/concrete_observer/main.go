package concreteobserver

import "fmt"

type Customer struct {
	Id string
}

func (c Customer) Update(str string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.Id, str)
}

func (c Customer) GetId() string {
	return c.Id
}
