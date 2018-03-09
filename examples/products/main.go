package main

import (
	"fmt"

	"github.com/brianseitel/charlatan"
)

// Product ...
type Product struct {
	UUID       string  `charlatan:"uuid"`
	Name       string  `charlatan:"name"`
	Brand      string  `charlatan:"name"`
	Price      float64 `charlatan:"price"`
	Categories struct {
		Name string `charlatan:"word"`
	} `charlatan:"struct"`
	Image     string `charlatan:"url"`
	OnSale    bool   `charlatan:"boolean"`
	UpdatedAt string `charlatan:"datetime"`
}

func main() {
	c := charlatan.New()

	foo, err := c.Generate(&Product{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%# v\n", foo)
}
