package main

import (
	"fmt"
	"go-api-practice/pkg/config"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(config.Get("id"))
}
