package main

import "fmt"

func main() {
	Learning("not")
	Learning("yes")
}

func Learning(s string) error {
	if s == "yes" {
		fmt.Println("learning tine")
		return nil
	}
	fmt.Println("not learning time")
	return fmt.Errorf("oops")
}
