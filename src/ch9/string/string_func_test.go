package string_test

import (
	"log"
	"strconv"
	"strings"
	"testing"
)

func TestStringsFunc(t *testing.T) {
	split := strings.Split("a,b,c", ",")

	for _, s := range split {
		log.Print(s)
	}
	t.Log(strings.Join(split, "-"))
}

func TestStringsConv(t *testing.T) {
	t.Log(strconv.Itoa(10))
}
