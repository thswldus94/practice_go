package main

import "fmt"

func main() {
	printNum()
	printGugu()
	drawStar()
}

func printNum() {
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}
}

func printGugu() {
	for i := 2; i < 10; i++ {
		if i >= 3 && i <= 6 {
			continue
		}

		for j := 1; j < 10; j++ {
			fmt.Println(i, "*", j, "=", i*j)
		}
		fmt.Println()
	}
}

func drawStar() {
	for i := 5; i >= 1; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
