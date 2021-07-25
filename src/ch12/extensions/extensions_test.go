package extensions

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) speak() {
	fmt.Println("...")
}

func (p *Pet) speakTo(host string) {
	p.speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) speak() {
	fmt.Println("wang")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.speak()
	dog.speakTo("hello")
}
