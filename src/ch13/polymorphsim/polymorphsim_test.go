package polymorphsim_test

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.println"
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWorld() string {
	return "System.out.println"
}

func printHelloWorld(p Programmer) {
	fmt.Println(p.WriteHelloWorld())
}

func TestPolymorphsim(t *testing.T) {
	goProgrammer := &GoProgrammer{}
	javaProgrammer := &JavaProgrammer{}
	printHelloWorld(goProgrammer)
	printHelloWorld(javaProgrammer)
}
