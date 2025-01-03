package main

import (
	"fmt"
	"log"
)

func main() {
	Learning("not")
	err := Learning("yes")
	if err != nil {
		log.Fatal(err)
	}
}

func Learning(s string) error {
	if s == "yes" {
		fmt.Println("learning tine")
		return nil
	}
	fmt.Println("not learning time")
	return fmt.Errorf("oops")
}

// var do_something string = "AAA"
// fmt.Println(do_something)
// on merge
