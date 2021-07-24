package encapsulations

import (
	"fmt"
	"testing"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	employee1 := Employee{"1", "bob", 3}
	employee2 := Employee{Name: "Mike", Age: 4}
	employee3 := new(Employee)
	employee3.Id = "2"
	employee3.Name = "Rose"
	employee3.Age = 5

	t.Log(employee1.Name)
	t.Log(employee2.Name)
	t.Log(employee3.Name)

	t.Logf("employee1 type %T", employee1)
	t.Logf("employee3 type %T", employee3)
}

//方法(行为)
func (e Employee) StringMethod() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func (e *Employee) StringPointerMethod() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e1 := Employee{"0", "Bob", 3}
	t.Log(e1.StringMethod())
	e2 := &Employee{"0", "Bob", 3}
	t.Log(e2.StringPointerMethod())
}
