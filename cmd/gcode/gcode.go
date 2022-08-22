package main

import "fmt"

func main() {
	fmt.Println(greet("World"))
}

func greet(who string) string {
	return fmt.Sprintf("Hi %v", who)
}
