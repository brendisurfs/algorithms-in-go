package main

import "fmt"

func main() {
	nth := GetNthFib(2)
	fmt.Println(nth)
}

func GetNthFib(n int) int {

	if n == 2 {
		return 1
	}

	if n == 1 {
		return 0
	}

}
