package main

import (
	"fmt"
	"testing"
)

const (
	a = iota - 1
	b
	c
)

func Test(t *testing.T) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
