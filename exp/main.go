package main

import (
	"html/template"
	"os"
)

type Dog struct {
	Name string
}
type User struct {
	Name  string
	Dog   Dog
	Age   int
	Map   map[string]int
	Float float64
	Slice []string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name: "Justin Persaud",
		Dog: Dog{
			Name: "Spike",
		},
		Age:   32,
		Float: 3.14,
		Slice: []string{"a", "b", "c"},
		Map: map[string]int{
			"abc": 123,
			"def": 456,
		},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
